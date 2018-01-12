package main

import (
	"fmt"
	"strconv"
)

func main() {
	blockchain := NewBlockchain()

	blockchain.AddBlock("1 hackcoin to Hendrik")
	blockchain.AddBlock("1 hackcoin to Peter")

	for _, block := range blockchain.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
