package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Kudzeri/yandex-practicum-metrics/internal/storage"
)

func PingHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running"))
}

func UpdateMetric(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 5 || parts[1] != "update" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		metricType, metricName, metricValue := parts[2], parts[3], parts[4]

		if metricName == "" {
			http.Error(w, "Metric name required", http.StatusNotFound)
			return
		}

		switch metricType {
		case "gauge":
			val, err := strconv.ParseFloat(metricValue, 64)
			if err != nil {
				http.Error(w, "Invalid gauge value", http.StatusBadRequest)
				return
			}
			storage.UpdateGauge(metricName, val)

		case "counter":
			val, err := strconv.ParseInt(metricValue, 10, 64)
			if err != nil {
				http.Error(w, "Invalid counter value", http.StatusBadRequest)
				return
			}
			storage.UpdateCounter(metricName, val)

		default:
			http.Error(w, "Unknown metric type", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func GetMetric(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 4 || parts[1] != "value" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		metricType, metricName := parts[2], parts[3]

		if metricName == "" {
			http.Error(w, "Metric name required", http.StatusNotFound)
			return
		}

		switch metricType {
		case "gauge":
			if val, ok := storage.GetGauge(metricName); ok {
				w.Write([]byte(fmt.Sprintf("%v", val)))
			} else {
				http.Error(w, "Metric not found", http.StatusNotFound)
			}

		case "counter":
			if val, ok := storage.GetCounter(metricName); ok {
				w.Write([]byte(fmt.Sprintf("%v", val)))
			} else {
				http.Error(w, "Metric not found", http.StatusNotFound)
			}

		default:
			http.Error(w, "Unknown metric type", http.StatusBadRequest)
		}
	}
}
