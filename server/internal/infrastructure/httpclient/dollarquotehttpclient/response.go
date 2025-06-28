package dollarquotehttpclient

import (
	"strconv"

	"github.com/dprio/currency-exchange/server/internal/domain/dollarquote"
)

type (
	exchangeRateResponse struct {
		USDBRL exchangeRate `json:"USDBRL"`
	}
	exchangeRate struct {
		Bid string `json:"bid"`
	}
)

func (r *exchangeRateResponse) ToDollarQuote() (*dollarquote.DollarQuote, error) {
	bid, err := strconv.ParseFloat(r.USDBRL.Bid, 64)
	if err != nil {
		return nil, err
	}

	return &dollarquote.DollarQuote{
		Value: float32(bid),
	}, nil
}
