package card

import "github.com/MirzovalievShodmon/bank/pkg/bank/types"

// Withdraw снимает деньги с карты.
func Withdraw(card *types.Card, amount types.Money) {
	const withDrawLimit = 20_000_00
	if amount < 0 {
		return
	}
	if amount > withDrawLimit {
		return
	}
	if !card.Active {
		return
	}
	if card.Balance < amount {
		return
	}
	card.Balance -= amount
}
