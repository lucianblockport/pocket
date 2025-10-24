package blockchain

import (
	"crypto/sha256"
	"errors"
	"sort"
)

// AccountState defines the state of an account.
type AccountState struct {
	Balance     uint64
	SequenceNum uint64
	Stake       uint64
	Reputation  int64
}

// State represents the state of the blockchain.
type State struct {
	Accounts map[string]*AccountState
	// UTXO-style state with Merkle tree for light client proofs
}

// NewState creates a new state.
func NewState() *State {
	return &State{
		Accounts: make(map[string]*AccountState),
	}
}

// ApplyTransaction applies a transaction to the state.
func (s *State) ApplyTransaction(tx *Transaction) error {
	if err := tx.Verify(); err != nil {
		return err
	}

	fromAddr := string(tx.From)
	toAddr := string(tx.To)

	fromAccount, ok := s.Accounts[fromAddr]
	if !ok {
		return errors.New("from account not found")
	}

	if fromAccount.SequenceNum+1 != tx.SequenceNum {
		return errors.New("invalid sequence number")
	}

	if fromAccount.Balance < tx.Amount {
		return errors.New("insufficient balance")
	}

	fromAccount.Balance -= tx.Amount
	fromAccount.SequenceNum = tx.SequenceNum

	toAccount, ok := s.Accounts[toAddr]
	if !ok {
		s.Accounts[toAddr] = &AccountState{Balance: 0, SequenceNum: 0}
		toAccount = s.Accounts[toAddr]
	}

	toAccount.Balance += tx.Amount

	return nil
}

// GetAccount returns the state of an account.
func (s *State) GetAccount(address string) (*AccountState, error) {
	account, ok := s.Accounts[address]
	if !ok {
		return nil, errors.New("account not found")
	}
	return account, nil
}

// UpdateStake updates the stake of an account.
func (s *State) UpdateStake(address string, amount int64) error {
	account, ok := s.Accounts[address]
	if !ok {
		return errors.New("account not found")
	}
	if amount < 0 && uint64(-amount) > account.Stake {
		return errors.New("insufficient stake")
	}
	account.Stake = uint64(int64(account.Stake) + amount)
	return nil
}

// ComputeStateRoot computes the Merkle root of the state.
func (s *State) ComputeStateRoot() ([]byte, error) {
	var addresses []string
	for addr := range s.Accounts {
		addresses = append(addresses, addr)
	}
	sort.Strings(addresses)

	var hashes [][]byte
	for _, addr := range addresses {
		account := s.Accounts[addr]
		hash := sha256.Sum256([]byte(addr))
		hashes = append(hashes, hash[:])
		hash = sha256.Sum256([]byte(string(account.Balance)))
		hashes = append(hashes, hash[:])
		hash = sha256.Sum256([]byte(string(account.SequenceNum)))
		hashes = append(hashes, hash[:])
	}

	if len(hashes) == 0 {
		return nil, nil
	}

	for len(hashes) > 1 {
		var nextLevel [][]byte
		for i := 0; i < len(hashes); i += 2 {
			if i+1 == len(hashes) {
				nextLevel = append(nextLevel, hashes[i])
			} else {
				combinedHash := append(hashes[i], hashes[i+1]...)
				hash := sha256.Sum256(combinedHash)
				nextLevel = append(nextLevel, hash[:])
			}
		}
		hashes = nextLevel
	}

	return hashes[0], nil
}
