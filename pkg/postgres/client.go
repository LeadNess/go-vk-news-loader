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
		fmt.Sprintf("%s:%s@%s:%s/%s", user, password, host, port, dbName))
	return &Storage{Conn: db}, err
}
