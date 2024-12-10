package main


import (
	"html/template"
	"database/sql"
	"net/http"
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


func launch_webpage(w http.ResponseWriter, r *http.Request) {
	conn_data := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)


	db, err := sql.Open("postgres", conn_data)
	if err != nil { log.Fatal("[ERROR]: failed to connect to the database: ", err) }
	defer db.Close()


	type Course struct {
		Name    string
		Day     string
		Time    string
		Room    string
		Level   int
	}

	rows, err := db.Query("SELECT name, day, time, room, level FROM timetable")
	if err != nil { log.Fatal("[ERROR]: 49: ", err) }
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var c Course
		if err := rows.Scan(&c.Name, &c.Day, &c.Time, &c.Room, &c.Level); err != nil { log.Fatal("[ERROR]: 55: ", err) }
		courses = append(courses, c)
	}

	if err := rows.Err(); err != nil { log.Fatal("[ERROR]: 59: ", err) }

	// Parse and execute the HTML template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, courses)
}


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
		{"Network Principles", "M", "11:30a 01:50p", "412NH", 2670},
		{"Social Engineering", "F", "11:30a 01:50p", "306NH", 2710},
		{"Intro to Sustainability", "F", "09:00a 11:20a", "webnet", 1000},
		{"Intro to Ethics", "W", "09:00a 11:20a", "305WH", 2110},
		{"Data Structures 2", "M", "09:00a 11:20a", "413NH", 3100},
		{"Public Speaking", "W", "11:30a 01:50p", "209NH", 1040},
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

	http.HandleFunc("/", launch_webpage)
	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


