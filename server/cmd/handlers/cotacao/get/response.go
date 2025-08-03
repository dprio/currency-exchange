package get

import "github.com/dprio/currency-exchange/server/internal/domain/dollarexchangerate"

type response struct {
	DollarQuote float32 `json:"bid"`
}

func newResponse(exchangeRate *dollarexchangerate.DollarExchangeRate) *response {
	return &response{
		DollarQuote: exchangeRate.Value,
	}
}
