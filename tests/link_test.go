package tests

import (
	"testing"

	"github.com/Supro/mail_ru"
)

func TestLinkCalculateMatches(t *testing.T) {
	tests := []struct {
		Url    string
		Sub    string
		Expect int
	}{
		{"http://google.com", "go", 1},
		{"http://gogodaddy.com", "go", 2},
		{"http://ya.ru", "go", 0},
		{"http://mail.ru", "il", 1},
	}

	for _, ex := range tests {
		l := &mail_ru.Link{Url: ex.Url}

		got := l.CalculateMatches(ex.Sub)

		if got != ex.Expect {
			t.Error("Expected submatches of string to be %v, got %v", ex.Expect, got)
		}
	}
}

func TestLinkSetMatches(t *testing.T) {
	l := &mail_ru.Link{Url: "http://google.com"}

	expect := 1

	l.SetMatches("go")

	got := l.Matches

	if got != expect {
		t.Error("Expected to matches will be setted, got %v expect %v", got, expect)
	}
}
