package database

import (
	"fmt"

	"github.com/Supro/mail_ru"
)

type LinkService struct {
	*Database
}

func (ls LinkService) FindByUrl(url string) (*mail_ru.Link, error) {
	l, ok := ls.Database.Links[url]

	if ok {
		return l, nil
	}

	return l, fmt.Errorf("Link with url %v not found", url)
}

func (ls LinkService) Create(l *mail_ru.Link) error {
	ls.Database.Links[l.Url] = l

	return nil
}

func (ls LinkService) Update(l *mail_ru.Link) error {
	return ls.Create(l)
}

func (ls LinkService) TotalMatches() int {
	var total int

	for _, l := range ls.Database.Links {
		total += l.Matches
	}

	return total
}
