package economics

const (
	// TotalSupply is the total supply of the native token.
	TotalSupply = 100_000_000_000

	// InitialCirculatingSupply is the initial circulating supply.
	InitialCirculatingSupply = 10_000_000_000
)

// Token represents the native token of the PocketChain network.
type Token struct {
	// In a real implementation, you would have a balance mapping.
}

// NewToken creates a new Token instance.
func NewToken() *Token {
	return &Token{}
}
