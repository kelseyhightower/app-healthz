package healthz

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	Hostname string
	Database DatabaseConfig
}

type DatabaseConfig struct {
	DriverName     string
	DataSourceName string
}

type handler struct {
	dc       *DatabaseChecker
	hostname string
	metadata map[string]string
}

func Handler(hc *Config) (http.Handler, error) {
	mc, err := NewDatabaseChecker(hc.Database.DriverName, hc.Database.DataSourceName)
	if err != nil {
		return nil, err
	}

	config, err := mysql.ParseDSN(hc.Database.DataSourceName)
	if err != nil {
		return nil, err
	}

	metadata := make(map[string]string)
	metadata["database_url"] = config.Addr
	metadata["database_user"] = config.User

	h := &handler{dc, hc.Hostname, metadata}
	return h, nil
}

type Response struct {
	Hostname string            `json:"hostname"`
	Metadata map[string]string `json:"metadata"`
	Errors   []Error           `json:"errors"`
}

type Error struct {
	Description string            `json:"description"`
	Error       string            `json:"error"`
	Metadata    map[string]string `json:"metadata"`
	Type        string            `json:"type"`
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Hostname: h.hostname,
		Metadata: h.metadata,
	}

	statusCode := http.StatusOK
	errors := make([]Error, 0)

	err = h.dc.Ping()
	if err != nil {
		errors = append(errors, Error{
			Type:        "DatabasePing",
			Description: "Database health check.",
			Error:       err.Error(),
		})
	}

	response.Errors = errors
	if len(response.Errors) > 0 {
		statusCode = http.StatusInternalServerError
		for _, e := range response.Errors {
			log.Println(e.Error)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	data, err := json.MarshalIndent(&response, "", "  ")
	if err != nil {
		log.Println(err)
	}
	w.Write(data)
}
