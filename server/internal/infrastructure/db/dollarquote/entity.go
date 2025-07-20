package dollarquote

import "time"

type DollarQuoteEntity struct {
	ID        *int64
	Value     float32
	CreatedAt time.Time
}
