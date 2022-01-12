package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello HttpRouter")
	})
	server := http.Server{
		Handler: router,
		Addr: "localhost:1700",
	}

	err := server.ListenAndServe()
	PanicError(err)
}

func PanicError(err error)  {
	if err == nil {
		panic(err)
	}
}
