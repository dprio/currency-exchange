package economiapi

import (
	"strconv"

	"github.com/dprio/currency-exchange/server/internal/domain/dollarexchangerate"
)

type (
	ExchangeRateResponse struct {
		USDBRL ExchangeRate `json:"USDBRL"`
	}

	ExchangeRate struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	}
)

func (r *ExchangeRateResponse) ToDollarExchangeRate() (*dollarexchangerate.DollarExchangeRate, error) {
	bid, err := strconv.ParseFloat(r.USDBRL.Bid, 64)
	if err != nil {
		return nil, err
	}

	return &dollarexchangerate.DollarExchangeRate{
		Value: float32(bid),
	}, nil
}
