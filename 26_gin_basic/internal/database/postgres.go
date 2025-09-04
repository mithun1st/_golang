package database

import (
	"database/sql"
	"fmt"
	"gin-basic/internal/config"

	_ "github.com/lib/pq"
)

type DB struct {
	SqlDB  *sql.DB
	SqlDB1 *sql.DB
}

func NewPostgresDb(dbConfig config.AppDBModel) (*DB, error) {

	var dataSrcName string = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbConfig.DbUser, dbConfig.DbPass, dbConfig.DbHost, dbConfig.DbPort, dbConfig.DbName)

	postgres, err := sql.Open("postgres", dataSrcName)
	if err != nil {
		return nil, err
	}
	return &DB{SqlDB: postgres}, nil

}

func Check(db *DB) error {
	err := db.SqlDB.Ping()
	return err
}

func Close(db *DB) error {
	err := db.SqlDB.Close()
	return err
}
