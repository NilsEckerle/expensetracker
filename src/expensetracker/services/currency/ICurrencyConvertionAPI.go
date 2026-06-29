package currency

type ICurrencyConvertionAPI interface {
	GetExchangeRate(from, to string) (float64, error)
}
