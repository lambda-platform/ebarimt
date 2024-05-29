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
			return FormatNumber((amount / 112) * 10)
		} else {
			return FormatNumber((amount / 110) * 10)
		}
	} else {
		return "0.00"
	}

}

func GetCityTax(amount float64, withVat bool, withCityTax bool) string {
	if withVat {
		if withCityTax {
			return FormatNumber((amount / 112) * 2)
		} else {
			return "0.00"
		}
	} else {
		if withCityTax {
			return FormatNumber((amount / 102) * 2)
		} else {
			return "0.00"
		}
	}
}

func Vat(amount float64, withVat bool, withCityTax bool) float64 {
	if withVat {
		if withCityTax {
			return (amount / 112) * 10
		} else {
			return (amount / 110) * 10
		}
	} else {
		return 0
	}

}

func CityTax(amount float64, withVat bool, withCityTax bool) float64 {
	if withVat {
		if withCityTax {
			return (amount / 112) * 2
		} else {
			return 0
		}
	} else {
		if withCityTax {
			return (amount / 102) * 2
		} else {
			return 0
		}
	}
}
