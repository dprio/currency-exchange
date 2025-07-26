package usecase

import (
	"github.com/dprio/currency-exchange/server/internal/gateway/httpgateway"
	"github.com/dprio/currency-exchange/server/internal/gateway/repository"
	"github.com/dprio/currency-exchange/server/internal/usecase/getdollarquote"
)

type UseCase struct {
	GetDollarQuoteUseCase getdollarquote.UseCase
}

func New(adapters httpgateway.HTTPGateways, repositories repository.Repositories) UseCase {
	return UseCase{
		GetDollarQuoteUseCase: getdollarquote.New(adapters.EconomiaAPIAdapter, repositories.DollarQuoteRepository),
	}
}
