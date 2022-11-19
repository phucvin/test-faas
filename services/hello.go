package hello

import (
	"encoding/json"
	"fmt"
	"net/http"

	"call"
)

func HandleHTTP(w http.ResponseWriter, req *http.Request) {
	randomRes := call.JSON("random", "{}")
	fmt.Fprintf(w, "hello "+randomRes)
}

func HandleJSON(message string) string {
	return `{"error": "not implemented"}`
}
