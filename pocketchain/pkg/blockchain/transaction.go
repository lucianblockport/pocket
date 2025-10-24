package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"errors"
)

// TransactionType defines the type of a transaction.
type TransactionType int

const (
	// StandardTransaction is a regular transaction.
	StandardTransaction TransactionType = iota
	// MeshTransaction is a transaction broadcasted over the mesh network.
	MeshTransaction
)

// Transaction defines the structure of a transaction in PocketChain.
type Transaction struct {
	Type           TransactionType
	From           []byte
	To             []byte
	Amount         uint64
	SequenceNum    uint64
	Signature      []byte
	MeshWitnesses  [][]byte
	TimelockExpire uint64
	RelayPath      [][]byte
}

// NewStandardTransaction creates a new standard transaction.
func NewStandardTransaction(from, to []byte, amount, sequenceNum uint64, signature []byte) *Transaction {
	return &Transaction{
		Type:        StandardTransaction,
		From:        from,
		To:          to,
		Amount:      amount,
		SequenceNum: sequenceNum,
		Signature:   signature,
	}
}

// NewMeshTransaction creates a new mesh transaction.
func NewMeshTransaction(from, to []byte, amount, sequenceNum, timelockExpire uint64, signature []byte, meshWitnesses, relayPath [][]byte) *Transaction {
	return &Transaction{
		Type:           MeshTransaction,
		From:           from,
		To:             to,
		Amount:         amount,
		SequenceNum:    sequenceNum,
		Signature:      signature,
		MeshWitnesses:  meshWitnesses,
		TimelockExpire: timelockExpire,
		RelayPath:      relayPath,
	}
}

// Serialize serializes the transaction into a byte slice.
func (t *Transaction) Serialize() ([]byte, error) {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(t)
	if err != nil {
		return nil, errors.New("failed to serialize transaction: " + err.Error())
	}
	return result.Bytes(), nil
}

// Hash calculates and returns the hash of the transaction.
func (t *Transaction) Hash() ([]byte, error) {
	serializedTx, err := t.Serialize()
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256(serializedTx)
	return hash[:], nil
}

// Verify performs basic validation on the transaction.
func (t *Transaction) Verify() error {
	if t.From == nil {
		return errors.New("transaction from is nil")
	}
	if t.To == nil {
		return errors.New("transaction to is nil")
	}
	if t.Amount == 0 {
		return errors.New("transaction amount is zero")
	}
	if t.Type == MeshTransaction {
		if len(t.MeshWitnesses) == 0 {
			return errors.New("mesh transaction has no witnesses")
		}
		if t.TimelockExpire == 0 {
			return errors.New("mesh transaction has no timelock expiration")
		}
	}
	// Add more validation logic here (e.g., signature verification, sequence number validation)
	return nil
}

// GetSize returns the size of the transaction in bytes.
func (t *Transaction) GetSize() (int, error) {
	serializedTx, err := t.Serialize()
	if err != nil {
		return 0, err
	}
	return len(serializedTx), nil
}
