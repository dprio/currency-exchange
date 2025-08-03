package usecase

import (
	"github.com/dprio/currency-exchange/server/internal/gateway/httpgateway"
	"github.com/dprio/currency-exchange/server/internal/gateway/repository"
	"github.com/dprio/currency-exchange/server/internal/usecase/getdollarexchangerate"
)

type UseCase struct {
	GetDollarQuoteUseCase getdollarexchangerate.UseCase
}

func New(adapters httpgateway.HTTPGateways, repositories repository.Repositories) UseCase {
	return UseCase{
		GetDollarQuoteUseCase: getdollarexchangerate.New(adapters.EconomiaAPIAdapter, repositories.DollarQuoteRepository),
	}
}
