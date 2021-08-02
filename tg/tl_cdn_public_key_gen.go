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

// CDNPublicKey represents TL type `cdnPublicKey#c982eaba`.
// Public key to use only during handshakes to CDN¹ DCs.
//
// Links:
//  1) https://core.telegram.org/cdn
//
// See https://core.telegram.org/constructor/cdnPublicKey for reference.
type CDNPublicKey struct {
	// CDN DC¹ ID
	//
	// Links:
	//  1) https://core.telegram.org/cdn
	DCID int
	// RSA public key
	PublicKey string
}

// CDNPublicKeyTypeID is TL type id of CDNPublicKey.
const CDNPublicKeyTypeID = 0xc982eaba

func (c *CDNPublicKey) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.DCID == 0) {
		return false
	}
	if !(c.PublicKey == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CDNPublicKey) String() string {
	if c == nil {
		return "CDNPublicKey(nil)"
	}
	type Alias CDNPublicKey
	return fmt.Sprintf("CDNPublicKey%+v", Alias(*c))
}

// FillFrom fills CDNPublicKey from given interface.
func (c *CDNPublicKey) FillFrom(from interface {
	GetDCID() (value int)
	GetPublicKey() (value string)
}) {
	c.DCID = from.GetDCID()
	c.PublicKey = from.GetPublicKey()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CDNPublicKey) TypeID() uint32 {
	return CDNPublicKeyTypeID
}

// TypeName returns name of type in TL schema.
func (*CDNPublicKey) TypeName() string {
	return "cdnPublicKey"
}

// TypeInfo returns info about TL type.
func (c *CDNPublicKey) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "cdnPublicKey",
		ID:   CDNPublicKeyTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "DCID",
			SchemaName: "dc_id",
		},
		{
			Name:       "PublicKey",
			SchemaName: "public_key",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CDNPublicKey) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode %s as nil", "cdnPublicKey#c982eaba")
	}
	b.PutID(CDNPublicKeyTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CDNPublicKey) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode %s as nil", "cdnPublicKey#c982eaba")
	}
	b.PutInt(c.DCID)
	b.PutString(c.PublicKey)
	return nil
}

// GetDCID returns value of DCID field.
func (c *CDNPublicKey) GetDCID() (value int) {
	return c.DCID
}

// GetPublicKey returns value of PublicKey field.
func (c *CDNPublicKey) GetPublicKey() (value string) {
	return c.PublicKey
}

// Decode implements bin.Decoder.
func (c *CDNPublicKey) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode %s to nil", "cdnPublicKey#c982eaba")
	}
	if err := b.ConsumeID(CDNPublicKeyTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "cdnPublicKey#c982eaba", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CDNPublicKey) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode %s to nil", "cdnPublicKey#c982eaba")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "cdnPublicKey#c982eaba", "dc_id", err)
		}
		c.DCID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "cdnPublicKey#c982eaba", "public_key", err)
		}
		c.PublicKey = value
	}
	return nil
}

// Ensuring interfaces in compile-time for CDNPublicKey.
var (
	_ bin.Encoder     = &CDNPublicKey{}
	_ bin.Decoder     = &CDNPublicKey{}
	_ bin.BareEncoder = &CDNPublicKey{}
	_ bin.BareDecoder = &CDNPublicKey{}
)
