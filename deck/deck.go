package deck

import "math/rand"

// Card represents a playing card, type is the face value ace to king
type Card struct {
	Type string
	Suit string
}

// New will generate a new deck of cards
func New() ([]Card, error) {
	var ret []Card
	joker := Card{"joker", "joker"}
	types := []string{"ace", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "jack", "queen", "king"}

	for _, t := range types {
		d := Card{t, "diamonds"}
		c := Card{t, "clubs"}
		s := Card{t, "spades"}
		h := Card{t, "hearts"}
		ret = append(ret, d, c, s, h)
	}

	ret = append(ret, joker, joker)

	return ret, nil
}

// Shuffle will shuffle the deck into random order
func Shuffle(deck []Card) {
	n := len(deck)
	for i := n - 1; i > 0; i-- {
		r := rand.Intn(i + 1)
		deck[r], deck[i] = deck[i], deck[r]
	}
}
