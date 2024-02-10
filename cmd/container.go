package main

import (
	"context"
	"os"

	"github.com/luchojuarez/call-sorter/internal/domain/ingestdata"
	"github.com/luchojuarez/call-sorter/internal/domain/invoice"
	"github.com/luchojuarez/call-sorter/internal/infrastructure/callservice"
	"github.com/luchojuarez/call-sorter/internal/infrastructure/user"
)

type InvoiceContainer struct {
	callRepo invoice.CallRepository
	userRepo invoice.UserRepository
}

func GetSimpleContainer() *InvoiceContainer {
	return &InvoiceContainer{
		callRepo: getLocalCallRepository(),
		userRepo: getRestUserRepository(),
	}
}

func getLocalCallRepository() invoice.CallRepository {
	callRepo := callservice.NewClient()
	file, err := os.Open("cmd/input.csv")
	if err != nil {
		panic(err)
	}
	if err := ingestdata.NewClient(callRepo).ReadAll(context.Background(), file); err != nil {
		panic(err)
	}
	return *callRepo
}

func getRestUserRepository() invoice.UserRepository {
	userRepo := user.NewRepository("https://fn-interview-api.azurewebsites.net")
	return userRepo
}
