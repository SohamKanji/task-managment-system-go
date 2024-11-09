package main

import (
	"database/sql"
	"os"

	"github.com/SohamKanji/task-management-system-go/api"
	db "github.com/SohamKanji/task-management-system-go/db/sqlc"
	_ "github.com/lib/pq"
)

func main() {
	psqldb, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/task-managment-system-db?sslmode=disable")
	if err != nil {
		os.Exit(1)
	}
	defer psqldb.Close()
	store := db.NewStore(psqldb)
	server := api.NewServer(store)
	server.Start(":8080")
}
