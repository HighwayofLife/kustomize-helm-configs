package main

import "testing"

func TestMain(t *testing.M) {
	logger = InitLogger()
}

func TestInitLogger(t *testing.T) {
	if logger == nil {
		t.Fatal("expected logger to never be nil")
	}
}
