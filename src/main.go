package main


import (
  "database/sql"
  "fmt"
	"log"

	_ "github.com/lib/pq"
)


const (
  host     = "localhost"
  port     = 8080
  user     = "admin"
  password = "1234"
  dbname   = "kamronbek_db"
)


func create_table() {
  conn_data := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)


	db, err := sql.Open("postgres", conn_data)
	if err != nil { log.Fatal("[ERROR]: failed to connect to the database: ", err) }
	defer db.Close()


	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS timetable (
		course_id SERIAL PRIMARY KEY,
		name TEXT,
    day TEXT,
    time TEXT,
    room TEXT,
		level INT
	)`)
	if err != nil { log.Fatal("[ERROR]: Failed to create table: ", err) }
}


func main() {
  create_table()
}


