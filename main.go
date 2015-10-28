package recorder2

import (
  "fmt"
  "io"
  "time"

  "html/template"
  "net/http"

  "appengine"
  "appengine/datastore"
  "appengine/user"
)

/*
Utility Methods
*/

func layout_variables(r *http.Request) map[string]interface{} {
  u := user.Current(appengine.NewContext(r))
  return map[string]interface{} {
    "user": u.String(),
  }
}

func render(t string, w io.Writer, data map[string]interface{}) {
  tmpl := template.Must(template.ParseFiles("views/layout.html", t))
  tmpl.Execute(w, data)
}

func save_user(w http.ResponseWriter, c appengine.Context, u *user.User) {
  el := Record {
    GroupId: "11",
    Name: u.String(),
    ID: u.ID,
    Email: u.Email,
    AuthDomain: u.AuthDomain,
    ReceivedAt: time.Now(),
  }
  _, err := datastore.Put(
    c,
    datastore.NewIncompleteKey(c, "records", nil),
    &el,
  )
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

/*
Routing
*/

func init() {
  http.HandleFunc("/", index)
  http.HandleFunc("/login", login)
  http.HandleFunc("/records", records)
}

/*
Request Handlers
*/

func index(w http.ResponseWriter, r *http.Request) {
  u := user.Current(appengine.NewContext(r))
  if u == nil {
    http.Redirect(w, r, "/login", http.StatusMovedPermanently)
    return
  }
  data := layout_variables(r)
  data["title"] = "Top Page"
  render("views/index.html", w, data)
}

func login(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  u := user.Current(c)
  url, err := user.LoginURL(c, r.URL.String())
  if u == nil {
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    w.Header().Set("Location", url)
    w.WriteHeader(http.StatusFound)
  } else {
    save_user(w, c, u)
    http.Redirect(w, r, "/", http.StatusMovedPermanently)
  }
}

func records(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
  case "POST": fmt.Fprintf(w, "records posted %s", http.PostForm) // TODO: parse JSON
  default: http.Error(w, "error occurred", http.StatusInternalServerError)
  }
}
