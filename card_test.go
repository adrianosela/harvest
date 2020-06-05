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
			{C: Card{}, ExpectedValue: -1},
		}
		// run tests
		for _, test := range tests {
			So(test.C.Value(), ShouldEqual, test.ExpectedValue)
		}
	})

	Convey("Flip Test", t, func() {
		tests := []struct {
			C                     Card
			BeforeFlipFaceUp      bool
			AfterFlipExpectFaceUp bool
		}{
			{
				C: Card{
					Rank:   CardRankA,
					FaceUp: false,
				},
				BeforeFlipFaceUp:      false,
				AfterFlipExpectFaceUp: true,
			},
			{
				C: Card{
					Rank:   CardRank2,
					FaceUp: true,
				},
				BeforeFlipFaceUp:      true,
				AfterFlipExpectFaceUp: false,
			},
		}
		// run tests
		for _, test := range tests {
			So(test.C.FaceUp, ShouldEqual, test.BeforeFlipFaceUp)
			test.C.Flip()
			So(test.C.FaceUp, ShouldEqual, test.AfterFlipExpectFaceUp)
		}
	})

}
