package main

import (
	"$GOPATH/web/myapp/app.go"
	"net/http"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
