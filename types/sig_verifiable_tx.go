package types

import (
	cryptotypes "github.com/gridironzone/service-sdk-go/crypto/types"
	"github.com/gridironzone/service-sdk-go/types/tx/signing"
)

// SigVerifiableTx defines a transaction interface for all signature verification
// handlers.
type SigVerifiableTx interface {
	Tx
	GetSigners() []AccAddress
	GetPubKeys() []cryptotypes.PubKey // If signer already has pubkey in context, this list will have nil in its place
	GetSignaturesV2() ([]signing.SignatureV2, error)
}

// Tx defines a transaction interface that supports all standard message, signature
// fee, memo, and auxiliary interfaces.
type SigTx interface {
	SigVerifiableTx

	TxWithMemo
	FeeTx
	TxWithTimeoutHeight
}
