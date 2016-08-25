package linker

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Linker An object which will attain links to event list urls
type Linker struct{}

// GetLatestLinks will search the main page of mtgo and return links
// the latest events posted there
func (l *Linker) GetLatestLinks() ([]string, error) {
	resp, err := http.Get("http://magic.wizards.com/section-articles-see-more-ajax?l=en&f=9041&fromDate=&toDate=&event_format=0&sort=DESC&word=&offset=0")
	if err != nil {
		return nil, err
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, err
	}

	type WizardsGetLatest struct {
		Status int
		Data   []string
	}

	wizardsGetLatest := new(WizardsGetLatest)
	jsonErr := json.Unmarshal(body, wizardsGetLatest)
	if jsonErr != nil {
		return nil, jsonErr
	}

	var urls []string

	// Grab all of the event links from the data of the latest
	for _, d := range wizardsGetLatest.Data {
		z := html.NewTokenizer(strings.NewReader(d))

		var link string
		for {
			tt := z.Next()
			if tt == html.ErrorToken {
				break
			}
			if tt == html.StartTagToken {
				t := z.Token()
				if t.Data == "a" {
					// Found link
					link = t.Attr[0].Val
				}
			}
		}

		if link != "" {
			urls = append(urls, "http://magic.wizards.com"+link)
		}
	}

	if len(urls) > 0 {
		return urls, nil
	}

	return nil, errors.New("No urls found, something went wrong with the linker!")
}
