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

// CDNConfig represents TL type `cdnConfig#5725e40a`.
// Configuration for CDN¹ file downloads.
//
// Links:
//  1) https://core.telegram.org/cdn
//
// See https://core.telegram.org/constructor/cdnConfig for reference.
type CDNConfig struct {
	// Vector of public keys to use only during handshakes to CDN¹ DCs.
	//
	// Links:
	//  1) https://core.telegram.org/cdn
	PublicKeys []CDNPublicKey
}

// CDNConfigTypeID is TL type id of CDNConfig.
const CDNConfigTypeID = 0x5725e40a

func (c *CDNConfig) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.PublicKeys == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CDNConfig) String() string {
	if c == nil {
		return "CDNConfig(nil)"
	}
	type Alias CDNConfig
	return fmt.Sprintf("CDNConfig%+v", Alias(*c))
}

// FillFrom fills CDNConfig from given interface.
func (c *CDNConfig) FillFrom(from interface {
	GetPublicKeys() (value []CDNPublicKey)
}) {
	c.PublicKeys = from.GetPublicKeys()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CDNConfig) TypeID() uint32 {
	return CDNConfigTypeID
}

// TypeName returns name of type in TL schema.
func (*CDNConfig) TypeName() string {
	return "cdnConfig"
}

// TypeInfo returns info about TL type.
func (c *CDNConfig) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "cdnConfig",
		ID:   CDNConfigTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PublicKeys",
			SchemaName: "public_keys",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CDNConfig) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode %s as nil", "cdnConfig#5725e40a")
	}
	b.PutID(CDNConfigTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CDNConfig) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode %s as nil", "cdnConfig#5725e40a")
	}
	b.PutVectorHeader(len(c.PublicKeys))
	for idx, v := range c.PublicKeys {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode %s: field %s element with index %d: %w", "cdnConfig#5725e40a", "public_keys", idx, err)
		}
	}
	return nil
}

// GetPublicKeys returns value of PublicKeys field.
func (c *CDNConfig) GetPublicKeys() (value []CDNPublicKey) {
	return c.PublicKeys
}

// Decode implements bin.Decoder.
func (c *CDNConfig) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode %s to nil", "cdnConfig#5725e40a")
	}
	if err := b.ConsumeID(CDNConfigTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "cdnConfig#5725e40a", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CDNConfig) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode %s to nil", "cdnConfig#5725e40a")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "cdnConfig#5725e40a", "public_keys", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value CDNPublicKey
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode %s: field %s: %w", "cdnConfig#5725e40a", "public_keys", err)
			}
			c.PublicKeys = append(c.PublicKeys, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for CDNConfig.
var (
	_ bin.Encoder     = &CDNConfig{}
	_ bin.Decoder     = &CDNConfig{}
	_ bin.BareEncoder = &CDNConfig{}
	_ bin.BareDecoder = &CDNConfig{}
)
