package database

import (
	"database/sql"
	"os"
	"time"

	"github.com/cyan-store/hook/log"
	_ "github.com/go-sql-driver/mysql"
)

var Conn *sql.DB
var DefaultTimeout = 3 * time.Second

func OpenDB(dsn string) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Error.Println("[OpenDB] Could not open DB", err)
		os.Exit(1)
	}

	if err = db.Ping(); err != nil {
		log.Error.Println("[OpenDB] Could not ping database", err)
		os.Exit(1)
	}

	Conn = db
	log.Info.Println("[OpenDB] Connected to database.")
}
