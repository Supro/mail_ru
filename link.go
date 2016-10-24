package mail_ru

import "strings"

// Main structure for working with incoming urls
type Link struct {
	Url     string
	Matches int
}

// Returns number of substring matches in link Url
func (l *Link) CalculateMatches(sub string) int {
	return strings.Count(l.Url, sub)
}

// Sets matches of substring in link Url
func (l *Link) SetMatches(sub string) {
	l.Matches = l.CalculateMatches(sub)
}

// Service for working with database source
// for link struct
type LinkService interface {
	// Returns link by it Url string
	FindByUrl(string) (*Link, error)

	// Creates link record in database
	Create(*Link) error

	// Updates link record in database
	Update(*Link) error

	// Returns total matches of all links
	TotalMatches() int
}
