package server

import (
	"fmt"
	_ "net"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request)  {
		fmt.Fprint(writer, "hello world")
	})

	mux.HandleFunc("/images", func(writer http.ResponseWriter, request *http.Request)  {
		fmt.Fprint(writer, "images")
	})

	mux.HandleFunc("/images/tumbnail", func(writer http.ResponseWriter, request *http.Request)  {
		fmt.Fprint(writer, "tumbnail")
	})

	server := http.Server{
		Addr: "localhost:8000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}