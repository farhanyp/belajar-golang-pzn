package server

import (
	"fmt"
	"io"
	_ "net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request){
	cookie := new(http.Cookie)
	cookie.Name = "X-yp-name"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Cookie has been created")
}

func GetCookie(writer http.ResponseWriter, request *http.Request){
	cookie, err := request.Cookie("X-yp-name")
	if err != nil {
		fmt.Fprint(writer, "No Cookie")
	}else{
		fmt.Fprintf(writer, "Hello %s", cookie.Value)
	}
}

// func TestTryCookie(t *testing.T) {

// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/set-cookie", SetCookie)
// 	mux.HandleFunc("/get-cookie", GetCookie)

// 	server := http.Server{
// 		Addr: "localhost:8000",
// 		Handler: mux,
// 	}

// 	err := server.ListenAndServe()
// 	if err != nil {
// 		panic(err)
// 	}

// }

func TestGetCookie(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-yp-name"
	cookie.Value = "yp"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}