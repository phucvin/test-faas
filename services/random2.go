package random2

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Req struct {
	Type *string `json:"type"`
}

type Res struct {
	Error       *string `json:"error,omitempty"`
	RandomValue *string `json:"randomValue"`
}

func HandleHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "not implemented")
}

func HandleJSON(message string) string {
	var req Req
	err := json.Unmarshal([]byte(message), &req)
	if err != nil {
		return `{"error": "bad request format"}`
	}
	res := handle(&req)
	resB, err := json.Marshal(res)
	if err != nil {
		fmt.Printf(err.Error())
		return `{"error": "internal error"}`
	}
	return string(resB)
}

func handle(req *Req) Res {
	r := rand.Intn(5000)
	time.Sleep(time.Duration(r) * time.Millisecond)
	randomValue := fmt.Sprintf("Bob %d", r)
	return Res{RandomValue: &randomValue}
}
