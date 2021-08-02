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

// ChannelsGetLeftChannelsRequest represents TL type `channels.getLeftChannels#8341ecc0`.
// Get a list of channels/supergroups¹ we left
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.getLeftChannels for reference.
type ChannelsGetLeftChannelsRequest struct {
	// Offset for pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Offset int
}

// ChannelsGetLeftChannelsRequestTypeID is TL type id of ChannelsGetLeftChannelsRequest.
const ChannelsGetLeftChannelsRequestTypeID = 0x8341ecc0

func (g *ChannelsGetLeftChannelsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Offset == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ChannelsGetLeftChannelsRequest) String() string {
	if g == nil {
		return "ChannelsGetLeftChannelsRequest(nil)"
	}
	type Alias ChannelsGetLeftChannelsRequest
	return fmt.Sprintf("ChannelsGetLeftChannelsRequest%+v", Alias(*g))
}

// FillFrom fills ChannelsGetLeftChannelsRequest from given interface.
func (g *ChannelsGetLeftChannelsRequest) FillFrom(from interface {
	GetOffset() (value int)
}) {
	g.Offset = from.GetOffset()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsGetLeftChannelsRequest) TypeID() uint32 {
	return ChannelsGetLeftChannelsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsGetLeftChannelsRequest) TypeName() string {
	return "channels.getLeftChannels"
}

// TypeInfo returns info about TL type.
func (g *ChannelsGetLeftChannelsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.getLeftChannels",
		ID:   ChannelsGetLeftChannelsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *ChannelsGetLeftChannelsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode %s as nil", "channels.getLeftChannels#8341ecc0")
	}
	b.PutID(ChannelsGetLeftChannelsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *ChannelsGetLeftChannelsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode %s as nil", "channels.getLeftChannels#8341ecc0")
	}
	b.PutInt(g.Offset)
	return nil
}

// GetOffset returns value of Offset field.
func (g *ChannelsGetLeftChannelsRequest) GetOffset() (value int) {
	return g.Offset
}

// Decode implements bin.Decoder.
func (g *ChannelsGetLeftChannelsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode %s to nil", "channels.getLeftChannels#8341ecc0")
	}
	if err := b.ConsumeID(ChannelsGetLeftChannelsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "channels.getLeftChannels#8341ecc0", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *ChannelsGetLeftChannelsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode %s to nil", "channels.getLeftChannels#8341ecc0")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "channels.getLeftChannels#8341ecc0", "offset", err)
		}
		g.Offset = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsGetLeftChannelsRequest.
var (
	_ bin.Encoder     = &ChannelsGetLeftChannelsRequest{}
	_ bin.Decoder     = &ChannelsGetLeftChannelsRequest{}
	_ bin.BareEncoder = &ChannelsGetLeftChannelsRequest{}
	_ bin.BareDecoder = &ChannelsGetLeftChannelsRequest{}
)

// ChannelsGetLeftChannels invokes method channels.getLeftChannels#8341ecc0 returning error if any.
// Get a list of channels/supergroups¹ we left
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  403 TAKEOUT_REQUIRED: A takeout session has to be initialized, first
//
// See https://core.telegram.org/method/channels.getLeftChannels for reference.
func (c *Client) ChannelsGetLeftChannels(ctx context.Context, offset int) (MessagesChatsClass, error) {
	var result MessagesChatsBox

	request := &ChannelsGetLeftChannelsRequest{
		Offset: offset,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Chats, nil
}
