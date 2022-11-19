package main

import (
	"fmt"
	"net/http"
	"strings"
	"os"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

const src = `
package code

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}

func Main(addr string) {
	handler := http.HandlerFunc(hello)

    fmt.Println("Listening...")
    http.ListenAndServe(addr, handler)
}
`

var i *interp.Interpreter
var fnMap map[string]bool

func handleHTTP(w http.ResponseWriter, req *http.Request) {
	fnName := strings.TrimPrefix(req.URL.Path, "/")
	fmt.Println("Executing fn: " + fnName)

	if _, exist := fnMap[fnName]; !exist {
		fnCode, err := os.ReadFile("services/" + fnName + ".go")
		if err != nil {
    		fmt.Fprintf(w, "not found\n")
			return
		}
		_, err = i.Eval(string(fnCode))
		if err != nil {
			panic(err)
		}
	}
	fnMap[fnName] = true
	v, _ := i.Eval(fnName + ".HandleHTTP")

	fn := v.Interface().(func(http.ResponseWriter, *http.Request))
	fn(w, req)
}

func main() {
	i = interp.New(interp.Options{})
	i.Use(stdlib.Symbols)

	fnMap = make(map[string]bool)

	handler := http.HandlerFunc(handleHTTP)
	fmt.Println("Listening...")
    http.ListenAndServe(":8080", handler)
}
