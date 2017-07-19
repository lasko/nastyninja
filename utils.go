package main

import "io/ioutil"

func loadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
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

func getAllPageLinks() ([]Page, error) {
	fileInfo, err := ioutil.ReadDir("data/")
	if err != nil {
		return nil, err
	}
	for _, v := range fileInfo {
		name := v.Name()
		//p := Page{Title: name[0 : len(name)-4], Body: []byte("")}
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
