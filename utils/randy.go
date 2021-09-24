package utils

import (
	"crypto/rand"
	"math/big"
)

func Roll(sides, amount int) []int {
	out := []int{}
	for i := 0; i < amount; i++ {
		r, _ := rand.Int(rand.Reader, big.NewInt(int64(sides)))
		out = append(out, int(r.Int64())+1)
	}
	return out
}
