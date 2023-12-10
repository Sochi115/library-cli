package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteDb struct {
	db *sqlx.DB
}

func ConnectToSqliteDb(datasource string) SqliteDb {
	db, err := sqlx.Connect("sqlite3", datasource)
	if err != nil {
		log.Fatal(err)
	}

	return SqliteDb{
		db: db,
	}
}

func (sqlite *SqliteDb) InitBookTable() {
	query := `
  CREATE TABLE IF NOT EXISTS bookEntries(
    key TEXT PRIMARY KEY,
    title TEXT,
    authors TEXT,
    isbn_13 TEXT,
    isbn_10 TEXT,
    publishers TEXT,
    publish_date TEXT,
    number_of_pages INT,
    rating INT
  );`

	_, err := sqlite.db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func (sqlite *SqliteDb) FetchAll() {
	query := "SELECT * FROM bookEntries"

	rows, err := sqlite.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		fmt.Println(rows)
	}
}

func (sqlite *SqliteDb) AddBook(bookEntry BookEntry) {
	if sqlite.IsExistsBookEntry(bookEntry.Key) {
		fmt.Println("Entry already saved")
		return
	}
	query := `INSERT OR IGNORE INTO bookEntries(key, title, authors, isbn_13, isbn_10, publishers, publish_date, number_of_pages, rating)
  VALUES (:key, :title, :authors, :isbn_13, :isbn_10, :publishers, :publish_date, :number_of_pages, :rating);`

	_, err := sqlite.db.NamedExec(query, bookEntry)
	if err != nil {
		log.Fatal(err)
	}
}

func (sqlite *SqliteDb) IsExistsBookEntry(key string) bool {
	var count int
	query := `SELECT COUNT(*) FROM bookEntries WHERE key = ?`

	err := sqlite.db.QueryRow(query, key).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count != 0
}

func (sqlite *SqliteDb) CloseDb() {
	sqlite.db.Close()
}
