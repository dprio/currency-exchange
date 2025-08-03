package httpclient

import (
	"net/http"
	"time"

	"github.com/dprio/currency-exchange/client/internal/infrastructure/httpclient/dollarexchangerate"
)

type HTTPClient struct {
	DollarExchangeRateHTTPClient dollarexchangerate.Client
}

func New() HTTPClient {
	return HTTPClient{
		DollarExchangeRateHTTPClient: dollarexchangerate.New(newClient(10 * time.Second)),
	}
}

func newClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
	}
}
