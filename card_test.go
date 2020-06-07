package harvest

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCard(t *testing.T) {
	Convey("Value Test", t, func() {
		tests := []struct {
			C             Card
			ExpectedValue int
		}{
			{C: Card{Rank: CardRankA}, ExpectedValue: CardValA},
			{C: Card{Rank: CardRank2}, ExpectedValue: CardVal2},
			{C: Card{Rank: CardRank3}, ExpectedValue: CardVal3},
			{C: Card{Rank: CardRank4}, ExpectedValue: CardVal4},
			{C: Card{Rank: CardRank5}, ExpectedValue: CardVal5},
			{C: Card{Rank: CardRank6}, ExpectedValue: CardVal6},
			{C: Card{Rank: CardRank7}, ExpectedValue: CardVal7},
			{C: Card{Rank: CardRank8}, ExpectedValue: CardVal8},
			{C: Card{Rank: CardRank9}, ExpectedValue: CardVal9},
			{C: Card{Rank: CardRank10}, ExpectedValue: CardVal10},
			{C: Card{Rank: CardRankJ}, ExpectedValue: CardValJ},
			{C: Card{Rank: CardRankQ}, ExpectedValue: CardValQ},
			{C: Card{Rank: CardRankK}, ExpectedValue: CardValK},
		}
		// run tests
		for _, test := range tests {
			So(test.C.Value(), ShouldEqual, test.ExpectedValue)
		}

		// test panic scenario
		So(func() { (&Card{}).Value() }, ShouldPanicWith, invalidCard)
	})
}
