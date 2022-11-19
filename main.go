package main

import (
	"sync"

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

func main() {
	i := interp.New(interp.Options{GoPath: "/workspace/go"})
	i.Use(stdlib.Symbols)

	_, err := i.Eval(src)
	if err != nil {
		panic(err)
	}

	v, err := i.Eval("code.Main")
	if err != nil {
		panic(err)
	}

	fn := v.Interface().(func(string))

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fn(":8080")
	}()
	go func() {
		fn(":8090")
	}()
	wg.Wait()
}
