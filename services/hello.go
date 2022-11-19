package hello

import (
	"encoding/json"
	"fmt"
	"net/http"

	"call"
)

type RandomRes struct {
	Error       *string `json:"error,omitempty"`
	RandomValue *string `json:"randomValue"`
}

func callRandom() RandomRes {
	resS := call.JSON("random2", "{}")
	var res RandomRes
	err := json.Unmarshal([]byte(resS), &res)
	if err != nil {
        errS := "error parsing result"
		return RandomRes{Error: &errS}
	}
	return res
}

func HandleHTTP(w http.ResponseWriter, req *http.Request) {
	randomRes := callRandom()
	fmt.Fprintf(w, "hello "+*randomRes.RandomValue)
}

func HandleJSON(message string) string {
	return `{"error": "not implemented"}`
}
