package main

import (
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

type user struct {
	UserName string
	First    string
	Last     string
}

var dbSessions = map[string]string{}
var dbUsers = map[string]user{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// check if session exists / cookie
	c, err := r.Cookie("session")
	if err != nil {
		// if not create one
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Value: sID.String(),
			Name:  "session",
		}
		http.SetCookie(w, c)
	}

	// find the user with the session
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		// get user data
		u = dbUsers[un]
	}

	// create user from POST
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}

	// exec template
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	u := dbUsers[un]

	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
