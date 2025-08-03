package dollarexchangeratedb

import (
	"time"

	"github.com/dprio/currency-exchange/server/internal/domain/dollarexchangerate"
)

type DollarExchangeRateEntity struct {
	ID        *int64
	Value     float32
	CreatedAt time.Time
}

func NewDollarQuoteEntity(exchangeRate dollarexchangerate.DollarExchangeRate) DollarExchangeRateEntity {
	return DollarExchangeRateEntity{
		Value:     exchangeRate.Value,
		CreatedAt: time.Now(),
	}
}

func (e *DollarExchangeRateEntity) ToDollarExchnageRate() *dollarexchangerate.DollarExchangeRate {
	return &dollarexchangerate.DollarExchangeRate{
		Value: e.Value,
	}
}
