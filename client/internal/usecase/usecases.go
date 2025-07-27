package usecase

import (
	"github.com/dprio/currency-exchange/client/internal/gateway/httpgateway"
	getdollarquote "github.com/dprio/currency-exchange/client/internal/usecase/getDollarQuote"
)

type UseCases struct {
	GetDollarQuoteUseCase getdollarquote.UseCase
}

func New(httpGateways httpgateway.HTTPGateway) UseCases {
	return UseCases{
		GetDollarQuoteUseCase: getdollarquote.New(httpGateways.DollarQuoteAPIGateway),
	}
}
