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

func TestHandlers(t *testing.T) {
	testTable := map[string]struct {
		Method   string
		Path     string
		Handler  http.Handler
		RespCode int
	}{
		"View": {
			Method:   "GET",
			Path:     "/view/TestPage",
			Handler:  makeHandler(viewHandler),
			RespCode: http.StatusOK,
		},
		"Edit": {
			Method:   "GET",
			Path:     "/edit/Test",
			Handler:  makeHandler(editHandler),
			RespCode: http.StatusFound,
		},
		"EditNew": {
			Method:   "GET",
			Path:     "/edit/ThisDoesNotExist",
			Handler:  makeHandler(editHandler),
			RespCode: http.StatusFound,
		},
		"NewView": {
			Method:   "GET",
			Path:     "/view/ThisDoesNotExist",
			Handler:  makeHandler(viewHandler),
			RespCode: http.StatusFound,
		},
	}
	for _, test := range testTable {
		p := newPage()
		p.Title = "MockTest"
		if err := p.save(); err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(test.Method, test.Path, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(makeHandler(viewHandler))

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != test.RespCode {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
		if err := os.Remove(DataDirectory + p.Title + ".txt"); err != nil {
			t.Errorf("expected nil for file removal, received %s", err)
		}
		if _, err := os.Stat(DataDirectory + "ThisDoesNotExist.txt"); !os.IsNotExist(err) {
			_ = os.Remove(DataDirectory + "ThisDoesNotExist.txt")
		}
	}
}
