package main

import (
	"encoding/json"
	"io"
	"log"
	"math/big"
	"net/http"

	"github.com/gorilla/mux"
)

func numHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number := vars["number"]
	n := new(big.Int)
	n, _ = n.SetString(number, 10)
	w.Header().Set("Content-Type", "application/json")
	result := map[string]bool{}
	result["isPrime"] = n.ProbablyPrime(0)
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println("Can not encode", result, "to json")
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}
