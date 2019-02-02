package main

import (
  "html/template"
  "path/filepath"
  "fmt"
  "net/http"
  "io"
  "log"
  "os"
  "github.com/gorilla/mux"
  "encoding/json"
)

type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}
type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

var people []Person

func hello( res http.ResponseWriter, req *http.Request ) {
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
  fmt.Println(r.URL.Path, "   serveTemplate function")
  lp := filepath.Join("templates", "layout.html")
  fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

  // Return a 404 if the template doesn't exist
  info, err := os.Stat(fp)
  if err != nil {
    fmt.Println(err.Error())
    if os.IsNotExist(err) {
      http.NotFound(w, r)
      return
    }
  }

  if info.IsDir() {
    http.NotFound(w, r)
    return
  }

  tmpl, err := template.ParseFiles(lp, fp)
  if err != nil {
    log.Println(err.Error())
    http.Error(w, http.StatusText(500), 500)
    return
  }

  if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
    log.Println(err.Error())
    http.Error(w, http.StatusText(500), 500)
  }
}

func login(w http.ResponseWriter, r *http.Request) {
  fmt.Println("method:", r.Method) //get request method
  if r.Method == "GET" {
    t, _ := template.ParseFiles("login.gtpl")
    t.Execute(w, nil)
  } else {
    r.ParseForm()
    // logic part of log in
    fmt.Println("username:", r.Form["username"])
    fmt.Println("password:", r.Form["password"])
  }
}

func backend(r *mux.Router) {
  r.HandleFunc("/backend/people", GetPeople)
}

func GetPeople(w http.ResponseWriter, r *http.Request)  {
  json.NewEncoder(w).Encode(people)
}

func route() {
  r := mux.NewRouter()
  r.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

  r.HandleFunc("/hello", hello)
  backend(r)

  r.HandleFunc("/", serveTemplate)
  //
  // fs := http.FileServer(http.Dir("./static/"))
  // http.Handle("/static/", http.StripPrefix("/static/", fs))

  http.Handle("/", r)
  http.ListenAndServe(":9000", r)
}

func main() {
  people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
  people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
  people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

  fmt.Println("Server is ready at port 9000")
  route()

}
