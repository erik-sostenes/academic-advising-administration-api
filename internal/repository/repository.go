package repository

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
)
// TODO: add comments
type Type uint 

const (
	SQL Type = iota
	NoSQL
)

type Configuration struct {
	Type Type
	SQL *sql.DB
	NoSQL *redis.Client
}
