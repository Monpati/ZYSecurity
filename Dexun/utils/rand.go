package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateRand() int64 {
	randNum, _ := rand.Int(rand.Reader, big.NewInt(1024))
	return randNum.Int64()
}
