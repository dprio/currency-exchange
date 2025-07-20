package dollarquote

import "database/sql"

type Client interface {
	SaveDollarQuote(entity DollarQuoteEntity) (*DollarQuoteEntity, error)
	FindDollarQuoteByID(id int64) (*DollarQuoteEntity, error)
}

type client struct {
	db *sql.DB
}

func NewClient(db *sql.DB) Client {
	return &client{db: db}
}

func (c *client) SaveDollarQuote(entity DollarQuoteEntity) (*DollarQuoteEntity, error) {
	query := "INSERT INTO dollar_quotes (value, created_at) VALUES (?, ?)"
	result, err := c.db.Exec(query, entity.Value, entity.CreatedAt)
	if err != nil {
		return nil, err
	}

	resp := entity
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	resp.ID = &id

	return &resp, nil
}

func (c *client) FindDollarQuoteByID(id int64) (*DollarQuoteEntity, error) {
	query := "SELECT id, value, created_at FROM dollar_quotes WHERE id = ?"
	stmt, err := c.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var resp DollarQuoteEntity
	err = row.Scan(&resp.ID, &resp.Value, &resp.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
