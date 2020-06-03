package harvest

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDeck(t *testing.T) {

	Convey("DeckHas52Cards", t, func() {
		d1 := NewDeck()
		So(len(d1.Cards), ShouldEqual, 52)
	})

}
