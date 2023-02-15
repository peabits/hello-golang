package example

import (
	"fmt"
	"net/http"
	"path"
	"runtime"
)

func HelloServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "<h1>Welcome to my website</h1><br /><a href=\"/static/\">跳转</a>")
	})

	// f, _ := exec.LookPath(os.Args[0])
	_, c, _, _ := runtime.Caller(0)
	// e, _ := os.Executable()
	// fmt.Println(f, c, e)

	fs := http.FileServer(http.Dir(path.Dir(c) + "/static/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
