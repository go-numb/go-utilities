package cards_test

import (
	"fmt"
	"sort"
	"testing"
	"time"

	"math/rand"

	"github.com/stretchr/testify/assert"
)

func TestNewCards(t *testing.T) {
	cardN := 52
	isJocker := true
	if isJocker {
		cardN = 54
	}
	cards := b21.NewCards(isJocker)

	sort.Sort(cards)

	var (
		n             int
		s, c, d, h, j int
	)
	for i, v := range *cards {
		switch v.Color {
		case b21.Clover:
			s++
		case b21.Spade:
			c++
		case b21.Diamond:
			d++
		case b21.Heart:
			h++
		default:
			j++
		}

		fmt.Printf("%d	%+v\n", i, v)
		n++
	}

	assert.Equal(t, n, cardN, nil)
	assert.Equal(t, s, 13, nil)
	assert.Equal(t, c, 13, nil)
	assert.Equal(t, d, 13, nil)
	assert.Equal(t, h, 13, nil)
	assert.Equal(t, j, 2, nil)

	sourse := rand.NewSource(time.Now().UnixNano())
	r := rand.New(sourse)

	for i := 0; i < cardN; i++ {
		card := cards.Draw(r)
		fmt.Printf("%+v\n", card)
		fmt.Printf("%+v\n", cards.Len())
	}

}
