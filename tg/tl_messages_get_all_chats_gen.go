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

// MessagesGetAllChatsRequest represents TL type `messages.getAllChats#eba80ff0`.
// Get all chats, channels and supergroups
//
// See https://core.telegram.org/method/messages.getAllChats for reference.
type MessagesGetAllChatsRequest struct {
	// Except these chats/channels/supergroups
	ExceptIDs []int
}

// MessagesGetAllChatsRequestTypeID is TL type id of MessagesGetAllChatsRequest.
const MessagesGetAllChatsRequestTypeID = 0xeba80ff0

func (g *MessagesGetAllChatsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ExceptIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetAllChatsRequest) String() string {
	if g == nil {
		return "MessagesGetAllChatsRequest(nil)"
	}
	type Alias MessagesGetAllChatsRequest
	return fmt.Sprintf("MessagesGetAllChatsRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetAllChatsRequest from given interface.
func (g *MessagesGetAllChatsRequest) FillFrom(from interface {
	GetExceptIDs() (value []int)
}) {
	g.ExceptIDs = from.GetExceptIDs()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetAllChatsRequest) TypeID() uint32 {
	return MessagesGetAllChatsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetAllChatsRequest) TypeName() string {
	return "messages.getAllChats"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetAllChatsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getAllChats",
		ID:   MessagesGetAllChatsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ExceptIDs",
			SchemaName: "except_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetAllChatsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode %s as nil", "messages.getAllChats#eba80ff0")
	}
	b.PutID(MessagesGetAllChatsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetAllChatsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode %s as nil", "messages.getAllChats#eba80ff0")
	}
	b.PutVectorHeader(len(g.ExceptIDs))
	for _, v := range g.ExceptIDs {
		b.PutInt(v)
	}
	return nil
}

// GetExceptIDs returns value of ExceptIDs field.
func (g *MessagesGetAllChatsRequest) GetExceptIDs() (value []int) {
	return g.ExceptIDs
}

// Decode implements bin.Decoder.
func (g *MessagesGetAllChatsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode %s to nil", "messages.getAllChats#eba80ff0")
	}
	if err := b.ConsumeID(MessagesGetAllChatsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "messages.getAllChats#eba80ff0", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetAllChatsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode %s to nil", "messages.getAllChats#eba80ff0")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.getAllChats#eba80ff0", "except_ids", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode %s: field %s: %w", "messages.getAllChats#eba80ff0", "except_ids", err)
			}
			g.ExceptIDs = append(g.ExceptIDs, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetAllChatsRequest.
var (
	_ bin.Encoder     = &MessagesGetAllChatsRequest{}
	_ bin.Decoder     = &MessagesGetAllChatsRequest{}
	_ bin.BareEncoder = &MessagesGetAllChatsRequest{}
	_ bin.BareDecoder = &MessagesGetAllChatsRequest{}
)

// MessagesGetAllChats invokes method messages.getAllChats#eba80ff0 returning error if any.
// Get all chats, channels and supergroups
//
// See https://core.telegram.org/method/messages.getAllChats for reference.
func (c *Client) MessagesGetAllChats(ctx context.Context, exceptids []int) (MessagesChatsClass, error) {
	var result MessagesChatsBox

	request := &MessagesGetAllChatsRequest{
		ExceptIDs: exceptids,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Chats, nil
}
