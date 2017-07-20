package main

import (
	"net/http/httptest"
	"os"
	"testing"
)

func TestPageSaveEmptyTitle(t *testing.T) {
	p := newPage()
	if err := p.save(); err == nil {
		t.Fatalf("expected error on page save, received %s", err)
	}
}

func TestPageSave(t *testing.T) {
	p := newPage()
	p.Title = "Testing"

	if err := p.save(); err != nil {
		t.Fatalf("expected nil on page save, received %s", err)
	}
	if err := os.Remove(DataDirectory + p.Title + ".txt"); err != nil {
		t.Errorf("expected nil for file removal, received %s", err)
	}
}

func TestRenderTemplate(t *testing.T) {
	index, _ := getAllPageLinks(DataDirectory)
	view, _ := loadPage("TestPage")
	edit, _ := loadPage("TestPage")
	var testPages = []struct {
		Method string
		Path   string
		Tmpl   string
		P      interface{}
	}{
		{"GET", "/", "index", index},
		{"GET", "/view/TestPage", "view", view},
		{"GET", "/edit/TestPage", "edit", edit},
	}
	for _, test := range testPages {
		w := httptest.NewRecorder()
		renderTemplate(w, test.Tmpl, test.P)
	}
}
