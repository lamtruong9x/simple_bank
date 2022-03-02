package util

const (
	USD = "USD"
	CAD = "CAD"
	EUR = "EUR"
)

func IsSupportedCurrency(c string) bool {
	switch c {
	case USD, CAD, EUR:
		return true
	}
	return false
}