package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestParams(t *testing.T){

	router := httprouter.New()

	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, error interface{}) {
		fmt.Fprint(writer, "Panic: ", error)
	}

	router.GET("/panic", func (writer http.ResponseWriter, request *http.Request, pararms httprouter.Params)  {
		panic("error")
	})

	request := httptest.NewRequest("GET", "http://localhost:8000/panic", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Panic: error", string(body))
}

func TestNamedParams(t *testing.T){

	router := httprouter.New()
	router.GET("/product/:id/item/:itemId", func (writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
		text := "Product " +  params.ByName("id") +" item " + params.ByName("itemId")
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:8000/product/1/item/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1 item 1", string(body))
}

func TestCatchAllParams(t *testing.T){

	router := httprouter.New()
	router.GET("/images/*image", func (writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
		text := "Image : " +  params.ByName("image")
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:8000/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Image : /small/profile.png", string(body))
}