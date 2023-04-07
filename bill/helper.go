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

func GetVat(amount float64, withVat bool, withCityTax bool) string {
	if withVat {
		if withCityTax {
			return FormatNumber((amount / 111) * 10)
		} else {
			return FormatNumber((amount / 110) * 10)
		}
	} else {
		return "0.00"
	}

}

func GetWithCityTax(amount float64, withVat bool, withCityTax bool) string {
	if withVat {
		if withCityTax {
			return FormatNumber((amount / 111) * 1)
		} else {
			return "0.00"
		}
	} else {
		if withCityTax {
			return FormatNumber((amount / 101) * 1)
		} else {
			return "0.00"
		}
	}
}
