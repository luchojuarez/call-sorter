package ingestdata

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/luchojuarez/call-sorter/internal/infrastructure/callservice"
	"go.uber.org/zap"
)

// TODO: choose a better logger
var log = zap.NewNop()

type CallRepository interface {
	Save(ctx context.Context, model callservice.Model) error
}

type Client struct {
	callRepository CallRepository
}

func NewClient(repo CallRepository) Client {
	return Client{
		callRepository: repo,
	}
}

func (c Client) ReadAll(ctx context.Context, input io.Reader) error {
	csvReader := csv.NewReader(input)
	records, err := csvReader.ReadAll()
	if err != nil {
		return fmt.Errorf("error in CSV reader: %w", err)
	}
	if len(records) < 2 {
		return ErrInvalidCsvInput
	}

	var readErrors []error

	for i := 1; i < len(records); i++ {
		csvCall := records[i]
		if len(csvCall) != 4 {
			readErrors = append(readErrors, errors.New(fmt.Sprintf("invalid row: %d", i)))
		} else {
			callDuration := -1
			if callDuration, err = strconv.Atoi(csvCall[2]); err != nil {
				readErrors = append(readErrors, errors.New(fmt.Sprintf("invalid call duration format '%s' at: %d", csvCall[2], i)))
				continue
			}

			callDate := time.Time{}
			if callDate, err = time.Parse(time.RFC3339, csvCall[3]); err != nil {
				readErrors = append(readErrors, errors.New(fmt.Sprintf("invalid call date format '%s' at: %d", csvCall[3], i)))
				continue
			}

			call := callservice.Model{
				OriginNumber:    csvCall[0],
				RecipientNumber: csvCall[1],
				Duration:        callDuration,
				Date:            callDate,
			}
			call.CallType = getCallType(call.OriginNumber, call.RecipientNumber)

			if err := c.callRepository.Save(ctx, call); err != nil {
				readErrors = append(readErrors, errors.New(fmt.Sprintf("error saving call '%+v' at: %d", call, i)))
			}
		}

	}

	return errors.Join(readErrors...)
}

func getCallType(originNumber, recipientNumber string) string {
	if originNumber[0:3] == recipientNumber[0:3] {
		return callservice.CallTypeNational
	}
	return callservice.CallTypeInternational
}
