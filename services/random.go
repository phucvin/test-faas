package random

import (
    "fmt"
    "net/http"
    "encoding/json"
)

func HandleHTTP(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "not implemented\n")
}

func HandleJSON(message string) string {
	return "{\"name\": \"John\"}"
}