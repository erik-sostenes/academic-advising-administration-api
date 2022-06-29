package repository

import "testing"

// TestNewRedis redis instance unit test
func TestNewRedis(t *testing.T) {
	db, err := LoadRedisConnection()
	if err != nil {
		t.Fatalf("Failed connection to Redis")
	}
	defer db.Close()
}
