package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + " my name is " + myPage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "John"}}`))
	tmpl.ExecuteTemplate(writer, "FUNCTION", MyPage{Name: "Doe"})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	tmpl.ExecuteTemplate(writer, "FUNCTION", MyPage{Name: "Doe"})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateFunctionCreateGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func(s string) string {
			return strings.ToUpper(s)
		},
	})
	tmpl := template.Must(t.Parse(`{{ upper .Name}}`))
	tmpl.ExecuteTemplate(writer, "FUNCTION", MyPage{Name: "Doe"})
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateFunctionCreateGlobalPipeline(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		"upper": func(s string) string {
			return strings.ToUpper(s)
		},
	})
	tmpl := template.Must(t.Parse(`{{ sayHello .Name |  upper }}`))
	tmpl.ExecuteTemplate(writer, "FUNCTION", MyPage{Name: "Doe"})
}

func TestTemplateFunctionCreateGlobalPipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobalPipeline(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
