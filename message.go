package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func handlerMessage(w http.ResponseWriter, r *http.Request) {
	type Msg struct {
		Message string `json:"msg"`
	}

	decoder := json.NewDecoder(r.Body)
	msg := Msg{}

	err := decoder.Decode(&msg)
	if err != nil {
		log.Printf("Error decoding message: %s", err)
		w.WriteHeader(500)
		return
	}

	type returnVals struct {
		CreatedAt time.Time `json:"created_at"`
		ID        int       `json:"id"`
	}

	respBody := returnVals{
		CreatedAt: time.Now(),
		ID:        123,
	}

	dat, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)
}
