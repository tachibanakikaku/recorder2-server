package recorder2

import (
  "fmt"
  "net/http"
)

func init() {
  http.HandleFunc("/", record)
}

func record(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "got request")
}
