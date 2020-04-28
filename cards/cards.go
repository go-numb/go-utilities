package cards

import (
	"math/rand"
)

type Color int

const (
	_ Color = iota
	Clover
	Spade
	Diamond
	Heart
	Jocker
)

func (p Color) String() string {
	switch p {
	case Clover:
		return "♧ CLOVER"
	case Spade:
		return "♤ SPADE"
	case Diamond:
		return "♢ DIAMOND"
	case Heart:
		return "♡ HEART"
	}
	return "💋 JOCKER"
}

type Cards []Card
type Card struct {
	Color  Color
	Number int
}

func (a Cards) Len() int           { return len(a) }
func (a Cards) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Cards) Less(i, j int) bool { return a[i].Number < a[j].Number }

func NewCards(inJocker bool) *Cards {
	n := 52
	if inJocker {
		n = 54
	}

	color := 0
	cards := make([]Card, n)
	for i := 0; i < n; i++ {
		mod13 := (i % 13) + 1
		if mod13 == 1 {
			color++
		}
		if 52 <= i {
			color = int(Jocker)
			mod13 = 0
		}

		cards[i] = Card{
			Color:  Color(color),
			Number: mod13,
		}
	}

	r := Cards(cards)
	return &r
}

func (p *Cards) Draw(r *rand.Rand) Card {
	n := r.Intn(p.Len())
	card := (*p)[n]
	*p = append((*p)[:n], (*p)[n+1:]...)

	return card
}
