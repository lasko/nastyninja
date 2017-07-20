package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
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

func TestMakeHandler(t *testing.T) {
	hf := makeHandler(viewHandler)
	out := reflect.TypeOf(hf).Kind()
	if out.String() != "func" {
		t.Errorf("got %v, wanted func", out.String())
	}
}

func TestHandlers(t *testing.T) {
	testTable := map[string]struct {
		Method        string
		Path          string
		HandlerMethod http.HandlerFunc
		RespCode      int
	}{
		"View": {
			Method:        "GET",
			Path:          "/view/TestPage",
			HandlerMethod: makeHandler(viewHandler),
			RespCode:      http.StatusOK,
		},
		"Edit": {
			Method:        "GET",
			Path:          "/edit/Test",
			HandlerMethod: makeHandler(editHandler),
			RespCode:      http.StatusOK,
		},
		"NewView": {
			Method:        "GET",
			Path:          "/view/ThisDoesNotExist",
			HandlerMethod: makeHandler(viewHandler),
			RespCode:      http.StatusFound,
		},
		"Save": {
			Method:        "POST",
			Path:          "/edit/ThisDoesNotExist",
			HandlerMethod: makeHandler(saveHandler),
			RespCode:      http.StatusFound,
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
			t.Error(err)
		}
		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(test.HandlerMethod)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != test.RespCode {
			t.Errorf("handler %v returned wrong status code: got %v want %v", test.Path, status, test.RespCode)
		}
		if err := os.Remove(DataDirectory + p.Title + ".txt"); err != nil {
			t.Errorf("expected nil for file removal, received %s", err)
		}
		if _, err := os.Stat(DataDirectory + "ThisDoesNotExist.txt"); !os.IsNotExist(err) {
			_ = os.Remove(DataDirectory + "ThisDoesNotExist.txt")
		}
	}
}
