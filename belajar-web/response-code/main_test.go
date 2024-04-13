package server

import (
	"fmt"
	"io"
	_ "net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request){

	name := request.URL.Query().Get("name")

	if name == ""{
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "name is empty")
	}else{
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "Hi %s ", name)
	}
}

func TestPostForm(t *testing.T) {

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/?name=yp", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))

}