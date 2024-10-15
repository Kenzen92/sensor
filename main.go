package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connect to SQLite database (placeholder.db is the database file)
	db, err := sql.Open("sqlite3", "./placeholder.db")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Create a simple table (if it doesn't exist) - placeholder SQL statement
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS example (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		log.Fatal("Failed to prepare table creation statement:", err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatal("Failed to execute table creation statement:", err)
	}

	// Insert a row into the table (this is just for demonstration)
	statement, err = db.Prepare("INSERT INTO example (name) VALUES (?)")
	if err != nil {
		log.Fatal("Failed to prepare insert statement:", err)
	}
	_, err = statement.Exec("Sample Name")
	if err != nil {
		log.Fatal("Failed to execute insert statement:", err)
	}

	// Define the handler for /helloworld route
	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		var name string

		// Query the database for a name
		err := db.QueryRow("SELECT name FROM example WHERE id = 1").Scan(&name)
		if err != nil {
			fmt.Fprintf(w, "Error querying the database: %v", err)
			return
		}

		// Respond with the queried name
		fmt.Fprintf(w, "Hello, %s!", name)
	})

	// Start the HTTP server
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
