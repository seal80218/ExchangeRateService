package model

type ExchangeRateDao struct {
	Uid             int64   `json:"uid,omitempty" xml:"uid"`
	Name            string  `json:"name,omitempty" xml:"name"`
	CashBuyingRate  float32 `json:"cash_buying_rate,omitempty" xml:"cash_buying_rate"`
	CashSellingRate float32 `json:"cash_selling_rate,omitempty" xml:"cash_selling_rate"`
	SignBuyingRate  float32 `json:"sign_buying_rate,omitempty" xml:"sign_buying_rate"`
	SignSellingRate float32 `json:"sign_selling_rate,omitempty" xml:"sign_selling_rate"`
	TimeStamp       string  `json:"time_stamp,omitempty" xml:"time_stamp"`
}
