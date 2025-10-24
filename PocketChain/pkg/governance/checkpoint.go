package governance

import (
	"errors"

	"pocketchain/pkg/blockchain"
	"pocketchain/pkg/node"
)

// Checkpoint represents a checkpoint of the blockchain state.
type Checkpoint struct {
	BlockID     string
	StateRoot   string
	Signatures  map[string][]byte // authority -> signature
	Authorities []string
}

// NewCheckpoint creates a new checkpoint.
func NewCheckpoint(block *blockchain.Block, authorities []*node.AuthorityNode) *Checkpoint {
	authSigs := make(map[string][]byte)
	authIDs := make([]string, len(authorities))
	for i, auth := range authorities {
		authSigs[auth.ID] = nil
		authIDs[i] = auth.ID
	}

	return &Checkpoint{
		BlockID:     block.ID,
		StateRoot:   block.StateRoot,
		Signatures:  authSigs,
		Authorities: authIDs,
	}
}

// Sign signs the checkpoint with the given authority's signature.
func (c *Checkpoint) Sign(authority *node.AuthorityNode, signature []byte) error {
	if _, ok := c.Signatures[authority.ID]; !ok {
		return errors.New("authority not in checkpoint authority set")
	}
	c.Signatures[authority.ID] = signature
	return nil
}

// VerifySignatures verifies the signatures in the checkpoint.
func (c *Checkpoint) VerifySignatures() (bool, error) {
	// In a real implementation, you would verify the BLS multi-signature.
	return true, nil
}

// IsFinalized checks if the checkpoint is finalized.
func (c *Checkpoint) IsFinalized() bool {
	// A checkpoint is finalized if it has been justified by a supermajority of authorities.
	// This is a simplified check. A real implementation would involve Casper FFG logic.
	return len(c.Signatures) >= (2*len(c.Authorities)/3 + 1)
}
