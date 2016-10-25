package tests

import (
	"bytes"
	"runtime"
	"sync"
	"testing"

	"github.com/Supro/mail_ru"
	"github.com/Supro/mail_ru/mocks"
)

func TestWoker(t *testing.T) {
	reader := bytes.NewBufferString("http://google.com\nhttp://yandex.ru\nhttp://mail.ru\nhttp://godaddy.com\nhttp://godaddy.com")

	w := &mail_ru.Worker{Match: "go", Limit: 1, CountChan: make(chan int)}
	wg := &sync.WaitGroup{}

	var routines int

	wg.Add(5)

	ls := mocks.LinkService{
		CreateFunc: func(l *mail_ru.Link) error {
			defer wg.Done()

			if w.Count > 1 {
				t.Errorf("Worker current count can't be greater than limit")
			}

			rs := runtime.NumGoroutine()

			if rs > routines+1 {
				t.Errorf("Too many routines are working now")
			}

			return nil
		},
	}

	w.LinkService = ls

	go w.Process(reader)

	routines = runtime.NumGoroutine()

	wg.Wait()
}
