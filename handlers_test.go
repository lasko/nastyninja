package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestViewHandler(t *testing.T) {
	p := newPage()
	p.Title = "MockTest"
	if err := p.save(); err != nil {
		t.Fatal(err)
	}
	defer os.Remove(DataDirectory + p.Title + ".txt")
	req, err := http.NewRequest("GET", "/view/TestPage", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(makeHandler(viewHandler))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestEditHandler(t *testing.T) {
	p := newPage()
	p.Title = "MockTest"
	if err := p.save(); err != nil {
		t.Fatal(err)
	}
	defer os.Remove(DataDirectory + p.Title + ".txt")
	req, err := http.NewRequest("GET", "/edit/MockTest", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(makeHandler(viewHandler))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
