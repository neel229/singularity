package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	alphabets = "abcdefghijklmnopqrstuvwxyz"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random number
// between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomInt32 generates a random number
// between min and max
func RandomInt32(min, max int64) int32 {
	return int32(min + rand.Int63n(max-min+1))
}

// RandomString generates a random string
// of length of n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomBool returns true if
// rand.Intn returns 1;else false
func RandomBool() bool {
	return rand.Intn(2) == 1
}

// RandomCurrencyCode generates a
// random currency code
func RandomCurrencyCode() string {
	currency_codes := []string{"USD", "ETH", "DAI", "USDC"}
	k := len(currency_codes)
	return currency_codes[rand.Intn(k)]
}

// RandomCurrencyName generates a
// random currency name
func CurrencyName(code string) string {
	switch code {
	case "USD":
		return "United States Dollar"
	case "ETH":
		return "Ether"
	case "DAI":
		return "DAI"
	case "USDC":
		return "USDC"
	default:
		return fmt.Sprint("Currency code not matching any currency")
	}
}

func IsBase(code string) bool {
	switch code {
	case "USD":
		return false
	case "ETH":
		return true
	case "DAI":
		return true
	case "USDC":
		return true
	default:
		return false
	}
}
