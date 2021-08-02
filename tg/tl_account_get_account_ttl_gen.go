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

// AccountGetAccountTTLRequest represents TL type `account.getAccountTTL#8fc711d`.
// Get days to live of account
//
// See https://core.telegram.org/method/account.getAccountTTL for reference.
type AccountGetAccountTTLRequest struct {
}

// AccountGetAccountTTLRequestTypeID is TL type id of AccountGetAccountTTLRequest.
const AccountGetAccountTTLRequestTypeID = 0x8fc711d

func (g *AccountGetAccountTTLRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetAccountTTLRequest) String() string {
	if g == nil {
		return "AccountGetAccountTTLRequest(nil)"
	}
	type Alias AccountGetAccountTTLRequest
	return fmt.Sprintf("AccountGetAccountTTLRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetAccountTTLRequest) TypeID() uint32 {
	return AccountGetAccountTTLRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetAccountTTLRequest) TypeName() string {
	return "account.getAccountTTL"
}

// TypeInfo returns info about TL type.
func (g *AccountGetAccountTTLRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getAccountTTL",
		ID:   AccountGetAccountTTLRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetAccountTTLRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode %s as nil", "account.getAccountTTL#8fc711d")
	}
	b.PutID(AccountGetAccountTTLRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetAccountTTLRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode %s as nil", "account.getAccountTTL#8fc711d")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetAccountTTLRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode %s to nil", "account.getAccountTTL#8fc711d")
	}
	if err := b.ConsumeID(AccountGetAccountTTLRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "account.getAccountTTL#8fc711d", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetAccountTTLRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode %s to nil", "account.getAccountTTL#8fc711d")
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetAccountTTLRequest.
var (
	_ bin.Encoder     = &AccountGetAccountTTLRequest{}
	_ bin.Decoder     = &AccountGetAccountTTLRequest{}
	_ bin.BareEncoder = &AccountGetAccountTTLRequest{}
	_ bin.BareDecoder = &AccountGetAccountTTLRequest{}
)

// AccountGetAccountTTL invokes method account.getAccountTTL#8fc711d returning error if any.
// Get days to live of account
//
// See https://core.telegram.org/method/account.getAccountTTL for reference.
func (c *Client) AccountGetAccountTTL(ctx context.Context) (*AccountDaysTTL, error) {
	var result AccountDaysTTL

	request := &AccountGetAccountTTLRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
