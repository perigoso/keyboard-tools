package main

import (
	"kle-app/web"
	"log"
	"os"
)

func main() {
	// CORS is enabled only in prod profile
	prod := os.Getenv("profile") == "prod"
	app := web.NewApp(prod)
	log.Fatal("Error", app.Serve())
}

