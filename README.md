# test-faas

go mod tidy

go run main.go

localhost:8080/hello

localhost:8080/_invalidateAll

TODO:
- Call https://randomuser.me/api/ or https://api.nationalize.io?name=michael from inside a service
- Test parallel calls to services
- In memory K/V store
