package hello

import (
    "fmt"
    "net/http"
    "encoding/json"

    "call"
)

func HandleHTTP(w http.ResponseWriter, req *http.Request) {
    randomRes := call.JSON("random", "{}")
    fmt.Fprintf(w, "hello " + randomRes)
}

func HandleJSON(message string) string {
	return "{\"err\": \"not implemented\"}"
}