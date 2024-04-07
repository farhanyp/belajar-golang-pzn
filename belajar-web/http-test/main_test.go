package server

import (
	"fmt"
	"io"
	_ "net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintln(writer, "Hello World")
	fmt.Fprintln(writer, request.URL)
}

func TestServer(t *testing.T) {


	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/test", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)

}