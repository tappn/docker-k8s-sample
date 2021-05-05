package db

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Connect is connection to database
func Connect() *sqlx.DB {
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	name := os.Getenv("DATABASE_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		user, pass, host, name)

	for {
		db, err := sqlx.Connect("mysql", dsn)
		if err == nil {
			return db
		}
		log.Println(err)
		time.Sleep(1 * time.Second)
	}
}
