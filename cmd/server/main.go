package main

import "github.com/Kudzeri/yandex-practicum-metrics/internal/server"

func main() {
	srv := server.NewServer()

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
