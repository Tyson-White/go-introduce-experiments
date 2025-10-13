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

const schema = `
	CREATE TABLE IF NOT EXISTS category (
		id serial not null unique,
		name varchar(100) unique not null
	);

	CREATE TABLE IF NOT EXISTS todos (
		id serial PRIMARY KEY not null unique, 
		title varchar(255) not null, 
		text text, 
		time timestamp default now(), 
		completed bool default false,
		category_id int REFERENCES category(id) ON DELETE SET NULL
	)`

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

	_, err = db.Exec(schema)

	if err != nil {
		panic(err)
	}

	log.Println("[INFO] Connected to postgres on port:", dbPort)

	return dbconn
}
