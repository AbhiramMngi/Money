package money

type Money struct {
	paise int
}

func NewMoney(paise int) Money {
	return Money{paise: paise}
}
