package main

import (
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
	defer os.Remove(DataDirectory + p.Title + ".txt")
	if err := p.save(); err != nil {
		t.Fatalf("expected nil on page save, received %s", err)
	}
}
