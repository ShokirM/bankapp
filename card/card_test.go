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
		{
			CardBalance: 5_00_000,
			CardActivity: true,
		},
		{
			CardBalance: -1_000_00,
			CardActivity: true,
		},
		{
			CardBalance: 15_000_00,
			CardActivity: true,
		},
}
	
	fmt.Println(Total(cards)) 

	// Output: 2200000
}
