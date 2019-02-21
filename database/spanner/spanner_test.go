package spanner

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"os"
	"testing"

	dt "github.com/golang-migrate/migrate/v4/database/testing"
	_ "github.com/golang-migrate/migrate/v4/source/stub"
)

func Test(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	db, ok := os.LookupEnv("SPANNER_DATABASE")
	if !ok {
		t.Skip("SPANNER_DATABASE not set, skipping test.")
	}

	s := &Spanner{}
	addr := fmt.Sprintf("spanner://%v", db)
	d, err := s.Open(addr)
	if err != nil {
		t.Fatalf("%v", err)
	}
	dt.Test(t, d, []byte("SELECT 1"))
}

func TestMigrate(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	db, ok := os.LookupEnv("SPANNER_DATABASE")
	if !ok {
		t.Skip("SPANNER_DATABASE not set, skipping test.")
	}

	s := &Spanner{}
	addr := fmt.Sprintf("spanner://%v", db)
	d, err := s.Open(addr)
	if err != nil {
		t.Fatalf("%v", err)
	}
	src, err := dt.GetStubSource()
	if err != nil {
		t.Fatalf("%v", err)
	}
	m, err := migrate.NewWithInstance("stub", src, "", d)
	if err != nil {
		t.Fatalf("%v", err)
	}
	dt.TestMigrate(t, m, []byte("SELECT 1"))
}
