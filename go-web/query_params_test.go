package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request)  {
	name := request.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(writer, "Hello User")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/8080/hello?name=Farhan", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleQueryParams(writer http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)

}

func TestMultipleQueryParams(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/8080/hello?first_name=Farhan&last_name=HusyenR", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParams(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleParameterValue(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]
	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestMultipleParameterValue(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/8080/hello?name=Farhan&name=Husyen&name=Ramadhan", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValue(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}