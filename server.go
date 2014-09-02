package main

import (
	"./models"
	"github.com/codegangsta/negroni"
	"html/template"
	"net/http"
	"strconv"
)

type Params struct {
	Title   string
	Members []models.Member
}

var templates = template.Must(template.ParseFiles("index.html", "add.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Params) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/add", addHandler)
	mux.HandleFunc("/delete", deleteHandler)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3001")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	members, err := models.LoadMembers(0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p := &Params{Title: "index", Members: members}
	renderTemplate(w, "index", p)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	m := models.Member{Name: name}
	err := m.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	err := models.Delete(int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
