package main

import (
	"crypto/rand"
	"math/big"
)

func GetRandomInt(i int) int {
	x, _ := rand.Int(rand.Reader, big.NewInt(int64(i)))
	return int(x.Int64())
}

func GetDiceRoll(i int) int {
	return GetRandomInt(i) + 1
}

func GetRandomBetween(low, high int) int {
	return GetDiceRoll(high-low) + high
}
