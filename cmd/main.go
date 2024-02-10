package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/luchojuarez/call-sorter/internal/domain/invoice"
)

func main() {
	ctx := context.Background()
	container := GetSimpleContainer()

	proccesor := invoice.NewProcessor(container.callRepo, container.userRepo)
	result, err := proccesor.Generate(ctx, "+5491167980953", time.April, 2020)
	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", string(b))
}
