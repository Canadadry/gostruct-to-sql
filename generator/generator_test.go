package generator

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
	"time"
)

const (
	dbProtocol = "sqlite3"
	dbUrl      = "test.db"
)

func TestIsUpToDateEmptyDB(t *testing.T) {
	db, err := sql.Open(dbProtocol, dbUrl)
	if err != nil {
		t.Fatalf("cannot open %s : %v", dbUrl, err)
	}
	defer db.Close()
	defer os.Remove(dbUrl)

	g := Generator{}
	err = g.RegisterType(struct{ test int }{})
	result := g.IsUpToDate(db)
	if result != false {
		t.Fatalf("should have said db is not sync")
	}
}

func TestIsUpToDate(t *testing.T) {
	db, err := sql.Open(dbProtocol, dbUrl)
	if err != nil {
		t.Fatalf("cannot open %s : %v", dbUrl, err)
	}
	defer db.Close()
	defer os.Remove(dbUrl)

	g := Generator{}
	err = g.RegisterType(struct {
		test        int
		creation    time.Time
		name        string
		uuid        string `type:"char" size:"36"`
		description string `type:"varchar" size:"500"`
	}{})
	if err != nil {
		t.Fatalf("Cannot generate database schema : %v", err)
	}

	query, err := g.Generate()
	if err != nil {
		t.Fatalf("Cannot generate database schema : %v", err)
	}

	_, err = db.Exec(query)
	if err != nil {
		t.Fatalf("Error while creating schema : %v", err)
	}

	result := g.IsUpToDate(db)
	if result != true {
		t.Fatalf("should have said db is sync with schema %s", query)
	}
}

func TestGenerate(t *testing.T) {
	g := Generator{}
	err := g.RegisterType(struct{ test int }{})
	if err != nil {
		t.Fatalf("Failed : %v", err)
	}
	_, err = g.Generate()
	if err != nil {
		t.Fatalf("Failed : %v", err)
	}
}
