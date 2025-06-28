package infrastructure

import (
	"net/http"
	"time"

	"github.com/dprio/currency-exchange/server/internal/infrastructure/httpclient/dollarquotehttpclient"
)

type HTTPClient struct {
	DollarQuoteHTTPClient dollarquotehttpclient.DollarQuoteHTTPClient
}

func New() HTTPClient {
	return HTTPClient{
		DollarQuoteHTTPClient: dollarquotehttpclient.New(newClient(10 * time.Second)),
	}
}

func newClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
	}
}
