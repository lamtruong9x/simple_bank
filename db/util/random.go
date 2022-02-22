package util

import (
	"math/rand"
	"strings"
	"time"
)
var alphabet = "abcdefghijklmnopqrstuvwxyz"
func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64{
	return min + rand.Int63n(max - min)
}
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}
func RandomString(n int) string {
	var s strings.Builder
	k := len(alphabet)
	for i:=0; i<n; i++ {
		c := alphabet[rand.Intn(k)]
		s.WriteByte(c)
	}
	return s.String()
}
func RandomOwner() string {
	return RandomString(6)
}

func RandomCurrency() string {
	curr := []string{"USD", "EUR", "CAD"}
	return curr[rand.Intn(len(curr))] 
}