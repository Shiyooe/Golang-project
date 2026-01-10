package response

import (
    "encoding/json"
    "net/http"
)

func JSON(w http.ResponseWriter, data interface{}, status int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, message string, status int) {
    JSON(w, map[string]string{"error": message}, status)
}