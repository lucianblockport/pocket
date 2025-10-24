package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"errors"
	"time"
)

// BlockHeader defines the structure of a block header.
type BlockHeader struct {
	Height             uint64
	Timestamp          int64
	PreviousHash       []byte
	StateRoot          []byte
	ValidatorSignature []byte
}

// BlockBody defines the structure of a block body.
type BlockBody struct {
	Transactions      []*Transaction
	ValidatorCommittee [][]byte
}

// Block defines the structure of a block in PocketChain.
type Block struct {
	Header   *BlockHeader
	Body     *BlockBody
	VRFProof []byte
}

// NewBlock creates and returns a new block.
func NewBlock(header *BlockHeader, body *BlockBody, vrfProof []byte) *Block {
	return &Block{
		Header:   header,
		Body:     body,
		VRFProof: vrfProof,
	}
}

// Serialize serializes the block into a byte slice.
func (b *Block) Serialize() ([]byte, error) {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		return nil, errors.New("failed to serialize block: " + err.Error())
	}
	return result.Bytes(), nil
}

// Deserialize deserializes a byte slice into a block.
func Deserialize(data []byte) (*Block, error) {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
if err != nil {
		return nil, errors.New("failed to deserialize block: " + err.Error())
	}
	return &block, nil
}

// Hash calculates and returns the hash of the block.
func (b *Block) Hash() ([]byte, error) {
	serializedBlock, err := b.Serialize()
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256(serializedBlock)
	return hash[:], nil
}

// Validate performs basic validation on the block.
func (b *Block) Validate() error {
	if b.Header == nil {
		return errors.New("block header is nil")
	}
	if b.Body == nil {
		return errors.New("block body is nil")
	}
	if b.Header.Timestamp > time.Now().UnixNano() {
		return errors.New("block timestamp is in the future")
	}
	// Add more validation logic here (e.g., signature verification, VRF proof validation)
	return nil
}
