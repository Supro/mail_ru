package tests

import (
	"bytes"
	"sync"
	"testing"

	"github.com/Supro/mail_ru"
	"github.com/Supro/mail_ru/mocks"
)

func TestWoker(t *testing.T) {
	reader := bytes.NewBufferString("http://google.com\nhttp://yandex.ru\nhttp://mail.ru\nhttp://godaddy.com\n")
	//done := make(chan bool)

	w := &mail_ru.Worker{Match: "go", Limit: 1, CountChan: make(chan int)}
	wg := &sync.WaitGroup{}

	wg.Add(4)

	ls := mocks.LinkService{
		CreateFunc: func(l *mail_ru.Link) error {
			defer wg.Done()

			if w.Count > 1 {
				t.Errorf("Worker current count can't be greater than limit")
			}

			return nil
		},
	}

	w.LinkService = ls

	go w.Process(reader)

	wg.Wait()
}
