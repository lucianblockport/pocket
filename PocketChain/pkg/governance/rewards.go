package governance

import (
	"math"

	"pocketchain/pkg/node"
)

const (
	baseReward          = 1.0
	timeOnNetMultiplier = 1.2
	locationMultiplier  = 1.5
)

// CalculateReward calculates the reward for a node.
func CalculateReward(node *node.FullNode) float64 {
	reward := baseReward

	// Apply time-on-net multiplier
	reward *= math.Pow(timeOnNetMultiplier, getTimeOnNet(node))

	// Apply location multiplier
	if isIncentivizedLocation(node) {
		reward *= locationMultiplier
	}

	return reward
}

// getTimeOnNet returns the time on net for a node.
func getTimeOnNet(node *node.FullNode) float64 {
	// In a real implementation, you would track the node's time on the network.
	return 1.0
}

// isIncentivizedLocation checks if a node is in an incentivized location.
func isIncentivizedLocation(node *node.FullNode) bool {
	// In a real implementation, you would use ZKP-based location proofs to verify the node's location.
	return false
}
