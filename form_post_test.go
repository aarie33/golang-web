package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()

	if err != nil {
		panic(err)
	}

	request.PostFormValue("first_name")

	first_name := request.PostForm.Get("first_name")
	last_name := request.PostForm.Get("last_name")

	fmt.Fprint(writer, "Hello ", first_name, " ", last_name)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Adi&last_name=Santoso")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
