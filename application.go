package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Page defines a wiki entry for a given page. It includes a Title for the Page
// and also a Body for the page contents.
type Page struct {
	Title string
	Body  []byte
}

/* Globals */

// templates defines all the template files we will use and is used in the 'renderTemplate function
//var templates = template.Must(template.ParseFiles("templates/base.html", "templates/edit.html", "templates/view.html", "templates/index.html"))
var templates = make(map[string]*template.Template)
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

/* End Global */

func init() {
	templates["index.html"] = template.Must(template.ParseFiles("templates/index.html", "templates/base.html"))
	templates["view.html"] = template.Must(template.ParseFiles("templates/view.html", "templates/base.html"))
	templates["edit.html"] = template.Must(template.ParseFiles("templates/edit.html", "templates/base.html"))
}

// Pages defines a slice of type Page to hold a list of all the pages in data/ dir.
var Pages = []Page{}

/* CURRENTLY UNUSED.
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression
}*/

func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	//err := templates.ExecuteTemplate(w, tmpl+".html", p)
	err := templates[tmpl+".html"].ExecuteTemplate(w, "base", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (p *Page) save() error {
	filename := "data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func main() {
	for _, r := range routes {
		http.HandleFunc(r.Path, r.Func)
	}
	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Println("Error during HTTP Server initialization", err)
	}
}
