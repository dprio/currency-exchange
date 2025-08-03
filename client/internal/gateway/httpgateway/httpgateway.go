package httpgateway

import (
	"github.com/dprio/currency-exchange/client/internal/gateway/httpgateway/dollarexchangerateapigateway"
	"github.com/dprio/currency-exchange/client/internal/infrastructure/httpclient"
)

type HTTPGateway struct {
	DollarQuoteAPIGateway dollarexchangerateapigateway.Gateway
}

func New(httpClients httpclient.HTTPClient) HTTPGateway {
	return HTTPGateway{
		DollarQuoteAPIGateway: dollarexchangerateapigateway.New(httpClients.DollarExchangeRateHTTPClient),
	}
}
