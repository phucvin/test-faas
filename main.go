package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

const (
	FN_LOAD_OK    = 0
	FN_NOT_FOUND  = 1
	FN_LOAD_ERROR = 2
)

var i *interp.Interpreter
var fnMap map[string]int
var fnMapMutex *sync.RWMutex

func load(fnName string) int {
	fnMapMutex.RLock()
	if val, exist := fnMap[fnName]; exist {
		fnMapMutex.RUnlock()
		return val
	}

	fnMapMutex.RUnlock()
	fnMapMutex.Lock()
	defer fnMapMutex.Unlock()

	fmt.Println("Reading fn: " + fnName)
	fnCode, err := os.ReadFile("services/" + fnName + ".go")
	if err != nil {
		fnMap[fnName] = FN_NOT_FOUND
		return 0
	}

	fnMap[fnName] = FN_LOAD_OK
	_, err = i.Eval(string(fnCode))
	if err != nil {
		fnMap[fnName] = FN_LOAD_ERROR
		fmt.Println(err)
	}
	return fnMap[fnName]
}

func handleHTTP(w http.ResponseWriter, req *http.Request) {
	fnName := strings.TrimPrefix(req.URL.Path, "/")
	fmt.Println("Serving fn: " + fnName)

	fnLoadStatus := load(fnName)
	if fnLoadStatus == FN_NOT_FOUND {
		fmt.Fprintf(w, "not found")
		return
	}
	if fnLoadStatus == FN_LOAD_ERROR {
		fmt.Fprintf(w, "error loading")
		return
	}
	v, _ := i.Eval(fnName + ".HandleHTTP")

	fn := v.Interface().(func(http.ResponseWriter, *http.Request))
	fn(w, req)
}

func callJSON(fnName string, message string) string {
	fmt.Println("Calling fn: " + fnName)
	fnLoadStatus := load(fnName)
	if fnLoadStatus == FN_NOT_FOUND {
		return `{"error": "not found"}`
	}
	if fnLoadStatus == FN_LOAD_ERROR {
		return `{"error": "error loading"}`
	}
	v, _ := i.Eval(fnName + ".HandleJSON")
	fn := v.Interface().(func(string) string)
	return fn(message)
}

func main() {
	i = interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	additionalSymbols := make(map[string]map[string]reflect.Value)
	additionalSymbols["call/call"] = map[string]reflect.Value{"JSON": reflect.ValueOf(callJSON)}
	i.Use(additionalSymbols)

	fnMap = make(map[string]int)
	fnMapMutex = &sync.RWMutex{}
	rand.Seed(time.Now().UnixNano())

	handler := http.HandlerFunc(handleHTTP)
	fmt.Println("Listening...")
	http.ListenAndServe(":8080", handler)
}
