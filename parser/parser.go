package parser

import (
	"errors"

	"github.com/levilovelock/mtgoscraper/magic"
)

// Parser A parser of webpages to pull decklists
type Parser struct{}

// ParseEvent is a function that will parse an entire magic event
// from the MTGO website, given a url
func (p *Parser) ParseEvent(url string) (*magic.Event, error) {
	return nil, errors.New("ParseEvent has not been implemented yet!")
}
