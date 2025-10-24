package consensus

import (
	"errors"
	"sync"
	"time"
)

// PBFT represents a PBFT consensus instance.
type PBFT struct {
	mu              sync.Mutex
	validators      []string
	leader          string
	view            uint64
	prepares        map[string]map[string][]byte // block hash -> validator -> signature
	commits         map[string]map[string][]byte // block hash -> validator -> signature
	viewChanges     map[uint64]map[string]bool  // view number -> validator -> seen
	currentBlock    []byte
	consensusTimer  *time.Timer
	leaderTimeout   time.Duration
	consensusStatus chan bool
}

// NewPBFT creates a new PBFT instance.
func NewPBFT(validators []string, leaderTimeout time.Duration) *PBFT {
	return &PBFT{
		validators:      validators,
		leader:          validators[0], // Simple leader selection
		view:            0,
		prepares:        make(map[string]map[string][]byte),
		commits:         make(map[string]map[string][]byte),
		viewChanges:     make(map[uint64]map[string]bool),
		leaderTimeout:   leaderTimeout,
		consensusStatus: make(chan bool),
	}
}

// HandlePrepare handles a prepare message.
func (p *PBFT) HandlePrepare(blockHash, validator string, signature []byte) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.prepares[blockHash]; !ok {
		p.prepares[blockHash] = make(map[string][]byte)
	}
	p.prepares[blockHash][validator] = signature

	if len(p.prepares[blockHash]) >= (2*len(p.validators))/3+1 {
		// Broadcast commit
	}

	return nil
}

// HandleCommit handles a commit message.
func (p *PBFT) HandleCommit(blockHash, validator string, signature []byte) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.commits[blockHash]; !ok {
		p.commits[blockHash] = make(map[string][]byte)
	}
	p.commits[blockHash][validator] = signature

	if len(p.commits[blockHash]) >= (2*len(p.validators))/3+1 {
		p.consensusStatus <- true
	}

	return nil
}

// HandleViewChange handles a view change message.
func (p *PBFT) HandleViewChange(view uint64, validator string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.viewChanges[view]; !ok {
		p.viewChanges[view] = make(map[string]bool)
	}
	p.viewChanges[view][validator] = true

	if len(p.viewChanges[view]) >= (2*len(p.validators))/3+1 {
		p.view++
		p.leader = p.validators[p.view%uint64(len(p.validators))] // Simple round-robin leader change
		// Start new view
	}

	return nil
}

// StartConsensus starts the consensus process.
func (p *PBFT) StartConsensus(block []byte) (bool, error) {
	p.mu.Lock()
	p.currentBlock = block
	// Broadcast pre-prepare
	p.mu.Unlock()

	p.consensusTimer = time.NewTimer(p.leaderTimeout)

	select {
	case <-p.consensusStatus:
		return true, nil
	case <-p.consensusTimer.C:
		// Broadcast view change
		return false, errors.New("consensus timeout")
	}
}
