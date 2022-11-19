package hello

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"

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

func callSomeExternalApi() (string, error) {
	res, err := http.Get("https://randomuser.me/api/")
	if err != nil {
		return "", err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(resBody), nil
}

func HandleHTTP(w http.ResponseWriter, req *http.Request) {
	var randomRes RandomRes
	var externalRes string
	var externalErr error
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		randomRes = callRandom()
		wg.Done()
	}()
	go func() {
		externalRes, externalErr := callSomeExternalApi()
		wg.Done()
	}()
	wg.Wait()

	if randomRes.Error != nil || externalErr != nil {
		fmt.Fprintf(w, "error getting a random name to say hello\n\n%v\n\n%v", randomRes.Error, externalErr)
	} else {
		fmt.Fprintf(w, "hello %s\n\nuser:\n%s", *randomRes.RandomValue, externalRes)
	}
}

func HandleJSON(message string) string {
	return `{"error": "not implemented"}`
}
