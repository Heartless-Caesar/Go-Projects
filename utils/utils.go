package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmonpqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) //Retorna num aleatorio entre minimo e maximo
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(100, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "R$"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
