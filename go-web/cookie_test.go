package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-FHR-Cookie"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Cookie Created")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-FHR-Cookie")
	if err != nil {
		fmt.Fprint(w, "No Cookie")
	} else {
		name := cookie.Value
		fmt.Fprintf(w, "Hello %s", name)
	}
	
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)
	
	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}


	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=Farhan", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s: %s \n", cookie.Name, cookie.Value)
	}
	
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-FHR-Cookie"
	cookie.Value = "Farhan HR"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()
	
	GetCookie(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}