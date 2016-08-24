package main

import (
	"fmt"

	"github.com/levilovelock/mtgoscraper/linker"
	"github.com/levilovelock/mtgoscraper/magic"
	"github.com/levilovelock/mtgoscraper/parser"
)

func main() {
	linker := new(linker.Linker)
	parser := new(parser.Parser)
	events := []magic.Event{}

	links, linkErr := linker.GetLatestLinks()
	if linkErr != nil {
		panic("Could not load the latest links!")
	}

	for _, l := range links {
		e, parseErr := parser.ParseEvent(l)
		if parseErr == nil {
			events = append(events, *e)
		} else {
			fmt.Printf("Error occurred when attempting to parse event from link '%s', error: '%s'\n", l, parseErr.Error())
		}
	}

}
