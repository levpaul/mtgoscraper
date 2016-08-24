package magic

// Event is a big net for all types of events
// reported by MTGO
type Event struct {
	Title string
	Decks []Deck
}
