package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

var i *interp.Interpreter
var fnMap map[string]bool
var fnMapMutex = &sync.RWMutex{}

func load(fnName string) bool {
	fnMapMutex.RLock()
	if val, exist := fnMap[fnName]; exist {
		fnMapMutex.RUnlock()
		return val
	}

	fnMapMutex.RUnlock()
	fnMapMutex.Lock()
	defer fnMapMutex.Unlock()

	fmt.Println("Reading fn code: " + fnName)
	fnCode, err := os.ReadFile("services/" + fnName + ".go")
	if err != nil {
		fnMap[fnName] = false
		return false
	}

	_, err = i.Eval(string(fnCode))
	if err != nil {
		panic(err)
	}
	fnMap[fnName] = true
	return true
}

func handleHTTP(w http.ResponseWriter, req *http.Request) {
	fnName := strings.TrimPrefix(req.URL.Path, "/")
	fmt.Println("Executing fn: " + fnName)

	if !load(fnName) {
		fmt.Fprintf(w, "not found\n")
		return
	}
	v, _ := i.Eval(fnName + ".HandleHTTP")

	fn := v.Interface().(func(http.ResponseWriter, *http.Request))
	fn(w, req)
}

func handleJSON(fnName string, message string) string {
	return ""
}

func main() {
	i = interp.New(interp.Options{})
	i.Use(stdlib.Symbols)

	fnMap = make(map[string]bool)

	handler := http.HandlerFunc(handleHTTP)
	fmt.Println("Listening...")
	http.ListenAndServe(":8080", handler)
}
