package invoice

import (
	"context"
	"fmt"
	"time"

	"github.com/luchojuarez/call-sorter/internal/infrastructure/callservice"
	"github.com/luchojuarez/call-sorter/internal/infrastructure/user"
)

// free friend calls: better in configs file.
const (
	totalfreeFriendsCalls           = 10
	nationalCallPricePerCall        = float64(2.5)
	internationalCallPricePerSecond = float64(0.75)
)

// best implementation: External service; this information could be useful for the rest of de company.
type CallRepository interface {
	FindByPhoneAndMonthAndYear(ctx context.Context, phoneNumber string, month time.Month, year int) ([]callservice.Model, error)
}

// External service.
type UserRepository interface {
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (user.Model, error)
}

type Processor struct {
	callRepo CallRepository
	userRepo UserRepository
}

func NewProcessor(callRepo CallRepository, userRepo UserRepository) Processor {
	return Processor{
		callRepo: callRepo,
		userRepo: userRepo,
	}
}

func (p Processor) Generate(ctx context.Context, phoneNumber string, month time.Month, year int) (Model, error) {
	newInvoice := Model{}
	user, err := p.userRepo.GetByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		return newInvoice, fmt.Errorf("user repository error: %w", err)
	}

	callList, err := p.callRepo.FindByPhoneAndMonthAndYear(ctx, phoneNumber, month, year)
	if err != nil {
		return newInvoice, fmt.Errorf("call repository error: %w", err)
	}

	friendCallsCounter := 0

	for _, call := range callList {
		isFrend := isFrend(user, call.RecipientNumber)
		callAmount := caluclateCallAmount(user, call)
		if isFrend {
			if friendCallsCounter < totalfreeFriendsCalls {
				callAmount = 0
			}
			friendCallsCounter++

			//frend calls count as an international and national call counter!
			newInvoice.TotalFriendsSeconds = newInvoice.TotalFriendsSeconds + call.Duration
		}

		switch call.CallType {
		case callservice.CallTypeInternational:
			newInvoice.TotalInternationalSeconds = newInvoice.TotalInternationalSeconds + call.Duration
		case callservice.CallTypeNational:
			newInvoice.TotalNationalSeconds = newInvoice.TotalNationalSeconds + call.Duration
		}

		InvoiceCall := FromCallRepo(call, callAmount)
		newInvoice.Calls = append(newInvoice.Calls, InvoiceCall)
		newInvoice.Total = newInvoice.Total + callAmount
	}

	newInvoice.User = FromUserRepo(user)
	newInvoice.Month = month

	return newInvoice, nil
}

func caluclateCallAmount(user user.Model, call callservice.Model) float64 {
	switch call.CallType {
	case callservice.CallTypeInternational:
		return internationalCallPricePerSecond * float64(call.Duration)
	case callservice.CallTypeNational:
		return nationalCallPricePerCall
	default:
		panic("invalid call type for pricing: " + call.CallType)
	}
}

func isFrend(user user.Model, phoneNumber string) bool {
	for _, friendNumber := range user.Friends {
		if friendNumber == phoneNumber {
			return true
		}
	}
	return false
}
