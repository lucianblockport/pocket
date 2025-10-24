package blockchain

import (
	"log"
	"time"

	"pocketchain-module/PocketChain/config"
)

// Blockchain represents the structure of a blockchain.
type Blockchain struct {
	Blocks []*Block
	State  *State
}

// NewBlockchainFromGenesis creates a new blockchain from a genesis configuration.
func NewBlockchainFromGenesis(genesisPath string) (*Blockchain, error) {
	genesisConfig, err := config.LoadGenesisConfig(genesisPath)
	if err != nil {
		return nil, err
	}

	genesisTime, err := time.Parse(time.RFC3339, genesisConfig.GenesisTime)
	if err != nil {
		return nil, err
	}

	genesisHeader := &BlockHeader{
		Height:       0,
		Timestamp:    genesisTime.UnixNano(),
		PreviousHash: []byte{},
		StateRoot:    []byte{},
		ExtraData:    []byte(genesisConfig.ExtraData),
	}
	genesisBody := &BlockBody{
		Transactions:      []*Transaction{},
		ValidatorCommittee: [][]byte{},
	}
	genesisBlock := NewBlock(genesisHeader, genesisBody, []byte{})
	bc := &Blockchain{
		Blocks: []*Block{genesisBlock},
		State:  NewState(),
	}
	return bc, nil
}

// AddBlock adds a new block to the blockchain.
func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	prevBlockHash, err := prevBlock.Hash()
	if err != nil {
		log.Printf("Error hashing previous block: %v", err)
		return
	}

	newHeader := &BlockHeader{
		Height:       uint64(len(bc.Blocks)),
		Timestamp:    time.Now().UnixNano(),
		PreviousHash: prevBlockHash,
		StateRoot:    []byte{}, // In a real scenario, this would be the Merkle root of the state.
	}
	newBody := &BlockBody{
		Transactions:      transactions,
		ValidatorCommittee: [][]byte{},
	}
	newBlock := NewBlock(newHeader, newBody, []byte{})
	bc.Blocks = append(bc.Blocks, newBlock)
}

// GetBlocks returns all the blocks in the blockchain.
func (bc *Blockchain) GetBlocks() []*Block {
	return bc.Blocks
}
