package main

import (
	"fmt"

	"github.com/solid-adventure/pocketchain/pkg/core"
)

func main() {
	// Create a new blockchain instance
	bc := core.NewBlockchain()

	// Add a new block to the blockchain
	bc.AddBlock("Send 1 BTC to Ivan")

	// Print all the blocks in the blockchain
	for _, block := range bc.GetBlocks() {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
