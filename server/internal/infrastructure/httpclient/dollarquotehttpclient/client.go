package dollarquotehttpclient

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/dprio/currency-exchange/server/internal/domain/dollarquote"
)

const url = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

type DollarQuoteHTTPClient interface {
	GetDollarQuote(ctx context.Context) (*dollarquote.DollarQuote, error)
}

type dollarQuoteHTTPClientImpl struct {
	client *http.Client
}

func New(client *http.Client) DollarQuoteHTTPClient {
	return &dollarQuoteHTTPClientImpl{client: client}
}

func (c *dollarQuoteHTTPClientImpl) GetDollarQuote(ctx context.Context) (*dollarquote.DollarQuote, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var exchangeRateResponse exchangeRateResponse
	if err := json.Unmarshal(body, &exchangeRateResponse); err != nil {
		return nil, err
	}

	return exchangeRateResponse.ToDollarQuote()
}
