package main

import (
	"io/ioutil"
)

// User defines someone using the site. By default a user is NOT an admin.
/*
type User struct {
	Name  string
	ID    int
	Admin bool
}
*/

// DataDirectory defines the location for our data files
var DataDirectory = "data/"

func loadPage(title string) (*Page, error) {
	filename := DataDirectory + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func pageInPages(p *Page) bool {
	for _, j := range Pages {
		if p.Title == j.Title {
			return true
		}
	}
	return false
}

//This is a stub function to be used when user permissions are added.
/*
func isAdmin(u *User) bool {
	if u.Admin {
		return true
	}
	return false
}
*/

func getAllPageLinks(dataDirectory string) ([]Page, error) {
	fileInfo, err := ioutil.ReadDir(dataDirectory)
	if err != nil {
		return nil, err
	}
	for _, v := range fileInfo {
		name := v.Name()
		p := newPage()
		p.Title, p.Body = name[0:len(name)-4], []byte("")

		if !pageInPages(p) {
			Pages = append(Pages, *p)
		}
	}
	return Pages, nil
}

func newPage() *Page {
	return &Page{Title: "", Body: []byte("")}
}
