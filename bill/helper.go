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
