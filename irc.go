package main

import (
	"log"
)

var clients []*Client

func setupirc() {
	// Find all clients
	rows, err := db.Queryx("SELECT * FROM clients")
	if err != nil {
		log.Println(err)
	}

	// For each client we want to add it to our slice of clients
	for rows.Next() {
		var client Client
		if err := rows.StructScan(&client); err != nil {
			log.Println(err)
		}

		clients = append(clients, &client)
	}

	// Start each client
	for _, client := range clients {
		go client.Connect()
	}
}
