package util

var currencies = []string{
	"USD",
	"EUR",
	"CAD",
}

func IsSupportedCurrency(currency string) bool {
	for _, x := range currencies {
		if x == currency {
			return true
		}
	}
	return false
}
