package _template

import (
    "fmt"
    "net/http"
    "encoding/json"
)

func HandleHTTP(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "not implemented")
}

func HandleJSON(message string) string {
	return "{\"err\": \"not implemented\"}"
}