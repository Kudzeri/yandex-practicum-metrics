package main

import (
	"github.com/Kudzeri/yandex-practicum-metrics/internal/storage"
	"github.com/Kudzeri/yandex-practicum-metrics/internal/server"
)

func main() {
	storage  := storage.InitMemStorage()
	srv := server.NewServer(storage)

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
