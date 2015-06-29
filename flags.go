package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	host     = kingpin.Flag("host", "Host to bind webserver to").Default("127.0.0.1").IP()
	port     = kingpin.Flag("port", "Port to bind webserver to").Default("3040").Int()
	database = kingpin.Flag("database", "SQLite3 database resource").Default("app/database/converse.db").String()
)
