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

func TemplateLayout(writer http.ResponseWriter, request *http.Request){
	t := template.Must(template.ParseFiles(
		"./resource/header.gohtml",
		"./resource/footer.gohtml",
		"./resource/index.gohtml",
	))

	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Name": "Yp",
		"Title": "Tes Template",
	})
}

func TestTemplateActionIf(t *testing.T) {

	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}