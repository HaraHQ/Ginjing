package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type MigrateTable struct {
	Name  string
	Query string
}

func Migrate() {
	migrates := []MigrateTable{
		{
			Name: "Users",
			Query: `CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);`,
		},
		// make changes below
	}

	token := os.Getenv("TURSO_TOKEN")
	url := os.Getenv("TURSO_URL") + token

	// Open database connection
	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Println("Error opening database connection:", err)
		return
	}
	defer db.Close()

	for _, m := range migrates {
		_, err = db.Exec(m.Query)
		if err != nil {
			fmt.Println("Error creating table:", err)
			return
		}

		fmt.Printf("Table %s created successfully.\n\n", m.Name)
	}
}
