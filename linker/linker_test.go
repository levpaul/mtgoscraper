package linker

import "testing"

// This test would be better if Linker was decomposed so that the web request was option/configurable
func TestGetLatestLinks(t *testing.T) {
	lkr := new(Linker)
	results, err := lkr.GetLatestLinks()
	if err != nil {
		t.Error(err)
	}
	if len(results) == 0 {
		t.Errorf("No results returned from linker")
	}
}
