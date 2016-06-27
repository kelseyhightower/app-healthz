package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/kelseyhightower/app-healthz/healthz"
)

var (
	dataSourceName string
	databaseName   string
	healthAddr     string
	tables         string
	vaultAddr      string
)

func main() {
	dataSourceName := os.Getenv("DATA_SOURCE_NAME")
	databaseName := os.Getenv("DATABASE_NAME")
	healthAddr := os.Getenv("HEALTH_ADDR")
	tables := os.Getenv("DATABASE_TABLES")
	vaultAddr := os.Getenv("VAULT_ADDRESS")

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	hc := &healthz.Config{
		Hostname: hostname,
		Database: healthz.DatabaseConfig{
			DriverName:     "mysql",
			DataSourceName: dataSourceName,
			DatabaseName:   databaseName,
			Tables:         strings.Split(tables, ","),
		},
		Vault: healthz.VaultConfig{
			Address: vaultAddr,
		},
	}

	healthzHandler, err := healthz.Handler(hc)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/healthz", healthzHandler)
	http.ListenAndServe(healthAddr, nil)
}
