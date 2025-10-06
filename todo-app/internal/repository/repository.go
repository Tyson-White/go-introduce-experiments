package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBConnection struct {
	DB *sqlx.DB
}

func ConnectDB(
	user,
	password,
	dbName string,
	dbPort int,
) DBConnection {

	dbconn := DBConnection{}

	connStr := fmt.Sprintf("user=%v password=%v port=%d dbname=%v sslmode=disable", user, password, dbPort, dbName)
	db, err := sqlx.Connect("postgres", connStr)

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	dbconn.DB = db

	log.Println("[INFO] Connected to postgres on port:", dbPort)

	return dbconn
}
