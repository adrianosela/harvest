package deck

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
			{
				C: Card{
					Type: CardA,
				},
				ExpectedValue: CardValA,
			},
			{
				C: Card{
					Type: Card2,
				},
				ExpectedValue: CardVal2,
			},
			{
				C: Card{
					Type: Card3,
				},
				ExpectedValue: CardVal3,
			},
			{
				C: Card{
					Type: Card4,
				},
				ExpectedValue: CardVal4,
			},
			{
				C: Card{
					Type: Card5,
				},
				ExpectedValue: CardVal5,
			},
			{
				C: Card{
					Type: Card6,
				},
				ExpectedValue: CardVal6,
			},
			{
				C: Card{
					Type: Card7,
				},
				ExpectedValue: CardVal7,
			},
			{
				C: Card{
					Type: Card8,
				},
				ExpectedValue: CardVal8,
			},
			{
				C: Card{
					Type: Card9,
				},
				ExpectedValue: CardVal9,
			},
			{
				C: Card{
					Type: Card10,
				},
				ExpectedValue: CardVal10,
			},
			{
				C: Card{
					Type: CardJ,
				},
				ExpectedValue: CardValJ,
			},
			{
				C: Card{
					Type: CardQ,
				},
				ExpectedValue: CardValQ,
			},
			{
				C: Card{
					Type: CardK,
				},
				ExpectedValue: CardValK,
			},
			{
				C: Card{
					Type: "NonExistentType",
				},
				ExpectedValue: -1,
			},
		}
		// run tests
		for _, test := range tests {
			So(test.C.Value(), ShouldEqual, test.ExpectedValue)
		}
	})

	Convey("FlipUp Test", t, func() {
		tests := []struct {
			C                     Card
			BeforeFlipFaceUp      bool
			AfterFlipExpectFaceUp bool
		}{
			{
				C: Card{
					Type:   CardA,
					FaceUp: false,
				},
				BeforeFlipFaceUp:      false,
				AfterFlipExpectFaceUp: true,
			},
			{
				C: Card{
					Type:   Card2,
					FaceUp: true,
				},
				BeforeFlipFaceUp:      true,
				AfterFlipExpectFaceUp: true,
			},
		}
		// run tests
		for _, test := range tests {
			So(test.C.FaceUp, ShouldEqual, test.BeforeFlipFaceUp)
			test.C.FlipUp()
			So(test.C.FaceUp, ShouldEqual, test.AfterFlipExpectFaceUp)
		}
	})

}
