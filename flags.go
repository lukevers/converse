package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	database   = kingpin.Flag("database", "SQLite3 database resource").Default("app/database/converse.db").String()
	host       = kingpin.Flag("host", "Host to bind webserver to").Default("127.0.0.1").IP()
	port       = kingpin.Flag("port", "Port to bind webserver to").Default("3040").Int()
	production = kingpin.Flag("production", "Run in production mode").Default("false").Bool()
)
