package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Supro/mail_ru"
	"github.com/Supro/mail_ru/database"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	db := &database.Database{Links: make(map[string]*mail_ru.Link)}
	ls := database.LinkService{db}

	w := mail_ru.NewWorker()

	w.Match = "go"
	w.Limit = 5
	w.LinkService = ls

	w.Process(reader)

	for _, l := range db.Links {
		fmt.Printf("Count for %v: %v \n", l.Url, l.Matches)
	}

	fmt.Printf("Total: %v \n", ls.TotalMatches())
}
