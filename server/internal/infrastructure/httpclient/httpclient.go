package httpclient

import (
	"net/http"
	"time"

	"github.com/dprio/currency-exchange/server/internal/infrastructure/httpclient/economiapi"
)

type HTTPClient struct {
	EconomiaAPIHTTPClient economiapi.Client
}

func New() HTTPClient {
	return HTTPClient{
		EconomiaAPIHTTPClient: economiapi.New(newClient(10 * time.Second)),
	}
}

func newClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
	}
}
