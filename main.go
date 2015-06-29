package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	// Parse flags
	kingpin.Parse()

	// Setup database
	dbconnect()

	// Setup router
	route()
}
