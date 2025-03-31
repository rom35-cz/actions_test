package main

import (
	"testing"
	"time"
)

func TestYesterday(t *testing.T) {
	expected := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	actual := yesterday().Format("2006-01-02")

	if actual != expected {
		t.Errorf("Expected yesterday's date to be %s, but got %s", expected, actual)
	}
}
