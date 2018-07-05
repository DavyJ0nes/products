package handlers

import "fmt"

func rootHandler(w http.responseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello")
}
