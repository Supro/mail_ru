package mocks

type LinkService struct {
	FindByUrlFunc func(string) (*mail_ru.Link, error)
	CreateFunc func(*mail_ru.Link) error
}

func(ls LinkService) FindByUrl(url string) (*mail_ru.Link, error) {
	return ls.FindByUrlFunc(url)
}

func(ls LinkService) Create(l *mail_ru.Link) error {
	return ls.CreateFunc(l)
}

func(ls LinkService) Update(l *mail_ru.Link) error {
	return ls.CreateFunc(l)
}
