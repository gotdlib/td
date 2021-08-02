// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// SecureValueHash represents TL type `secureValueHash#ed1ecdb0`.
// Secure value hash
//
// See https://core.telegram.org/constructor/secureValueHash for reference.
type SecureValueHash struct {
	// Secure value type
	Type SecureValueTypeClass
	// Hash
	Hash []byte
}

// SecureValueHashTypeID is TL type id of SecureValueHash.
const SecureValueHashTypeID = 0xed1ecdb0

func (s *SecureValueHash) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Type == nil) {
		return false
	}
	if !(s.Hash == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureValueHash) String() string {
	if s == nil {
		return "SecureValueHash(nil)"
	}
	type Alias SecureValueHash
	return fmt.Sprintf("SecureValueHash%+v", Alias(*s))
}

// FillFrom fills SecureValueHash from given interface.
func (s *SecureValueHash) FillFrom(from interface {
	GetType() (value SecureValueTypeClass)
	GetHash() (value []byte)
}) {
	s.Type = from.GetType()
	s.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SecureValueHash) TypeID() uint32 {
	return SecureValueHashTypeID
}

// TypeName returns name of type in TL schema.
func (*SecureValueHash) TypeName() string {
	return "secureValueHash"
}

// TypeInfo returns info about TL type.
func (s *SecureValueHash) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "secureValueHash",
		ID:   SecureValueHashTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SecureValueHash) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode %s as nil", "secureValueHash#ed1ecdb0")
	}
	b.PutID(SecureValueHashTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SecureValueHash) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode %s as nil", "secureValueHash#ed1ecdb0")
	}
	if s.Type == nil {
		return fmt.Errorf("unable to encode %s: field %s is nil", "secureValueHash#ed1ecdb0", "type")
	}
	if err := s.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "secureValueHash#ed1ecdb0", "type", err)
	}
	b.PutBytes(s.Hash)
	return nil
}

// GetType returns value of Type field.
func (s *SecureValueHash) GetType() (value SecureValueTypeClass) {
	return s.Type
}

// GetHash returns value of Hash field.
func (s *SecureValueHash) GetHash() (value []byte) {
	return s.Hash
}

// Decode implements bin.Decoder.
func (s *SecureValueHash) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode %s to nil", "secureValueHash#ed1ecdb0")
	}
	if err := b.ConsumeID(SecureValueHashTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "secureValueHash#ed1ecdb0", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SecureValueHash) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode %s to nil", "secureValueHash#ed1ecdb0")
	}
	{
		value, err := DecodeSecureValueType(b)
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "secureValueHash#ed1ecdb0", "type", err)
		}
		s.Type = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "secureValueHash#ed1ecdb0", "hash", err)
		}
		s.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for SecureValueHash.
var (
	_ bin.Encoder     = &SecureValueHash{}
	_ bin.Decoder     = &SecureValueHash{}
	_ bin.BareEncoder = &SecureValueHash{}
	_ bin.BareDecoder = &SecureValueHash{}
)
