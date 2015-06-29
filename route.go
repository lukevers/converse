package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
)

var router *httprouter.Router

func route() {
	router = httprouter.New()

	

	// Log
	log.Fatal(http.ListenAndServe(addr, router))
}
