package example

import (
	"html/template"
	"net/http"
	"path"
	"runtime"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func Templates() {
	_, c, _, _ := runtime.Caller(0)

	tmpl := template.Must(template.ParseFiles(path.Dir(c) + "/layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO List",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)
}
