package main

import (
	"os"
)

// Config carries the api configuration
type Config struct {
	Version        string
	MongoDBConnStr string
	MongoDBName    string
}

func config(version string) Config {
	c := Config{
		Version:        version,
		MongoDBConnStr: os.Getenv("MONGO_CONN_STR"),
		MongoDBName:    os.Getenv("MONGO_DB_NAME"),
	}
	return c
}
