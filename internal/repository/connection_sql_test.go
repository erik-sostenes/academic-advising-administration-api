package repository

import "testing"

// TestNewMySQL MySQL instance unit test
func TestNewMySQL(t *testing.T) {
	db, err := LoadMySQLConnection()
	if err != nil {
		t.Fatalf("Failed connection to MySQL")
	}
	defer db.Close()
}
