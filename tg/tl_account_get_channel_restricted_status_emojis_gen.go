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
	"github.com/gotd/td/tdjson"
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
	_ = tdjson.Encoder{}
)

// AccountGetChannelRestrictedStatusEmojisRequest represents TL type `account.getChannelRestrictedStatusEmojis#35a9e0d5`.
//
// See https://core.telegram.org/method/account.getChannelRestrictedStatusEmojis for reference.
type AccountGetChannelRestrictedStatusEmojisRequest struct {
	// Hash field of AccountGetChannelRestrictedStatusEmojisRequest.
	Hash int64
}

// AccountGetChannelRestrictedStatusEmojisRequestTypeID is TL type id of AccountGetChannelRestrictedStatusEmojisRequest.
const AccountGetChannelRestrictedStatusEmojisRequestTypeID = 0x35a9e0d5

// Ensuring interfaces in compile-time for AccountGetChannelRestrictedStatusEmojisRequest.
var (
	_ bin.Encoder     = &AccountGetChannelRestrictedStatusEmojisRequest{}
	_ bin.Decoder     = &AccountGetChannelRestrictedStatusEmojisRequest{}
	_ bin.BareEncoder = &AccountGetChannelRestrictedStatusEmojisRequest{}
	_ bin.BareDecoder = &AccountGetChannelRestrictedStatusEmojisRequest{}
)

func (g *AccountGetChannelRestrictedStatusEmojisRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetChannelRestrictedStatusEmojisRequest) String() string {
	if g == nil {
		return "AccountGetChannelRestrictedStatusEmojisRequest(nil)"
	}
	type Alias AccountGetChannelRestrictedStatusEmojisRequest
	return fmt.Sprintf("AccountGetChannelRestrictedStatusEmojisRequest%+v", Alias(*g))
}

// FillFrom fills AccountGetChannelRestrictedStatusEmojisRequest from given interface.
func (g *AccountGetChannelRestrictedStatusEmojisRequest) FillFrom(from interface {
	GetHash() (value int64)
}) {
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetChannelRestrictedStatusEmojisRequest) TypeID() uint32 {
	return AccountGetChannelRestrictedStatusEmojisRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetChannelRestrictedStatusEmojisRequest) TypeName() string {
	return "account.getChannelRestrictedStatusEmojis"
}

// TypeInfo returns info about TL type.
func (g *AccountGetChannelRestrictedStatusEmojisRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getChannelRestrictedStatusEmojis",
		ID:   AccountGetChannelRestrictedStatusEmojisRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetChannelRestrictedStatusEmojisRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getChannelRestrictedStatusEmojis#35a9e0d5 as nil")
	}
	b.PutID(AccountGetChannelRestrictedStatusEmojisRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetChannelRestrictedStatusEmojisRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getChannelRestrictedStatusEmojis#35a9e0d5 as nil")
	}
	b.PutLong(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetChannelRestrictedStatusEmojisRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getChannelRestrictedStatusEmojis#35a9e0d5 to nil")
	}
	if err := b.ConsumeID(AccountGetChannelRestrictedStatusEmojisRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getChannelRestrictedStatusEmojis#35a9e0d5: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetChannelRestrictedStatusEmojisRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getChannelRestrictedStatusEmojis#35a9e0d5 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode account.getChannelRestrictedStatusEmojis#35a9e0d5: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetHash returns value of Hash field.
func (g *AccountGetChannelRestrictedStatusEmojisRequest) GetHash() (value int64) {
	if g == nil {
		return
	}
	return g.Hash
}

// AccountGetChannelRestrictedStatusEmojis invokes method account.getChannelRestrictedStatusEmojis#35a9e0d5 returning error if any.
//
// See https://core.telegram.org/method/account.getChannelRestrictedStatusEmojis for reference.
func (c *Client) AccountGetChannelRestrictedStatusEmojis(ctx context.Context, hash int64) (EmojiListClass, error) {
	var result EmojiListBox

	request := &AccountGetChannelRestrictedStatusEmojisRequest{
		Hash: hash,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.EmojiList, nil
}
