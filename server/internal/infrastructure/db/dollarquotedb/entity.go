package dollarquotedb

import (
	"time"

	"github.com/dprio/currency-exchange/server/internal/domain/dollarquote"
)

type DollarQuoteEntity struct {
	ID        *int64
	Value     float32
	CreatedAt time.Time
}

func NewDollarQuoteEntity(dollarQuote dollarquote.DollarQuote) DollarQuoteEntity {
	return DollarQuoteEntity{
		Value:     dollarQuote.Value,
		CreatedAt: time.Now(),
	}
}

func (e *DollarQuoteEntity) ToDollarQuote() *dollarquote.DollarQuote {
	return &dollarquote.DollarQuote{
		Value: e.Value,
	}
}
