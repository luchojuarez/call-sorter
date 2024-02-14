package invoice

import (
	"time"

	"github.com/luchojuarez/call-sorter/internal/infrastructure/callservice"
	user "github.com/luchojuarez/call-sorter/internal/infrastructure/user"
)

type Model struct {
	Month                     time.Month  `json:"-"`
	User                      InvoiceUser `json:"user"`
	Calls                     []Call      `json:"calls"`
	TotalInternationalSeconds int         `json:"total_international_seconds"`
	TotalNationalSeconds      int         `json:"total_national_seconds"`
	TotalFriendsSeconds       int         `json:"total_friends_seconds"`
	Total                     float64     `json:"total"`
}

type InvoiceUser struct {
	Address     string `json:"address"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

func FromUserRepo(input user.Model) InvoiceUser {
	return InvoiceUser{
		Address:     input.Address,
		Name:        input.Name,
		PhoneNumber: input.PhoneNumber,
	}
}

type Call struct {
	PhoneNumber string    `json:"phone_number"`
	Duration    int       `json:"duration"`
	Timestamp   time.Time `json:"timestamp"`
	Amount      float64   `json:"amount"`
}

func FromCallRepo(input callservice.Model, amount float64) Call {
	return Call{
		PhoneNumber: input.RecipientNumber,
		Duration:    input.Duration,
		Timestamp:   input.Date,
		Amount:      amount,
	}
}
