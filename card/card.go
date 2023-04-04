package card

import (
	"GitHub/pkg/bank/card/types"
)

func PaymentSources(cards []types.Card) []types.PaymentSource {
	payment := make([]types.PaymentSource, 10)
	j := 0

	for _, operations := range cards {
		if !operations.Active || operations.Balance <= 0 {
			continue
		}

		payment[j].Balance = operations.Balance
		j++
	}

	return payment
}
