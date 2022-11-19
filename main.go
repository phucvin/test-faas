package main

import (
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

func Main() {
    http.HandleFunc("/hello", hello)

    fmt.Println("Listening...")
    http.ListenAndServe(":8080", nil)
}
`

func main() {
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)

	_, err := i.Eval(src)
	if err != nil {
		panic(err)
	}

	v, err := i.Eval("code.Main")
	if err != nil {
		panic(err)
	}

	fn := v.Interface().(func())

	fn()
}
