package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

const uri = "/users/"

type Repository struct {
	restClient *resty.Client
}

func NewRepository(host string) *Repository {
	return &Repository{
		restClient: resty.New().SetHostURL(host),
	}
}

func (r Repository) GetByPhoneNumber(ctx context.Context, phoneNumber string) (Model, error) {

	resp, err := r.restClient.R().
		ForceContentType("application/json").
		Get(uri + phoneNumber)

	if err != nil {
		return Model{}, fmt.Errorf("rest error: %w", err)
	}

	switch resp.StatusCode() {
	case http.StatusOK:
		response := Model{}

		if err = json.Unmarshal(resp.Body(), &response); err != nil {
			return Model{}, fmt.Errorf("unmarshal error: %w", err)
		}
		return response, nil
	default:
		return Model{}, errors.New("invalid response status: " + resp.Status())
	}

}
