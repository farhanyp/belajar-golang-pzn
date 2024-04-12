package server

import (
	"fmt"
	"io"
	_ "net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request){
	contentType := request.Header.Get("content-type")
	
	fmt.Fprint(writer, contentType)
}

func TestHeader(t *testing.T) {


	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/", nil)
	request.Header.Add("Content-type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}

func ResoponseHeader(writer http.ResponseWriter, request *http.Request){
	writer.Header().Add("X-Powered-By", "Tes 1 dan 2")
	
}

func TestResoponseHeader(t *testing.T) {


	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/", nil)
	recorder := httptest.NewRecorder()

	ResoponseHeader(recorder, request)

	PoweredBy := recorder.Header().Get("X-Powered-By")

	fmt.Println(PoweredBy)

}