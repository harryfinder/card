package card_test

import (
	"GitHub/pkg/bank/card"
	"GitHub/pkg/bank/card/types"
	"fmt"
)

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			ID:         12345,
			PAN:        "5058 xxxx xxxx 8888",
			Balance:    3_000_00,
			MinBalance: 1_000_00,
			Currency:   "TJS",
			Color:      "White",
			Name:       "Visa",
			Active:     false,
		},
		{
			ID:         12345,
			PAN:        "5058 xxxx xxxx 8888",
			Balance:    8_000_00,
			MinBalance: 2_000_00,
			Currency:   "TJS",
			Color:      "White",
			Name:       "Visa",
			Active:     true,
		},
		{
			ID:         12345,
			PAN:        "5058 xxxx xxxx 8888",
			Balance:    0_000_00,
			MinBalance: 0_000_00,
			Currency:   "TJS",
			Color:      "White",
			Name:       "Visa",
			Active:     true,
		},
		{
			ID:         12345,
			PAN:        "5058 xxxx xxxx 8888",
			Balance:    0_000_00,
			MinBalance: 0_000_00,
			Currency:   "TJS",
			Color:      "White",
			Name:       "Visa",
			Active:     false,
		},
	}
	fmt.Println(card.PaymentSources(cards))

	//output:
}
