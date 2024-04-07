package server

import (
	"fmt"
	"io"
	_ "net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func singleQuery(writer http.ResponseWriter, request *http.Request){
	name := request.URL.Query().Get("name")
	
	fmt.Fprint(writer, name)
}

func TestSingleQuery(t *testing.T) {


	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/?name=yp", nil)
	recorder := httptest.NewRecorder()

	singleQuery(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}

func multipleQuery(writer http.ResponseWriter, request *http.Request){
	first_name := request.URL.Query().Get("first_name")
	last_name := request.URL.Query().Get("last_name")
	
	fmt.Fprint(writer, first_name)
	fmt.Fprint(writer, last_name)
}

func TestMultipleQuery(t *testing.T) {


	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/?first_name=yp&last_name=yp2", nil)
	recorder := httptest.NewRecorder()

	multipleQuery(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}

func multipleValue(writer http.ResponseWriter, request *http.Request){
	var query url.Values = request.URL.Query()
	var names []string = query["name"]
	
	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestMultipleValue(t *testing.T) {


	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/?name=yp&name=yp2", nil)
	recorder := httptest.NewRecorder()

	multipleValue(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}