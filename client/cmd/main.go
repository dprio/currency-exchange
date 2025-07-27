package main

import (
	"context"
	"fmt"

	"github.com/dprio/currency-exchange/client/cmd/app"
)

func main() {
	application := app.New()
	fmt.Println("Client is running")

	application.Start(context.Background())
}
