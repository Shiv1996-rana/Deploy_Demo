package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// connStr := "host=localhost port=5432 user=postgres password=postgres dbname=multitenant_db sslmode=disable "
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER_NAME")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("SSL_MODE")
	Port, err := strconv.Atoi(port)
	if err != nil {
		log.Printf("no converted string to integer %v", err)
	}
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, Port, user, password, dbname, sslmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("failed connection Driver %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetConnMaxIdleTime(5)
	db.SetConnMaxIdleTime(30 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	if err = db.Ping(); err != nil {
		// log.Printf("failed connection database %v", err)
		log.Fatalf("failed connection database: %v", err)
	}
	DB = db
	fmt.Println("database connection successfully....")
	CreateTable()

}

func CreateTable() {
	createTable := `
	 CREATE TABLE IF NOT EXISTS crud(
	 id SERIAL PRIMARY KEY,
	 name VARCHAR(200) NOT NULL,
	 email VARCHAR(100) UNIQUE NOT NULL,
	 mobile_no BIGINT ,
     address JSONB NOT NULL
	 )
	`
	_, err := DB.Exec(createTable)
	if err != nil {
		log.Printf("no created crud Table %v", err)
	}
	fmt.Println("crud Table created successfully..")
}
