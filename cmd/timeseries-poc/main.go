package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/brunodrugowick/go-timeseries-poc/cmd/timeseries-poc/internal/config"
	config_reader "github.com/brunodrugowick/go-timeseries-poc/pkg/config-reader"
	"github.com/brunodrugowick/go-timeseries-poc/pkg/infrastructure/database"
	"github.com/brunodrugowick/go-timeseries-poc/pkg/server"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {

	props := readProperties()

	// TODO temp (replace with service layer)
	// THIS IS JUST TEST CODE
	ctx := context.Background()
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		props.Database.Username,
		props.Database.Password,
		props.Database.Name)
	db, err := sql.Open(props.Database.Driver, connectionString)
	if err != nil {
		panic(err)
	}

	queries := database.New(db)

	srv := server.NewDefaultServerBuilder().
		SetPort(props.Server.Port).

		// TODO temp (replace with handlers
		WithHandlerFunc("/", func(resp http.ResponseWriter, req *http.Request) {
			name := req.URL.Query().Get("name")
			if len(name) == 0 {
				t, _ := template.New("index").Parse("... try adding a parameter the URL, like ?name=John")
				data := struct {
					Path string
				}{
					Path: req.URL.Path,
				}

				_ = t.Execute(resp, data)

				return
			}
			response := []struct {
				Hello string `json:"hello"`
			}{
				{
					Hello: name,
				},
			}

			resp.Header().Set("Content-Type", "application/json")
			resp.WriteHeader(http.StatusOK)
			err := json.NewEncoder(resp).Encode(response)
			if err != nil {
				resp.WriteHeader(http.StatusInternalServerError)
			}
		}).
		WithHandlerFunc("/measurements", func(resp http.ResponseWriter, req *http.Request) {
			if req.Method == "GET" {
				measurements, err := queries.ListMeasurements(ctx)
				if err != nil {
					panic(err)
				}

				resp.Header().Set("Content-Type", "application/json")
				resp.WriteHeader(http.StatusOK)

				err = json.NewEncoder(resp).Encode(measurements)
				if err != nil {
					resp.WriteHeader(http.StatusInternalServerError)
				}
			} else if req.Method == "POST" {
				var meas database.CreateMeasurementParams
				err := json.NewDecoder(req.Body).Decode(&meas)
				if err != nil {
					return
				}
				createMeasurement, err := queries.CreateMeasurement(ctx, meas)
				if err != nil {
					return
				}
				resp.Header().Set("Content-Type", "application/json")
				resp.WriteHeader(http.StatusCreated)

				err = json.NewEncoder(resp).Encode(createMeasurement)
				if err != nil {
					resp.WriteHeader(http.StatusInternalServerError)
				}
			}
		}).
		Build()

	srv.Run()
}

func readProperties() config.Properties {
	const configLocationEnvVar = "CONFIG"
	configLocation, ok := os.LookupEnv(configLocationEnvVar)

	var reader config_reader.ConfigReader
	if ok {
		log.Printf("Config location found in environment variable CONFIG=%s", configLocation)
		reader = config_reader.ConfigReader{
			File:        configLocation,
			Environment: true,
		}
	} else {
		reader = config_reader.DefaultConfigReader()
	}

	var props config.Properties
	err := reader.Read(&props)
	if err != nil {
		log.Printf("Could not read properties")
		return config.Properties{}
	}
	return props
}
