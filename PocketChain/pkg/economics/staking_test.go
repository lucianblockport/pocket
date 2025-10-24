package economics

import (
	"math/big"
	"testing"
)

func TestStaking(t *testing.T) {
	s := NewStaking()

	nodeID := "test-node"
	stakeAmount := big.NewInt(100)

	// Test Stake
	err := s.Stake(nodeID, stakeAmount)
	if err != nil {
		t.Fatalf("Stake() error = %v", err)
	}

	// Test GetStake
	stake, err := s.GetStake(nodeID)
	if err != nil {
		t.Fatalf("GetStake() error = %v", err)
	}
	if stake.Cmp(stakeAmount) != 0 {
		t.Errorf("GetStake() got = %v, want %v", stake, stakeAmount)
	}

	// Test Unstake
	unstakeAmount := big.NewInt(50)
	err = s.Unstake(nodeID, unstakeAmount)
	if err != nil {
		t.Fatalf("Unstake() error = %v", err)
	}

	// Test GetStake after Unstake
	stake, err = s.GetStake(nodeID)
	if err != nil {
		t.Fatalf("GetStake() error = %v", err)
	}
	expectedStake := big.NewInt(50)
	if stake.Cmp(expectedStake) != 0 {
		t.Errorf("GetStake() got = %v, want %v", stake, expectedStake)
	}

	// Test Unstake with insufficient stake
	err = s.Unstake(nodeID, big.NewInt(100))
	if err == nil {
		t.Errorf("Unstake() expected error for insufficient stake, but got nil")
	}
}
