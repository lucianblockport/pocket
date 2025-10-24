package consensus

import (
	"errors"

	"pocketchain/pkg/blockchain"
)

// Validator represents a validator node.
type Validator struct {
	ID              string
	vrfSecretKey    []byte
	vrfPublicKey    []byte
	uptime          int64
	blocksProduced  int64
	slashingEvents  int64
	batteryLevel    int
}

// NewValidator creates a new validator.
func NewValidator(id string, vrfSecretKey, vrfPublicKey []byte) *Validator {
	return &Validator{
		ID:             id,
		vrfSecretKey:   vrfSecretKey,
		vrfPublicKey:   vrfPublicKey,
	}
}

// IsSelected checks if the validator is selected for the current epoch.
func (v *Validator) IsSelected(seed []byte, stakeWeight uint64, numValidators int) (bool, error) {
	// This is a simplified check. In a real implementation, the validator would
	// need to check against the list of selected validators for the current epoch.
	return true, nil
}

// ProposeBlock proposes a new block.
func (v *Validator) ProposeBlock(header *blockchain.BlockHeader, body *blockchain.BlockBody, vrfProof []byte) (*blockchain.Block, error) {
	// Check if validator is the leader for the current round
	block := blockchain.NewBlock(header, body, vrfProof)
	if err := block.Validate(); err != nil {
		return nil, err
	}
	return block, nil
}

// VoteOnBlock votes on a block.
func (v *Validator) VoteOnBlock(block *blockchain.Block) ([]byte, error) {
	// Sign the block hash with the validator's private key
	blockHash, err := block.Hash()
	if err != nil {
		return nil, err
	}
	// Replace with actual signing logic
	signature := []byte("signature-" + v.ID)
	return signature, nil
}

// SignBlock signs a block.
func (v *Validator) SignBlock(block *blockchain.Block) ([]byte, error) {
	if v.batteryLevel < 20 { // Example battery level check
		return nil, errors.New("battery level too low to sign block")
	}
	return v.VoteOnBlock(block)
}
