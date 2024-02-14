package callservice

import (
	"context"
	"sort"
	"time"
)

type Client struct {
	repo *[]Model
}

func NewClient() *Client {
	return &Client{
		repo: &[]Model{},
	}
}

func NewInMemoryClient(initialData []Model) Client {
	return Client{
		repo: &initialData,
	}
}

// Return calls in chronological order.
func (cli Client) FindByPhoneAndMonthAndYear(ctx context.Context, phoneNumber string, month time.Month, year int) ([]Model, error) {
	result := []Model{}
	for _, call := range *cli.repo {
		if call.OriginNumber == phoneNumber && call.Date.Year() == year && call.Date.Month() == month {
			result = append(result, call)
		}
	}

	sort.Slice(result, func(i, j int) bool { return result[i].Date.Before(result[j].Date) })

	return result, nil
}

func (cli Client) Save(ctx context.Context, model Model) error {
	*(cli.repo) = append(*cli.repo, model)
	return nil
}
