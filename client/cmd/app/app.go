package app

import (
	"context"
	"fmt"

	"github.com/dprio/currency-exchange/client/internal/gateway/httpgateway"
	"github.com/dprio/currency-exchange/client/internal/infrastructure/httpclient"
	"github.com/dprio/currency-exchange/client/internal/usecase"
)

type App struct {
	UseCases usecase.UseCases
}

func New() *App {

	httpClients := httpclient.New()

	httpGateways := httpgateway.New(httpClients)

	useCases := usecase.New(httpGateways)

	return &App{
		UseCases: useCases,
	}

}

func (a *App) Start(ctx context.Context) {

	fmt.Println("Application is running =) ")

	var input string
	for {
		fmt.Print("> ")
		fmt.Scan(&input)

		switch input {
		case "quote":
			quote, err := a.UseCases.GetDollarQuoteUseCase.Execute(ctx)
			if err != nil {
				continue
			}
			fmt.Printf("Dollar Quote -> US$ %f\n", quote)
		case "exit":
			fmt.Println("Parando a execução...")
			return
		default:
			fmt.Println("Comando não encontrado !")
		}
	}
}
