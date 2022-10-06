package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	var sequences = [10]int64{8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096}
	for i := 0; i < 10; i++ {
		possible_key := possible_keys(sequences[i])
		rand_key := random_key(possible_key)

		fmt.Printf("Keys for %dbit: %d\n", sequences[i], possible_key)
		fmt.Printf("Random key: %d in hex: 0x%X\n", rand_key, rand_key)
		fmt.Printf("Brute-force time: %s\n", brute_force(rand_key))
		println("-----------------------------------")
	}
}

func possible_keys(bit int64) *big.Int {
	var bigBit = big.NewInt(bit)
	bigBit.Exp(big.NewInt(2), bigBit, nil)
	return bigBit
}

func random_key(max *big.Int) *big.Int {
	res, err := rand.Int(rand.Reader, max.Sub(max, big.NewInt(1)))
	if err != nil {
		return big.NewInt(0)
	}
	return res
}

func brute_force(key *big.Int) string {
	start_time := time.Now()
	i := big.NewInt(0)
	for key.Cmp(i) != 0 {
		i.Add(i, big.NewInt(1))
	}
	t := time.Now()
	return t.Sub(start_time).String()
}
