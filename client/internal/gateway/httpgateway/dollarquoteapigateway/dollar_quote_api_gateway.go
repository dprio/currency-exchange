package dollarquoteapigateway

import (
	"context"
	"time"

	"github.com/dprio/currency-exchange/client/internal/infrastructure/httpclient/dollarquote"
)

type Gateway interface {
	GetDollarQuote(ctx context.Context) (float32, error)
}

type (
	client interface {
		GetDollarQuote(ctx context.Context) (*dollarquote.DollarQuoteResponse, error)
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

func (g *gateway) GetDollarQuote(ctx context.Context) (float32, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	dollarQuote, err := g.client.GetDollarQuote(ctxTimeout)
	if err != nil {
		return 0, err
	}

	return dollarQuote.Bid, nil
}
