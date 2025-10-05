package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBConnection struct {
	DB *sql.DB
}

func ConnectDB(
	user,
	password,
	dbName string,
) DBConnection {

	dbconn := DBConnection{}

	connStr := fmt.Sprintf("user=%v password=%v port=5433 dbname=%v sslmode=disable", user, password, dbName)
	fmt.Println(connStr)
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

	return dbconn
}
