package harvest

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPlayer(t *testing.T) {
	Convey("Test Compute Hand Value", t, func() {
		p := &Player{
			Hand: [CardsPerPlayer]Card{
				{Rank: CardRankA},
				{Rank: CardRank2},
				{Rank: CardRank10},
				{Rank: CardRankJ},
				{Rank: CardRank5},
				{Rank: CardRank7},
			},
		}
		sum := CardValA + CardVal2 + CardVal10 + CardValJ + CardVal5 + CardVal7
		So(p.ComputeScore(), ShouldEqual, sum)
	})
}
