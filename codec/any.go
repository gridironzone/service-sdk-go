package codec

import (
	"fmt"

	"github.com/gogo/protobuf/proto"

	"github.com/gridironzone/service-sdk-go/codec/types"
)

// MarshalAny is a convenience function for packing the provided value in an
// Any and then proto marshaling it to bytes
func MarshalAny(m BinaryMarshaler, x interface{}) ([]byte, error) {
	msg, ok := x.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("can't proto marshal %T", x)
	}

	any := &types.Any{}
	if err := any.Pack(msg); err != nil {
		return nil, err
	}

	return m.MarshalBinaryBare(any)
}

// UnmarshalAny is a convenience function for proto unmarshaling an Any from
// bz and then unpacking it to the interface pointer passed in as iface using
// the provided AnyUnpacker or returning an error
//
// Ex:
//		var x MyInterface
//		err := UnmarshalAny(unpacker, &x, bz)
func UnmarshalAny(m BinaryMarshaler, iface interface{}, bz []byte) error {
	any := &types.Any{}

	if err := m.UnmarshalBinaryBare(bz, any); err != nil {
		return err
	}

	return m.UnpackAny(any, iface)
}
