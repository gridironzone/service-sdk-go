package token

import (
	"github.com/gridironzone/service-sdk-go/codec"
	"github.com/gridironzone/service-sdk-go/codec/types"
	cryptocodec "github.com/gridironzone/service-sdk-go/crypto/codec"
	sdk "github.com/gridironzone/service-sdk-go/types"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgIssueToken{},
		&MsgEditToken{},
		&MsgMintToken{},
		&MsgTransferTokenOwner{},
	)
	registry.RegisterInterface("irismod.token.TokenI", (*TokenInterface)(nil), &Token{})
}
