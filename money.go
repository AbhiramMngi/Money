package money

type Money struct {
	paise int
}

func NewMoney(rupees, paise int) Money {
	if paise < 0 || rupees < 0 {
		panic("paise should be non negative")
	}
	paise += rupees * 100
	return Money{paise: paise}
}

func (money *Money) AddMoney(moneyToAdd Money) {
	if moneyToAdd.paise == 0 {
		return
	}
	if money.paise == 0 {
		money.paise = moneyToAdd.paise
		return
	}
	money.paise += moneyToAdd.paise
}

func (money *Money) SubtractMoney(moneyToSubtract Money) {
	if money.paise == 0 {
		panic("balance is 0 cannot deduct")
	}
	if moneyToSubtract.paise == 0 {
		return
	}
	if moneyToSubtract.paise > money.paise {
		panic("balance drops below 0")
	}
	money.paise -= moneyToSubtract.paise
}
