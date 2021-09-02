package logger

import (
	"database/sql"
	"log"
	"time"

	clickhouse "github.com/ClickHouse/clickhouse-go"
)

var DB *sql.DB

type Logger struct {
	DB *sql.DB
}

func Connect() *sql.DB {
	connect, err := sql.Open("clickhouse", "http://localhost:9000/default?debug=true")
	if err != nil {
		log.Println(err)
	}
	if err := connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			log.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			log.Println(err)
		}
		return nil
	}
	_, err = connect.Exec(`
		CREATE TABLE IF NOT EXISTS logs (
			id 			 UUID,
			message      Varchar,
			date  		 DateTime(3, 'Europe/Moscow')
		) engine=TinyLog
	`)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Logger has been started")
	DB = connect

	return connect
}

func (l *Logger) Log(message string) {

	log.Println(message)

	var (
		tx, _   = DB.Begin()
		stmt, _ = tx.Prepare("INSERT INTO logs (message, date) VALUES (?, ?)")
	)
	defer stmt.Close()
	if _, err := stmt.Exec(
		message,
		time.Now(),
	); err != nil {
		log.Println(err.Error())
	}
	if err := tx.Commit(); err != nil {
		log.Println(err.Error())
	}
}
