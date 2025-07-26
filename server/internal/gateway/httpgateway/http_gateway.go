package httpgateway

import (
	"github.com/dprio/currency-exchange/server/internal/gateway/httpgateway/economiaapiagateway"
	"github.com/dprio/currency-exchange/server/internal/infrastructure/httpclient"
)

type HTTPGateways struct {
	EconomiaAPIAdapter economiaapiagateway.Adapter
}

func New(httpClient httpclient.HTTPClient) HTTPGateways {
	return HTTPGateways{
		EconomiaAPIAdapter: economiaapiagateway.New(httpClient.EconomiaAPIHTTPClient),
	}
}
