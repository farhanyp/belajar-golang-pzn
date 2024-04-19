package server

import (
	"fmt"
	"io"
	_ "net"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

type MyPage struct{
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name is " + myPage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request){
	t := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello "Yp" }}`))

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Farhan",
	})
}

func TestTemplateFunction(t *testing.T) {

	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}