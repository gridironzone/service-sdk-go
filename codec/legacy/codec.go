package legacy

import (
	"github.com/gridironzone/service-sdk-go/codec"
	cryptocodec "github.com/gridironzone/service-sdk-go/crypto/codec"
)

// Cdc defines a global generic sealed Amino codec to be used throughout sdk. It
// has all Tendermint crypto and evidence types registered.
//
// TODO: Deprecated - remove this global.
var Cdc *codec.LegacyAmino

func init() {
	Cdc = codec.NewLegacyAmino()
	cryptocodec.RegisterCrypto(Cdc)
	codec.RegisterEvidences(Cdc)
	Cdc.Seal()
}
