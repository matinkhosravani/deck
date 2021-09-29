package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: King, Suit: Heart})
	// Output:
	// King of Hearts
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Error("Cards are not enough")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	if cards[0].Rank != Ace && cards[0].Suit != Spade {
		t.Error("Sort is not ok ")
	}
}

func TestCustomSort(t *testing.T) {
	cards := New(CustomSort(CustomLess))
	if cards[1].Rank != Ace && cards[1].Suit != Diamond {
		t.Error("Custome Sort is not ok ")
	}
}

func TestDefaultShuffle(t *testing.T) {
	cards := New(DefaultShuffle)
	if len(cards) != 52 {
		t.Error("Cards are not enough")
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(2))
	count := 0
	for _, card := range cards {
		if card.Suit == Joker {
			count++
		}
	}

	if count != 2 {
		t.Errorf("Expected 2 Jokers , got %d", count)
	}

}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two {
			t.Error("Twos should not be present in filtered cards")
		}
	}

}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))

	if len(cards) != 3*52 {
		t.Errorf("expected %d got %d", 3*52,len(cards))
	}
}
