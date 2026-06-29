package currency

import "math"

type Converter struct {
	api ICurrencyConvertionAPI
}

func NewConverter(api ICurrencyConvertionAPI) *Converter {
	return &Converter{api: api}
}

func (c *Converter) ConvertTo(toCurrencyCode string, totalInCent int, fromCurrencyCode string) (int, error) {
	if fromCurrencyCode == toCurrencyCode {
		return totalInCent, nil
	}
	rate, err := c.api.GetExchangeRate(fromCurrencyCode, toCurrencyCode)
	if err != nil {
		return 0, err
	}
	return int(math.Round(float64(totalInCent) * rate)), nil
}
