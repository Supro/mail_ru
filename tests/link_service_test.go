package tests

import (
	"testing"

	"github.com/Supro/mail_ru"
	"github.com/Supro/mail_ru/database"
)

func TestLinkServiceFindByUrl(t *testing.T) {
	db := buildDatabase()

	s := database.LinkService{db}

	url := "http://google.com"

	l, err := s.FindByUrl(url)

	if err != nil {
		t.Error(err)
	}

	expect := url
	got := l.Url

	if got != expect {
		t.Errorf("Expected to find correct url, got %v, searching for %v", got, expect)
	}

	l, err = s.FindByUrl("http://ya.ru")

	if err == nil {
		t.Error("Shouldn't find link")
	}
}

func TestLinkServiceCreate(t *testing.T) {
	db := buildDatabase()

	s := database.LinkService{db}

	url := "http://ya.ru"

	l := &mail_ru.Link{Url: url, Matches: 0}

	s.Create(l)

	_, err := s.FindByUrl(url)

	if err != nil {
		t.Error(err)
	}
}

func TestLinkServiceUpdate(t *testing.T) {
	db := buildDatabase()

	s := database.LinkService{db}

	url := "http://google.com"

	l, err := s.FindByUrl(url)

	if err != nil {
		t.Error(err)
	}

	expect := 888

	l.Matches = 888

	s.Update(l)

	fl, err := s.FindByUrl(url)

	if err != nil {
		t.Error(err)
	}	

	got := fl.Matches

	if expect != got {
		t.Errorf("Record not updated, v% != %v", expect, got)
	}
}

func TestLinkServiceTotalMatches(t *testing.T) {
	db := buildDatabase()

	s := database.LinkService{db}

	l := &mail_ru.Link{"http://godaddy.com", 1}

	s.Create(l)

	expect := 2

	got := s.TotalMatches()

	if got != expect {
		t.Errorf("Expected total matches to be %v, got %v")
	}
}
