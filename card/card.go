package card

import (
	"gitproject3/types"
)

func Total(cards []types.Card) types.Money {
	sum := types.Money(0)
	for _, card := range cards {
		if !card.CardActivity {
			continue
		}
		if card.CardBalance <= 0{
			continue
		}
		sum += types.Money(card.CardBalance)
			}
	return types.Money(sum) 
}	