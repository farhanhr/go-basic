package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")

}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	/* recorder adalah salah satu implementasi dari response writer
	Response recorder merupakan struct bantuan untuk merekam Http 
	response dari hasil testing */
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)

	fmt.Println(bodyString)
}