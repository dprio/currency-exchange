package dollarexchangerateapigateway

import (
	"context"
	"time"

	"github.com/dprio/currency-exchange/client/internal/infrastructure/httpclient/dollarexchangerate"
)

type Gateway interface {
	GetDollarExchangeRate(ctx context.Context) (float32, error)
}

type (
	client interface {
		GetDollarExchangeRate(ctx context.Context) (*dollarexchangerate.DollarExchangeRateResponse, error)
	}

	gateway struct {
		client client
	}
)

func New(client client) Gateway {
	return &gateway{
		client: client,
	}
}

func (g *gateway) GetDollarExchangeRate(ctx context.Context) (float32, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	dollarQuote, err := g.client.GetDollarExchangeRate(ctxTimeout)
	if err != nil {
		return 0, err
	}

	return dollarQuote.Bid, nil
}
