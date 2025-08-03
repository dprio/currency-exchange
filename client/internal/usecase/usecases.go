package usecase

import (
	"github.com/dprio/currency-exchange/client/internal/gateway/httpgateway"
	"github.com/dprio/currency-exchange/client/internal/usecase/getdollarexchangerate"
)

type UseCases struct {
	GetDollarExchangeRateUseCase getdollarexchangerate.UseCase
}

func New(httpGateways httpgateway.HTTPGateway) UseCases {
	return UseCases{
		GetDollarExchangeRateUseCase: getdollarexchangerate.New(httpGateways.DollarQuoteAPIGateway),
	}
}
