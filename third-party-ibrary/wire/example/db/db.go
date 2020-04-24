package db

import "database/sql"

type Config struct {

}

func NewDB(c *Config) (*sql.DB, error) {
	return &sql.DB{}, nil
}
