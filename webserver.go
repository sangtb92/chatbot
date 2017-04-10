package main

import (
	"net/http"
	"io"
)

func hello(r http.ResponseWriter, req *http.Request) {
	r.Header().Set(
		"Content-type",
		"text/html",
	)
	content := `<!DOCTYPE html>
                <html>
                    <head>
                        <title>Sample Go Web Server</title>
                    </head>
                    <body>
                        <h1>It Worked!</h1>
                    </body>
                </html>`
	io.WriteString(r, content)

}
func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
