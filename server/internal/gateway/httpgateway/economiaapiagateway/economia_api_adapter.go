package economiaapiagateway

import (
	"context"
	"fmt"
	"time"

	"github.com/dprio/currency-exchange/server/internal/domain/dollarexchangerate"
	"github.com/dprio/currency-exchange/server/internal/infrastructure/httpclient/economiapi"
)

type Adapter interface {
	GetUSDExchangeRate(ctx context.Context) (*dollarexchangerate.DollarExchangeRate, error)
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

func (a *adapter) GetUSDExchangeRate(ctx context.Context) (*dollarexchangerate.DollarExchangeRate, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	resp, err := a.client.GetDollarQuote(ctxTimeout)
	if err != nil {
		fmt.Printf("error calling economiaAPI. [Error: %s]\n", err.Error())
		return nil, err
	}

	return resp.ToDollarExchangeRate()
}
