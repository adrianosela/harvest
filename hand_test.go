package harvest

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHand(t *testing.T) {
	Convey("Test Compute Hand Value", t, func() {
		h := &Hand{
			Card{Rank: CardRankA},
			Card{Rank: CardRank2},
			Card{Rank: CardRank10},
			Card{Rank: CardRankJ},
			Card{Rank: CardRank5},
			Card{Rank: CardRank7},
		}
		sum := CardValA + CardVal2 + CardVal10 + CardValJ + CardVal5 + CardVal7
		So(h.ComputeScore(), ShouldEqual, sum)
	})
}
