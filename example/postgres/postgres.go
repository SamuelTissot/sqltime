package main

import (
	"database/sql"
	"fmt"
	"github.com/SamuelTissot/sqltime"
	_ "github.com/lib/pq"
	"reflect"
)

type Memory struct {
	ID      int64
	Moment  sqltime.Time
	Content string
}

func main() {
	// change the dataSourceName value
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	setup(db)

	m := Memory{
		Moment:  sqltime.Now(),
		Content: "When a beer is ravishing, the shabby Hops Alligator Ale throws a nuclear Heineken at a psychotic lager",
	}

	query := `INSERT INTO memories (moment, content)
				VALUES ($1, $2) 
				RETURNING id`

	stmt, err := db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		panic(err)
	}

	err = stmt.QueryRow(m.Moment, m.Content).Scan(&m.ID)
	if err != nil {
		panic(err)
	}

	var fetchMemory Memory
	err = db.QueryRow("SELECT id, moment, content FROM memories WHERE id = $1", m.ID).Scan(&fetchMemory.ID, &fetchMemory.Moment, &fetchMemory.Content)
	if err != nil {
		panic(err)
	}

	if reflect.DeepEqual(m, fetchMemory) {
		fmt.Println("Wow equal")
	} else {
		fmt.Println("NOT EQUAL :/")
		fmt.Println(m)
		fmt.Println(fetchMemory)
	}
}

func setup(db *sql.DB) {
	_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS memories
			(
				id 			SERIAL,
				moment      timestamp,
				content        varchar(255)
			);
		`)

	if err != nil {
		panic(err)
	}
}
