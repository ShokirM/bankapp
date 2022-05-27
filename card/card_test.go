package card

import (
	"gitproject3/types"
	"fmt"
)
func ExampleTotal() {
	cards := []types.Card {
		{
		CardBalance: 1_000_00,
		CardActivity: true,
		},
		{
			CardBalance: 1_000_00,
			CardActivity: true,
		},
}
	
	fmt.Println(Total(cards)) 

	// Output: 200000
}
