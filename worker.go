package mail_ru

import (
	"bufio"
	"io"
)

type Worker struct {
	// String for string submatching count
	Match string

	// Limit of goroutins
	Limit int

	// Count of current working routines
	Count int

	// Chan for queueing
	CountChan chan int

	// Strategy for wrinting link data
	// into database source
	LinkService
}

func NewWorker() *Worker {
	w := &Worker{}

	return w
}

// Starts to scann each line of source
// with queueing max goroutines
func (w *Worker) Process(src io.Reader) {
	s := bufio.NewScanner(src)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		if w.Count == w.Limit {
			for <-w.CountChan == w.Limit {
			}
		}

		url := s.Text()
		w.Count += 1
		go w.Calculate(url)
	}

	w.Wait()
}

// Calculates matches in incoming url
// and saves it to database then inform
// that goroutin is done
func (w *Worker) Calculate(url string) {
	l := &Link{Url: url}

	l.SetMatches(w.Match)

	w.LinkService.Create(l)

	w.Done()
}

// Inform count channel that one gouruting is done
func (w *Worker) Done() {
	w.Count--
	w.CountChan <- w.Count
}

// Waits untill all goroutines will completed
func (w *Worker) Wait() {
	for <-w.CountChan > 0 {
	}

	close(w.CountChan)
}
