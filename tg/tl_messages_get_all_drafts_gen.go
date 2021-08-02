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

// MessagesGetAllDraftsRequest represents TL type `messages.getAllDrafts#6a3f8d65`.
// Save get all message drafts¹.
//
// Links:
//  1) https://core.telegram.org/api/drafts
//
// See https://core.telegram.org/method/messages.getAllDrafts for reference.
type MessagesGetAllDraftsRequest struct {
}

// MessagesGetAllDraftsRequestTypeID is TL type id of MessagesGetAllDraftsRequest.
const MessagesGetAllDraftsRequestTypeID = 0x6a3f8d65

func (g *MessagesGetAllDraftsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetAllDraftsRequest) String() string {
	if g == nil {
		return "MessagesGetAllDraftsRequest(nil)"
	}
	type Alias MessagesGetAllDraftsRequest
	return fmt.Sprintf("MessagesGetAllDraftsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetAllDraftsRequest) TypeID() uint32 {
	return MessagesGetAllDraftsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetAllDraftsRequest) TypeName() string {
	return "messages.getAllDrafts"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetAllDraftsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getAllDrafts",
		ID:   MessagesGetAllDraftsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetAllDraftsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode %s as nil", "messages.getAllDrafts#6a3f8d65")
	}
	b.PutID(MessagesGetAllDraftsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetAllDraftsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode %s as nil", "messages.getAllDrafts#6a3f8d65")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetAllDraftsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode %s to nil", "messages.getAllDrafts#6a3f8d65")
	}
	if err := b.ConsumeID(MessagesGetAllDraftsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "messages.getAllDrafts#6a3f8d65", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetAllDraftsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode %s to nil", "messages.getAllDrafts#6a3f8d65")
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetAllDraftsRequest.
var (
	_ bin.Encoder     = &MessagesGetAllDraftsRequest{}
	_ bin.Decoder     = &MessagesGetAllDraftsRequest{}
	_ bin.BareEncoder = &MessagesGetAllDraftsRequest{}
	_ bin.BareDecoder = &MessagesGetAllDraftsRequest{}
)

// MessagesGetAllDrafts invokes method messages.getAllDrafts#6a3f8d65 returning error if any.
// Save get all message drafts¹.
//
// Links:
//  1) https://core.telegram.org/api/drafts
//
// See https://core.telegram.org/method/messages.getAllDrafts for reference.
func (c *Client) MessagesGetAllDrafts(ctx context.Context) (UpdatesClass, error) {
	var result UpdatesBox

	request := &MessagesGetAllDraftsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
