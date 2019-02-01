package main

import (
  "html/template"
  "path/filepath"
  "fmt"
  "net/http"
  "io"
)

func hello(
res http.ResponseWriter,
req *http.Request) {
  res.Header().Set(
    "Content-Type",
    "text/html",
  )
  io.WriteString(
    res,
    `<doctype html>
    <html>
      <head>
      <title>Hello World</title>
      </head>
      <body>
      <h1>Hello World!</h1>
      </body>
    </html>`,
)}


func serveTemplate(w http.ResponseWriter, r *http.Request) {
  lp := filepath.Join("templates", "layout.html")
  fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

  tmpl, _ := template.ParseFiles(lp, fp)
  tmpl.ExecuteTemplate(w, "layout", nil)
}

func route() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

  http.HandleFunc("/hello", hello)
  http.HandleFunc("/default/", serveTemplate)

  fs := http.FileServer(http.Dir("./static/"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))
}

func main() {
  fmt.Println("Server is ready at port 9000")

  route()

  http.ListenAndServe(":9000", nil)
}
