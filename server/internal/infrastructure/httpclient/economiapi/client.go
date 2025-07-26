package economiapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url = "https://economia.awesomeapi.com.br/json/last/USD-BRL?token=1fec47184a95b44f47f6ba765947c3914734b5a2ec66bbe759f83061b5781263"

type Client interface {
	GetDollarQuote(ctx context.Context) (*ExchangeRateResponse, error)
}

type dollarQuoteHTTPClient struct {
	client *http.Client
}

func New(client *http.Client) Client {
	return &dollarQuoteHTTPClient{client: client}
}

func (c *dollarQuoteHTTPClient) GetDollarQuote(ctx context.Context) (*ExchangeRateResponse, error) {
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

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error calling economi-api. Error: [%s]", string(body))
	}

	fmt.Println(string(body))

	var exchangeRateResponse ExchangeRateResponse
	if err := json.Unmarshal(body, &exchangeRateResponse); err != nil {
		return nil, err
	}

	return &exchangeRateResponse, err
}
