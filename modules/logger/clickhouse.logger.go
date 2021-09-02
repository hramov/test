package logger

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	clickhouse "github.com/ClickHouse/clickhouse-go"
)

type Logger struct {
	DB *sql.DB
}

func Connect() *sql.DB {
	connect, err := sql.Open("clickhouse", "tcp://127.0.0.1:9000?debug=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
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

	return connect
}

func (l *Logger) Log(message string) {
	rows, err := l.DB.Query(fmt.Sprintf("INSERT INTO logs (message, date) VALUES (%s, %v)", message, time.Now()))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
	}
}
