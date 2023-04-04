package types

// Money представляет собой денежную сумму в минимальных единицах (центы, копейкы, дирамы и т.д.)
type Money int64

// Currency представляет собой денежную единицу
type Currency string

// коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN представляет номер карты
type PAN string

// Card представдяет информацию о платежной карте
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

// Payment представляет информацию о платеже
type Payment struct {
	ID     int
	Amount Money
}

// PaymentSource представяет из себя источник оплаты
type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}
