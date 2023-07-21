package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RadomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomStrings generate a random string of len n
func RandomString(n int) string {
	var sb strings.Builder
	count := len(alphabet)

	for i := 0; i < n; i++ {
		byte := alphabet[rand.Intn(count)]
		sb.WriteByte(byte)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	count := len(currencies)

	return currencies[rand.Intn(count)]
}
