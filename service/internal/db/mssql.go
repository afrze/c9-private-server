package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

func InitMSSQL(conn string) *sql.DB {
	db, err := sql.Open("sqlserver", conn)
	if err != nil {
		panic(err)
	}

	// pool settings - adjust to preference
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		panic(err)
	}

	log.Println("[db] Connected to MSSQL")

	return db
}
