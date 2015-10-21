package recorder2

import (
  "fmt"
  "net/http"
  "time"
)

func init() {
  http.HandleFunc("/", root)
  http.HandleFunc("/records", records)
}

func root(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "please login") // TODO: add login function
}

func records(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
  case "POST": fmt.Fprintf(w, "records posted %s", http.PostForm) // TODO: parse JSON
  default: http.Error(w, "error occurred", http.StatusInternalServerError) // TODO: use template file
  }
}

// TODO: separate into other file
type Record struct {
  GroupId int
  Name string
  ReceivedAt time.Time
}
