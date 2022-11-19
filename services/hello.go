package hello

import (
	"encoding/json"
	"fmt"
	"net/http"
	"math/rand"

	"call"
)

type RandomRes struct {
	Error       *string `json:"error,omitempty"`
	RandomValue *string `json:"randomValue"`
}

func callRandom() RandomRes {
	token := rand.Intn(1000000)
	fmt.Printf("calling random, token-%d\n", token)
	resS := call.JSON("random2", "{}")
	fmt.Printf("done calling random, token-%d\n", token)
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
	if randomRes.Error != nil {
		fmt.Fprintf(w, "error getting a random name to say hello")
	} else {
		fmt.Fprintf(w, "hello "+*randomRes.RandomValue)
	}
}

func HandleJSON(message string) string {
	return `{"error": "not implemented"}`
}
