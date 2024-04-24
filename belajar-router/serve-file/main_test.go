package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T){

	router := httprouter.New()
	dir, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(dir))

	request := httptest.NewRequest("GET", "http://localhost:8000/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(dir)

	assert.Equal(t, "Hai httprouter", string(body))
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