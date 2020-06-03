package player

import (
	"testing"

	"github.com/adrianosela/harvest/game/deck"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPlayer(t *testing.T) {
	Convey("Test Compute Hand Value", t, func() {
		p := Player{
			PlayerID: "PlayerA",
			Hand: [6]deck.Card{
				{
					Type: deck.CardA,
				},
				{
					Type: deck.Card2,
				},
				{
					Type: deck.Card10,
				},
				{
					Type: deck.CardJ,
				},
				{
					Type: deck.Card5,
				},
				{
					Type: deck.Card7,
				},
			},
		}
		sum := deck.CardValA + deck.CardVal2 + deck.CardVal10 + deck.CardValJ + deck.CardVal5 + deck.CardVal7
		So(p.ComputeScore(), ShouldEqual, sum)
	})

}
