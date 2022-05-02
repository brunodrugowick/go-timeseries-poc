package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/brunodrugowick/go-timeseries-poc/pkg/infrastructure/database"
	"github.com/brunodrugowick/go-timeseries-poc/pkg/server"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	srv := server.NewDefaultServerBuilder().
		SetPort(8080).

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
		Build()

	// TODO temp (replace with service layer)
	// THIS IS JUST TEST CODE
	ctx := context.Background()
	db, err := sql.Open("postgres", "user=postgres password=password dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	queries := database.New(db)

	// list all
	measurements, err := queries.ListMeasurements(ctx)
	if err != nil {
		panic(err)
	}
	log.Println(measurements)

	insertedMeasurement, err := queries.CreateMeasurement(ctx, database.CreateMeasurementParams{
		CreatedDate: sql.NullInt64{
			Int64: time.Now().Unix(),
			Valid: true,
		},
		HeartRate: sql.NullInt32{
			Int32: 123,
			Valid: true,
		},
		High: sql.NullInt32{
			Int32: 128,
			Valid: true,
		},
		Low: sql.NullInt32{
			Int32: 98,
			Valid: true,
		},
		Username: sql.NullString{
			String: "drugo",
			Valid:  true,
		},
	})
	if err != nil {
		panic(err)
	}
	log.Println(insertedMeasurement)

	srv.Run()
}
