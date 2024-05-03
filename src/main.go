package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nats-io/nats.go"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("Error connecting to NATS server")
	}
	defer nc.Close()

	nc.Subscribe("updates", func(m *nats.Msg) {
		log.Printf("Received a message: %s\n", string(m.Data))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			Message: "Hello World!",
			Status:  "OK",
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Fatal("Error converting struct to json")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})

	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
