package example

import (
	"fmt"
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	fmt.Printf("%T\n", f)
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func foo(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "bar")
}

func Middleware() {
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))

	http.ListenAndServe(":8080", nil)
}
