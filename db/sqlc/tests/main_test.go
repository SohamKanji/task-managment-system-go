package db

import (
	"database/sql"
	"os"
	"testing"

	db "github.com/SohamKanji/task-management-system-go/db/sqlc"
	_ "github.com/lib/pq"
)

var testQueries *db.Queries

func TestMain(m *testing.M) {
	testDB, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/task-managment-system-db?sslmode=disable")
	if err != nil {
		os.Exit(1)
	}
	testQueries = db.New(testDB)
	os.Exit(m.Run())
}
