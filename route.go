package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lukevers/converse/routes"
	"log"
	"net/http"
	"time"
)

var router *gin.Engine

func route() {
	// If we're in production mode don't run gin in develop
	if *production {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create new router
	router = gin.Default()

	// Compile html templates
	router.LoadHTMLGlob("app/html/*.html")

	// Add all routes
	addRoutes()

	// Add static routes
	router.Static("/", "./public/")

	// Create http server based off of net/http
	addr := fmt.Sprintf("%s:%d", *host, *port)
	s := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Log
	log.Fatal(s.ListenAndServe())
}

func addRoutes() {
	router.GET("/init", routes.Init)
	router.GET("/", routes.Index)
}
