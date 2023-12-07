package util

var Currencies = []string{
	"USD",
	"EUR",
	"CAD",
}

func IsSupportedCurrency(currency string) bool {
	for _, x := range Currencies {
		if x == currency {
			return true
		}
	}
	return false
}
