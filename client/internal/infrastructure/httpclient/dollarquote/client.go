package dollarquote

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url = "http://localhost:8080/cotacao"

type Client interface {
	GetDollarQuote(ctx context.Context) (*DollarQuoteResponse, error)
}

type dollarQuoteHTTPClient struct {
	client *http.Client
}

func New(client *http.Client) Client {
	return &dollarQuoteHTTPClient{client: client}
}

func (c *dollarQuoteHTTPClient) GetDollarQuote(ctx context.Context) (*DollarQuoteResponse, error) {
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

	var dollarQuoreResponse DollarQuoteResponse
	if err := json.Unmarshal(body, &dollarQuoreResponse); err != nil {
		return nil, err
	}

	return &dollarQuoreResponse, err
}
