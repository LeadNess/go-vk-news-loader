package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	Conn *sql.DB
}

func OpenConnection(user, password, host, port, dbName string) (*Storage, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	return &Storage{Conn: db}, err
}
