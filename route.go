package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"fmt"
	"github.com/lukevers/converse/routes"
)

var router *httprouter.Router

func route() {
	// Create new router
	router = httprouter.New()

	router.GET("/login", routes.Login)

	// Addr
	addr := fmt.Sprintf("%s:%d", *host, *port)

	// Log
	log.Fatal(http.ListenAndServe(addr, router))
}
