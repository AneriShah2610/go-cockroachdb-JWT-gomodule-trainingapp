package driver

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type DB struct {
	DatabaseConn *sql.DB
}

var dbConn = &DB{}

func connectDb() (*DB, error) {
	connectionString := "postgresql://root@localhost:26257/traning?sslmode=disable"
	dbConnection, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error while connecting database", err)
	}
	dbConn.DatabaseConn = dbConnection
	return dbConn, err
}
