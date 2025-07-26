package economiaapiagateway

import (
	"context"
	"time"

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
	ctxTimeout, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	resp, err := a.client.GetDollarQuote(ctxTimeout)
	if err != nil {
		return nil, err
	}

	return resp.ToDollarQuote()
}
