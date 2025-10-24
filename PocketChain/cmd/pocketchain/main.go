package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"pocketchain-module/PocketChain/pkg/blockchain"
)

func main() {
	// Create a new blockchain instance from the genesis file
	bc, err := blockchain.NewBlockchainFromGenesis("config/genesis.json")
	if err != nil {
		log.Fatalf("Failed to create blockchain from genesis: %v", err)
	}

	// Print the hash of the genesis block
	genesisBlock := bc.GetBlocks()[0]
	genesisHash, err := genesisBlock.Hash()
	if err != nil {
		log.Fatalf("Failed to hash genesis block: %v", err)
	}

	fmt.Printf("Genesis block hash: %x\n", genesisHash)
	fmt.Println("Pocketchain node running. Press Ctrl+C to exit.")

	// Wait for a shutdown signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	fmt.Println("\nShutting down pocketchain node...")
}
