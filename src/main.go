package main


import (
  "database/sql"
  "fmt"
	"log"

	_ "github.com/lib/pq"
)


const (
  host     = "localhost"
  port     = 5432
  user     = "admin"
  password = "admin"
  dbname   = "kamronbek_db"
)


//cmd command
//docker run --name postgres-container -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=kamronbek_db -p 5432:5432 -d postgres


func add_courses() {
  conn_data := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)


	db, err := sql.Open("postgres", conn_data)
	if err != nil { log.Fatal("[ERROR]: failed to connect to the database: ", err) }
	defer db.Close()


	course := []struct {
		name    string
		day     string
    time    string
    room    string
    level   int
	}{
    {"Network Principles", "M", "11:30a 01:50p", "412NH", 3},
    {"Social Engineering", "F", "11:30a 01:50p", "306NH", 4},
    {"Intro to Sustainability", "F", "09:00a 11:20a", "webnet", 9},
    {"Intro to Ethics", "W", "09:00a 11:20a", "305WH", 2},
    {"Data Structures 2", "M", "09:00a 11:20a", "413NH", 1},
    {"Public Speaking", "W", "11:30a 01:50p", "209NH", 3},
	}


	for _, i := range course {
		query := `INSERT INTO timetable (name, day, time, room, level) VALUES ($1, $2, $3, $4, $5)`

		_, err := db.Exec(query, i.name, i.day, i.time, i.room, i.level)
		if err != nil { log.Printf("[ERROR]: Failed to insert timetable %s: %v\n", i.name, err) }
	}
}


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
  add_courses()
}


