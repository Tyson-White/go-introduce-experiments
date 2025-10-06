package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DBConnection struct {
	DB *sql.DB
}

func ConnectDB(
	user,
	password,
	dbName string,
	dbPort int,
) DBConnection {

	dbconn := DBConnection{}

	connStr := fmt.Sprintf("user=%v password=%v port=%d dbname=%v sslmode=disable", user, password, dbPort, dbName)
	db, err := sql.Open("postgres", connStr)

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS todos (id serial not null unique, title varchar(255) not null, text text, time timestamp default now())")

	if err != nil {
		panic(err)
	}

	dbconn.DB = db

	log.Println("[INFO] Connected to postgres on port:", dbPort)

	return dbconn
}
