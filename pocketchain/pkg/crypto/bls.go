package crypto

import (
	"errors"

	"github.com/herumi/bls-go-binary/bls"
)

func init() {
	bls.Init(bls.BLS12_381)
}

// Sign generates a BLS signature.
func Sign(secretKey *bls.SecretKey, message []byte) *bls.Sign {
	return secretKey.Sign(string(message))
}

// Verify verifies a BLS signature.
func Verify(publicKey *bls.PublicKey, signature *bls.Sign, message []byte) bool {
	return signature.Verify(publicKey, string(message))
}

// AggregateSignatures aggregates multiple BLS signatures into one.
func AggregateSignatures(signatures []*bls.Sign) (*bls.Sign, error) {
	if len(signatures) == 0 {
		return nil, errors.New("no signatures to aggregate")
	}
	aggregatedSig := new(bls.Sign)
	aggregatedSig.Aggregate(signatures)
	return aggregatedSig, nil
}

// VerifyAggregate verifies an aggregated BLS signature.
func VerifyAggregate(publicKeys []*bls.PublicKey, aggregatedSig *bls.Sign, message []byte) bool {
	return aggregatedSig.FastAggregateVerify(publicKeys, message)
}
