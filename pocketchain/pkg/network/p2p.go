package network

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

// P2PNode represents a libp2p node.
type P2PNode struct {
	host      host.Host
	gossipSub *pubsub.PubSub
}

// NewP2PNode creates a new P2PNode instance.
func NewP2PNode(ctx context.Context) (*P2PNode, error) {
	host, err := libp2p.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create libp2p host: %w", err)
	}

	gossipSub, err := pubsub.NewGossipSub(ctx, host)
	if err != nil {
		return nil, fmt.Errorf("failed to create gossipsub: %w", err)
	}

	return &P2PNode{
		host:      host,
		gossipSub: gossipSub,
	}, nil
}

// Start starts the P2P node.
func (n *P2PNode) Start() {
	// In a real implementation, you would start the DHT for peer discovery.
	fmt.Println("P2P node started:", n.host.ID())
}

// Connect connects to a peer.
func (n *P2PNode) Connect(ctx context.Context, peerInfo peer.AddrInfo) error {
	return n.host.Connect(ctx, peerInfo)
}

// Broadcast broadcasts a message to a topic.
func (n *P2PNode) Broadcast(ctx context.Context, topic string, data []byte) error {
	return n.gossipSub.Publish(topic, data)
}

// Subscribe subscribes to a topic.
func (n *P2PNode) Subscribe(topic string) (*pubsub.Subscription, error) {
	return n.gossipSub.Subscribe(topic)
}
