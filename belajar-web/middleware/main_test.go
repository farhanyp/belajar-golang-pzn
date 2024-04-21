package server

import (
	"fmt"
	_ "net"
	"net/http"
	"testing"
)

type LogMiddleware struct{
	Handler http.Handler
}

type ErrorHandler struct{
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request){
	defer func ()  {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()

	errorHandler.Handler.ServeHTTP(writer, request)

}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request){
	fmt.Println("Before Execute Handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Execute Handler")
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprintf(writer, "Hello Middleware")
	})

	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprintf(writer, "panic Middleware")
		panic("Panic cuy")
	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	ErrorMiddleware := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: ErrorMiddleware,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}