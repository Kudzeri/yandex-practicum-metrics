package server

import (
	"net/http"

	"github.com/Kudzeri/yandex-practicum-metrics/internal/database"
	"github.com/Kudzeri/yandex-practicum-metrics/internal/handlers"
	"github.com/Kudzeri/yandex-practicum-metrics/internal/storage"
)

func NewServer(storage storage.Storage) *http.Server {
	database.InitDB()

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", handlers.PingHandle)
	mux.HandleFunc("/update/", handlers.UpdateMetric(storage))
	mux.HandleFunc("/value/", handlers.GetMetric(storage))

	port := ":8080"
	return &http.Server{
		Addr:    port,
		Handler: mux,
	}
}
