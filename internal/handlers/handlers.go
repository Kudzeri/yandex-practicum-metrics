package handlers

import "net/http"

func PingHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running"))
}
