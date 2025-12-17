package wargame

import (

)

type (
	// Details about the players deck.
	Deck struct{
		Code string
		Nation uint16
		Specialization uint8
		Era uint8
	}
)
