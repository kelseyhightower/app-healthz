package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kelseyhightower/app-healthz/healthz"
)

var version = "1.0.0"

func main() {
	log.Println("Starting app...")

	httpAddr := os.Getenv("HTTP_ADDR")
	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseHost := os.Getenv("DATABASE_HOST")
	databaseName := os.Getenv("DATABASE_NAME")

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Initializing database connection pool...")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		databaseUsername, databasePassword, databaseHost, databaseName)

	hc := &healthz.Config{
		Hostname: hostname,
		Database: healthz.DatabaseConfig{
			DriverName:     "mysql",
			DataSourceName: dataSourceName,
		},
	}

	healthzHandler, err := healthz.Handler(hc)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/healthz", healthzHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, html, hostname, version)
	})

	log.Printf("HTTP service listening on %s", httpAddr)
	http.ListenAndServe(httpAddr, nil)
}
