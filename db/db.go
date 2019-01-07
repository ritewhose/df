package db

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Quote struct {
	Id       int64
	Creator  string
	Count    int
	Quote    string
	AddTime  time.Time
	LastTime time.Time
}

var db *sql.DB
var count int64 // I may put this in a struct so it's harder to screw up.

func init() {
	var err error
	db, err = sql.Open("sqlite3", "file:db/quotes.db?cache=shared")
	if err != nil {
		panic(err)
	}
	err = db.QueryRow("SELECT COUNT(*) FROM quotes").Scan(&count)
	if err != nil {
		panic(err)
	}
}

func (quote *Quote) Get() error {
	row := db.QueryRow("SELECT count, quote, addtime, lasttime FROM quotes WHERE id = ?", quote.Id)
	err := row.Scan(&quote.Count, &quote.Quote, &quote.AddTime, &quote.LastTime)
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE quotes SET count = count + 1, lasttime = ? WHERE id = ?", time.Now().UTC(), quote.Id)
	if err != nil {
		return err
	}

	return err
}

func (qt *Quote) Add() error {
	res, err := db.Exec("INSERT INTO quotes (creator, quote) VALUES (?, ?)", qt.Creator, qt.Quote)
	if err != nil {
		return err
	}
	qt.Id, _ = res.LastInsertId()
	count++
	return err
}

func QuoteCount() int64 {
	return count.Val()
}

func SafeId(id int64) int64 {
	// Kill Me Baby
	if id < 0 {
		id++
	}
	id += count

	// Zero is NOT safe.
	id--
	id %= count
	id++
	return id
}
