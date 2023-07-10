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

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	sb := strings.Builder{}
	alphaLen := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(alphaLen)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner(n int) string {
	return RandomString(n)
}

func RandomMoney() int64 {
	return RandomInt(500, 25000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "CAD", "EGP"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
