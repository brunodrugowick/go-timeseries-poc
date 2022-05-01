package main

import (
	"encoding/json"
	"github.com/brunodrugowick/go-timeseries-poc/pkg/server"
	"html/template"
	"net/http"
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

	srv.Run()
}
