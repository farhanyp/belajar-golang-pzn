package server

import (
	"fmt"
	_ "net"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {

	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request)  {
		fmt.Fprint(writer, "hello world")
	}

	server := http.Server{
		Addr: "localhost:8000",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}