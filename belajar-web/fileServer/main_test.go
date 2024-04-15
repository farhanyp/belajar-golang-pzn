package server

import (
	"embed"
	"io/fs"
	_ "net"
	"net/http"
	"testing"
)

// func TestFileServer(t *testing.T) {

// 	directory := http.Dir("./resource")
// 	fileServer := http.FileServer(directory)

// 	mux := http.NewServeMux()
// 	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

// 	server := http.Server{
// 		Addr: "localhost:8000",
// 		Handler: mux,
// 	}

// 	err := server.ListenAndServe()
// 	if err != nil {
// 		panic(err)
// 	}

// }

//go:embed resource
	var resource embed.FS

func TestFileServerGoEmbed(t *testing.T) {

	directory, _ := fs.Sub(resource, "resource")
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr: "localhost:8000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}