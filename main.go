package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseMessage struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		resBody := ResponseMessage{
			Message: "Success send first body",
			Status:  http.StatusOK,
		}

		err := json.NewEncoder(w).Encode(resBody)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to load server: ", err)
	}
}
