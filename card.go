package harvest

// Rank and Suit definitions MUST match those in:
// https://www.npmjs.com/package/react-playing-card

const (
	// CardRankA represents an Ace
	CardRankA = "A"
	// CardRank2 represents a Two
	CardRank2 = "2"
	// CardRank3 represents a Three
	CardRank3 = "3"
	// CardRank4 represents a Four
	CardRank4 = "4"
	// CardRank5 represents a Five
	CardRank5 = "5"
	// CardRank6 represents a Six
	CardRank6 = "6"
	// CardRank7 represents a Seven
	CardRank7 = "7"
	// CardRank8 represents an Eight
	CardRank8 = "8"
	// CardRank9 represents a Nine
	CardRank9 = "9"
	// CardRank10 represents a Ten
	CardRank10 = "10"
	// CardRankJ represents a Jack
	CardRankJ = "J"
	// CardRankQ represents a Queen
	CardRankQ = "Q"
	// CardRankK represents a King
	CardRankK = "K"

	// CardValA represents the value of an Ace
	CardValA = 1
	// CardVal2 represents the value of a Two
	CardVal2 = -2
	// CardVal3 represents the value of a Three
	CardVal3 = 3
	// CardVal4 represents the value of a Four
	CardVal4 = 4
	// CardVal5 represents the value of a Five
	CardVal5 = 5
	// CardVal6 represents the value of a Six
	CardVal6 = 6
	// CardVal7 represents the value of a Seven
	CardVal7 = 7
	// CardVal8 represents the value of an Eight
	CardVal8 = 8
	// CardVal9 represents the value of a Nine
	CardVal9 = 9
	// CardVal10 represents the value of a Ten
	CardVal10 = 10
	// CardValJ represents the value of a Jack
	CardValJ = 0
	// CardValQ represents the value of a Queen
	CardValQ = 10
	// CardValK represents the value of a King
	CardValK = 0

	// CardSuitHeart represents the Heart suit
	CardSuitHeart = "H"
	// CardSuitDiamond represents the Diamond suit
	CardSuitDiamond = "D"
	// CardSuitClub represents the Club suit
	CardSuitClub = "C"
	// CardSuitSpade represents the Spade suit
	CardSuitSpade = "S"

	// used by panic in Value()
	invalidCard = "invalid card"
)

var (
	// Types represents the possible cards a player can hold
	Types = []string{CardRankA, CardRank2, CardRank3, CardRank4, CardRank5,
		CardRank6, CardRank7, CardRank8, CardRank9, CardRank10,
		CardRankJ, CardRankQ, CardRankK}
	// Suits represents the possible suits a card can belong to
	Suits = []string{CardSuitHeart, CardSuitDiamond, CardSuitClub, CardSuitSpade}
)

// Card represents the state of card of a particular type and suit
type Card struct {
	Rank         string `json:"rank,omitempty"`
	Suit         string `json:"suit,omitempty"`
	FaceUp       bool   `json:"face_up,omitempty"`
	OwnerVisible bool   `json:"owner_visible,omitempty"`
}

// Value returns the value of a given card
func (c *Card) Value() int {
	switch c.Rank {
	case CardRankA:
		return CardValA
	case CardRank2:
		return CardVal2
	case CardRank3:
		return CardVal3
	case CardRank4:
		return CardVal4
	case CardRank5:
		return CardVal5
	case CardRank6:
		return CardVal6
	case CardRank7:
		return CardVal7
	case CardRank8:
		return CardVal8
	case CardRank9:
		return CardVal9
	case CardRank10:
		return CardVal10
	case CardRankJ:
		return CardValJ
	case CardRankQ:
		return CardValQ
	case CardRankK:
		return CardValK
	default:
		panic(invalidCard)
	}
}
