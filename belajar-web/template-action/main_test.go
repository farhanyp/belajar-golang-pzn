package server

import (
	"embed"
	"fmt"
	"io"
	_ "net"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

//go:embed resource/*.gohtml
var templates embed.FS

func TemplateIf(writer http.ResponseWriter, request *http.Request){
	t, err := template.ParseFS(templates, "resource/*.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "index.gohtml", map[string]interface{}{
		"FinalValue": 60,
	})
}

func TestTemplateActionIf(t *testing.T) {

	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateIf(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateRange(writer http.ResponseWriter, request *http.Request){
	t, err := template.ParseFS(templates, "resource/*.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Hobbies": []string{
			"Gaming", "Reading", "Coding",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {

	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateRange(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateWith(writer http.ResponseWriter, request *http.Request){
	t, err := template.ParseFS(templates, "resource/*.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "with.gohtml", map[string]interface{}{
		"Name":"Eko",
		"Address": map[string]interface{}{
			"Street": "Jalan Belum Jadi",
			"City": "Mars",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {

	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateWith(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}