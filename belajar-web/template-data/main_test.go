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

func TemplateDirWithMap(writer http.ResponseWriter, request *http.Request){
	t, err := template.ParseFS(templates, "resource/*.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "index.gohtml", map[string]interface{}{
		"Title": "Template Data Struct",
		"Name": "Yp",
	})
}

type Addres struct{
	Street string
}

type Page struct{
	Title string
	Name string
	Addres Addres
}

func TemplateDirWithStruct(writer http.ResponseWriter, request *http.Request){
	t, err := template.ParseFS(templates, "resource/*.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "index.gohtml", Page{
		Title: "Template Data Struct",
		Name: "Yp",
		Addres: Addres{
			Street:  "Jauh",
		},
	})
}

func TestTemplateDirWithMap(t *testing.T) {

	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateDirWithStruct(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}