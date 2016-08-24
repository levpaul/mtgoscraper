package linker

import "errors"

// Linker An object which will attain links to event list urls
type Linker struct{}

// GetLatestLinks will search the main page of mtgo and return links
// the latest events posted there
func (l *Linker) GetLatestLinks() ([]string, error) {
	return nil, errors.New("Not implemented yet")
}
