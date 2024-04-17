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

func SimpleHtml (writer http.ResponseWriter, request *http.Request){
	templateText := `<html><body>{{.}}</body></html>`
	t, err := template.New("SIMPLE").Parse(templateText)
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "SIMPLE", "Hello World")
}

func SimpleTemplateHtml (writer http.ResponseWriter, request *http.Request){
	t, err := template.ParseFiles("./resource/index.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "index.gohtml", "Hello World template")
}

func TemplateDir (writer http.ResponseWriter, request *http.Request){
	t, err := template.ParseGlob("./resource/*.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "index.gohtml", "Hello World template")
}

//go:embed resource/*.gohtml
var templates embed.FS

func TemplateEmbed (writer http.ResponseWriter, request *http.Request){
	t, err := template.ParseFS(templates, "resource/*.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "index.gohtml", "Hello World template")
}

func TestSimpleHtml(t *testing.T) {

	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	SimpleHtml(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestSimpleTemplateHtml(t *testing.T) {

	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateDir(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestTemplateEmbed(t *testing.T) {

	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}