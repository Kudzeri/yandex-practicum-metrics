package server

import (
	"net/http"

	"github.com/Kudzeri/yandex-practicum-metrics/internal/database"
	"github.com/Kudzeri/yandex-practicum-metrics/internal/handlers"
)

func NewServer() *http.Server {
	database.InitDB()

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", handlers.PingHandle)

	port := ":8080"
	return &http.Server{
		Addr:    port,
		Handler: mux,
	}
}
