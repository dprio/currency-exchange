package handlers

import (
	"github.com/dprio/currency-exchange/server/cmd/handlers/cotacao/get"
	"github.com/dprio/currency-exchange/server/internal/usecase"
	"github.com/dprio/currency-exchange/server/pkg/handler"
)

type Handlers struct {
	GetCotacaoHandler handler.HandlerInterface
}

func New(usecase usecase.UseCase) Handlers {
	return Handlers{
		GetCotacaoHandler: get.New(usecase.GetDollarQuoteUseCase),
	}
}
