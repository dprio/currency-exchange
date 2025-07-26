package economiaapiagateway

import (
	"context"

	"github.com/dprio/currency-exchange/server/internal/domain/dollarquote"
	"github.com/dprio/currency-exchange/server/internal/infrastructure/httpclient/economiapi"
)

type Adapter interface {
	GetUSDQuote(ctx context.Context) (*dollarquote.DollarQuote, error)
}

type (
	client interface {
		GetDollarQuote(ctx context.Context) (*economiapi.ExchangeRateResponse, error)
	}

	adapter struct {
		client client
	}
)

func New(client client) Adapter {
	return &adapter{
		client: client,
	}
}

func (a *adapter) GetUSDQuote(ctx context.Context) (*dollarquote.DollarQuote, error) {
	resp, err := a.client.GetDollarQuote(ctx)
	if err != nil {
		return nil, err
	}

	return resp.ToDollarQuote()
}
