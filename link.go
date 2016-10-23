package mail_ru

type Link struct {
	Url   string
	Count int
}

type LinkService interface {
	FindByUrl(string) (*Link, error)
	Create(*Link) error
	Update(*Link) error
}
