//go:generate stringer  -type=Suit,Rank
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

const (
	maxRank = King
	minRank = Ace
)

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func CustomSort(less func(cards []Card) func(i, j int) bool) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func CustomLess(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return cards[i].Rank < cards[j].Rank
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

func DefaultShuffle(cards []Card) []Card {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	return cards
}

func Jokers(n int) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Suit: Joker,
				Rank: Rank(i),
			})
		}

		return cards
	}
}

func Filter(f func(card Card) bool) func(cards []Card) []Card {
	var filtered []Card
	return func(cards []Card) []Card {
		for _, c := range cards {
			if !f(c) {
				filtered = append(filtered, c)
			}
		}

		return filtered
	}
}

func New(options ...func([]Card) []Card) []Card {
	var cards []Card
	for s := Spade; s < Joker; s++ {
		for r := minRank; r <= maxRank; r++ {
			card := Card{Rank: r, Suit: s}
			cards = append(cards, card)
		}
	}

	for _, option := range options {
		cards = option(cards)
	}

	return cards
}

func Deck(n int) func(cards []Card) []Card {

	var ret []Card
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}

		return ret
	}

}
