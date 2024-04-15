package server

import (
	_ "embed"
	"fmt"
	_ "net"
	"net/http"
	"testing"
)

// func ServeFile (writer http.ResponseWriter, request *http.Request){

// 	if request.URL.Query().Get("name") != "" {
// 		http.ServeFile(writer, request, "./resource/ok.html")
// 	}else{
// 		http.ServeFile(writer, request, "./resource/not-found.html")
// 	}

// }

// func TestServerFile(t *testing.T) {

// 	server := http.Server{
// 		Addr: "localhost:8000",
// 		Handler: http.HandlerFunc(ServeFile),
// 	}

// 	err := server.ListenAndServe()
// 	if err != nil {
// 		panic(err)
// 	}

// }

//go:embed resource/ok.html
	var okfile string

//go:embed resource/not-found.html
var notFound string

func ServerFileWithEmbed (writer http.ResponseWriter, request *http.Request){

	if request.URL.Query().Get("name") != "" {
		fmt.Fprint(writer, okfile)
	}else{
		fmt.Fprint(writer, notFound)
	}

}

func TestServerFileWithEmbed(t *testing.T) {

	server := http.Server{
		Addr: "localhost:8000",
		Handler: http.HandlerFunc(ServerFileWithEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}