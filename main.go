package main

import (
	"fmt"
	"html"
	"net/http"
)

const template = `<!DOCTYPE html>
<html>
	<head>
	</head>
	<body>
		<h2>Thanks for visiting my path %q</h2>
	</body>
</html>
`

func main() {
	port := ":3000"
	http.HandleFunc("/", root)
	fmt.Printf("Listening on port %s\n", port)
	panic(http.ListenAndServe(port, nil))
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html; charset=utf-8")
	fmt.Fprintf(w, template, html.EscapeString(r.URL.Path))
	return
}
