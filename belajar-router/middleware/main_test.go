package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	// "github.com/stretchr/testify/assert"
)

type LogMiddleware struct{
	http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request){
	fmt.Fprint(writer, "Recieve Request")
	middleware.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T){

	router := httprouter.New()

	router.GET("/", func (writer http.ResponseWriter, request *http.Request, pararms httprouter.Params)  {
		fmt.Fprint(writer, "Ini root")
	})

	middleware := LogMiddleware{Handler: router}

	request := httptest.NewRequest("GET", "http://localhost:8000/", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	// assert.Equal(t, "Method tidak ada", string(body))
}