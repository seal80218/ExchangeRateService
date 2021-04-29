package model

import "strconv"

type ExchangeRate struct {
	Name            string  `json:"name,omitempty" xml:"name"`
	CashBuyingRate  float32 `json:"cash_buying_rate,omitempty" xml:"cash_buying_rate"`
	CashSellingRate float32 `json:"cash_selling_rate,omitempty" xml:"cash_selling_rate"`
	SignBuyingRate  float32 `json:"sign_buying_rate,omitempty" xml:"sign_buying_rate"`
	SignSellingRate float32 `json:"sign_selling_rate,omitempty" xml:"sign_selling_rate"`
}

func NewExchangeRate(name string, cashBuyingRate string, cashSellingRate string, signBuyingRate string, signSellingRate string) ExchangeRate {
	cb, _ := strconv.ParseFloat(cashBuyingRate, 32)
	cs, _ := strconv.ParseFloat(cashSellingRate, 32)
	sb, _ := strconv.ParseFloat(signBuyingRate, 32)
	ss, _ := strconv.ParseFloat(signSellingRate, 32)

	return ExchangeRate{
		Name:            name,
		CashBuyingRate:  float32(cb),
		CashSellingRate: float32(cs),
		SignBuyingRate:  float32(sb),
		SignSellingRate: float32(ss),
	}
}
