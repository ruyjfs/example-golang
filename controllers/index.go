package controllers

import (
	"fmt"
	"html"
	"net/http"
)

type Index struct {
	// core.Controller
}

func (i *Index) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

}
