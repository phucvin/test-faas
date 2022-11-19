package hello

import (
    "fmt"
    "net/http"
    "encoding/json"
)

func HandleHTTP(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}

func HandleJSON(message string) string {
	return ""
}