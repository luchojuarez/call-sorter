package container

import (
	"context"
	"os"

	"github.com/luchojuarez/call-sorter/internal/domain/ingestdata"
	"github.com/luchojuarez/call-sorter/internal/infrastructure/callservice"
	"github.com/luchojuarez/call-sorter/internal/infrastructure/user"
)

type InvoiceContainer struct {
	callRepo *callservice.Client
	userRepo *user.Repository
}

func GetSimpleContainer() *InvoiceContainer {
	return &InvoiceContainer{}
}

func (ic InvoiceContainer) GetLocalCallRepository() *callservice.Client {
	if ic.callRepo != nil {
		return ic.callRepo
	}
	ic.callRepo = callservice.NewClient()
	file, err := os.Open("cmd/input.csv")
	if err != nil {
		panic(err)
	}
	if err := ingestdata.NewClient(ic.callRepo).ReadAll(context.Background(), file); err != nil {
		panic(err)
	}
	return ic.callRepo
}

func (ic InvoiceContainer) GetRestUserRepository() *user.Repository {
	if ic.userRepo != nil {
		return ic.userRepo
	}
	ic.userRepo = user.NewRepository("https://fn-interview-api.azurewebsites.net")
	return ic.userRepo
}
