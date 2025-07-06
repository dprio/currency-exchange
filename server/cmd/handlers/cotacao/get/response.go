package get

import "github.com/dprio/currency-exchange/server/internal/domain/dollarquote"

type response struct {
	DollarQuote float32 `json:"bid"`
}

func newResponse(dollarQuote *dollarquote.DollarQuote) *response {
	return &response{
		DollarQuote: dollarQuote.Value,
	}
}
