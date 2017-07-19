package main

import (
	"testing"
)

func TestLoadExistingPage(t *testing.T) {
	title := "TestPage"
	p, err := loadPage(title)
	if err != nil {
		t.Fatal(err)
	}
	if p.Title != "TestPage" {
		t.Fatalf("expected TestPage, got %s", p.Title)
	}
}

func TestLoadNonExistingPage(t *testing.T) {
	title := "DOESNOTEXIST"
	_, err := loadPage(title)
	if err == nil {
		t.Fatalf("expecting error, received %s", err)
	}
}

func TestGetAllPagesNoError(t *testing.T) {
	// getAllPageLinks will fill in the []Pages struct slice.
	p, err := getAllPageLinks(DataDirectory)
	if err != nil {
		t.Fatalf("expected nil, received %s", err)
	}
	if len(p) < 1 {
		t.Fatalf("expected pages, received 0: %d", len(p))
	}
}

func TestGetAllPagesWithError(t *testing.T) {
	// getAllPageLinks will fill in the []Pages struct slice.
	_, err := getAllPageLinks("BREAKIT/")
	if err == nil {
		t.Fatalf("expected error, received nil - %s", err)
	}
}

func TestPageInPagesTrue(t *testing.T) {
	// getAllPageLinks will fill in the []Pages struct slice.
	_, err := getAllPageLinks(DataDirectory)
	if err != nil {
		t.Fatalf("expected pages, received error response. %s", err)
	}

	p, _ := loadPage("TestPage")
	if !pageInPages(p) {
		t.Fatalf("expected true, received false. %s", p.Title)
	}
}
