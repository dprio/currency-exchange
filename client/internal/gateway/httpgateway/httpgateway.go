package httpgateway

import (
	"github.com/dprio/currency-exchange/client/internal/gateway/httpgateway/dollarquoteapigateway"
	"github.com/dprio/currency-exchange/client/internal/infrastructure/httpclient"
)

type HTTPGateway struct {
	DollarQuoteAPIGateway dollarquoteapigateway.Gateway
}

func New(httpClients httpclient.HTTPClient) HTTPGateway {
	return HTTPGateway{
		DollarQuoteAPIGateway: dollarquoteapigateway.New(httpClients.DollarQuoteHTTPClient),
	}
}
