package bill

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateBillIdSuffix() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
func FormatNumber(n float64) string {
	return fmt.Sprintf("%.2f", n)
}

func GetVat(amount float64, withCityTax bool) string {
	if withCityTax {
		return FormatNumber((amount / 111) * 10)
	} else {
		return FormatNumber((amount / 110) * 10)
	}
}

func GetWithCityTax(amount float64, witVat bool) string {
	if witVat {
		return FormatNumber((amount / 111) * 1)
	} else {
		return FormatNumber((amount / 101) * 1)
	}
}
