package db

import (
	"fmt"
)

type QuoteDriver interface {
	InsertQuote(*Quote) error
	SelectRandomQuote() (*Quote, error)
	SelectQuoteByNumber(int) (*Quote, error)
}

type Quote struct {
	Creator  string
	QuoteMsg string
}

func (q Quote) String() string {
	return fmt.Sprintf("<@%s> %s", q.Creator, q.QuoteMsg)
}

func (db *DB) InsertQuote(q *Quote) error {
	statement, err := db.Prepare("INSERT INTO Quote(creator, quoteMsg) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("[InsertQuote] %s", err)
	}
	defer statement.Close()

	_, err = statement.Exec(q.Creator, q.QuoteMsg)
	if err != nil {
		return fmt.Errorf("[InsertQuote] %s", err)
	}

	return nil
}

func (db *DB) SelectRandomQuote() (*Quote, error) {
	q := &Quote{}
	row := db.QueryRow("SELECT * FROM Quote ORDER BY RANDOM() LIMIT 1")

	err := row.Scan(&q)
	if err != nil {
		return nil, fmt.Errorf("[SelectRandomQuote] %s", err)
	}

	return q, nil
}

func (db *DB) SelectQuoteByNumber(number int) (*Quote, error) {
	q := &Quote{}
	row := db.QueryRow("SELECT * FROM Quote LIMIT 1 OFFSET $1", number)

	err := row.Scan(&q)
	if err != nil {
		return nil, fmt.Errorf("[SelectQuoteByNumber] %s", err)
	}

	return q, nil
}
