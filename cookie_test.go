package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	// cookie := http.Cookie{
	// 	Name:  "my-cookie",
	// 	Value: "my-value",
	// 	Path:  "/",
	// }
	// http.SetCookie(w, &cookie)

	cookie := new(http.Cookie)
	cookie.Name = "my-cookie"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Cookie set")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("my-cookie")
	if err != nil {
		fmt.Fprint(w, "Cookie not found")
	} else {
		fmt.Fprint(w, cookie.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/set-cookie?name=BudiUtomo", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	response := recorder.Result()

	cookies := response.Cookies()

	for _, cookie := range cookies {
		fmt.Println(cookie)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/get-cookie", nil)
	recorder := httptest.NewRecorder()

	cookie := new(http.Cookie)
	cookie.Name = "my-cookie"
	cookie.Value = "BudiUtomo"
	request.AddCookie(cookie)

	recorder = httptest.NewRecorder()

	GetCookie(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
