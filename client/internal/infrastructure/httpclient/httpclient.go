package httpclient

import (
	"net/http"
	"time"

	"github.com/dprio/currency-exchange/client/internal/infrastructure/httpclient/dollarquote"
)

type HTTPClient struct {
	DollarQuoteHTTPClient dollarquote.Client
}

func New() HTTPClient {
	return HTTPClient{
		DollarQuoteHTTPClient: dollarquote.New(newClient(10 * time.Second)),
	}
}

func newClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
	}
}
