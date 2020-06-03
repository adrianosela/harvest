package harvest

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHand(t *testing.T) {
	Convey("Test Compute Hand Value", t, func() {
		h := &Hand{
			Card{Type: CardA},
			Card{Type: Card2},
			Card{Type: Card10},
			Card{Type: CardJ},
			Card{Type: Card5},
			Card{Type: Card7},
		}
		sum := CardValA + CardVal2 + CardVal10 + CardValJ + CardVal5 + CardVal7
		So(h.ComputeScore(), ShouldEqual, sum)
	})
}
