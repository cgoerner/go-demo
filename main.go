package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

var counter int
var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

type Response struct {
	Url     string `json:"url"`
	Counter int    `json:"counter"`
}

func main() {
	http.HandleFunc("/counter", handleCounter)
	logger.Info("Listening...")
	http.ListenAndServe(":8080", logRequest(http.DefaultServeMux))
}

func handleCounter(w http.ResponseWriter, r *http.Request) {
	response := Response{Url: r.URL.String(), Counter: counter}
	json, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(json)
	counter++
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info(r.URL.String(), "method", r.Method, "status", 200, "counter", counter)
		handler.ServeHTTP(w, r)
	})
}
