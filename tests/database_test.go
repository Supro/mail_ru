package tests

import (
	"github.com/Supro/mail_ru"
	"github.com/Supro/mail_ru/database"
)

func buildDatabase() *database.Database {
	url := "http://google.com"

	ls := make(map[string]*mail_ru.Link, 1)

	l := &mail_ru.Link{url, 1}

	ls[url] = l

	return &database.Database{Links: ls}
}
