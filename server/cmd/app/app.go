package app

import (
	"fmt"
	"net/http"

	"github.com/dprio/currency-exchange/server/cmd/handlers"
	"github.com/dprio/currency-exchange/server/internal/gateway/httpgateway"
	"github.com/dprio/currency-exchange/server/internal/gateway/repository"
	"github.com/dprio/currency-exchange/server/internal/infrastructure/db"
	"github.com/dprio/currency-exchange/server/internal/infrastructure/httpclient"
	"github.com/dprio/currency-exchange/server/internal/usecase"
)

type App struct {
	HTTPClients httpclient.HTTPClient
	UseCases    usecase.UseCase
	Handlers    handlers.Handlers
	mux         *http.ServeMux
}

func New() *App {

	dataBase := db.New()

	httpClients := httpclient.New()

	httpGateways := httpgateway.New(httpClients)

	repositories := repository.New(dataBase)

	usecases := usecase.New(httpGateways, repositories)

	handlers := handlers.New(usecases)

	mux := http.NewServeMux()
	registerRoutes(mux, handlers)

	return &App{
		HTTPClients: httpClients,
		UseCases:    usecases,
		Handlers:    handlers,
		mux:         mux,
	}
}

func (a *App) Start() {
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", a.mux)
}

func registerRoutes(mux *http.ServeMux, handlers handlers.Handlers) {
	mux.HandleFunc("/cotacao", handlers.GetCotacaoHandler.Serve)
}
