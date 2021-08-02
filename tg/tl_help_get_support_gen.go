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

// HelpGetSupportRequest represents TL type `help.getSupport#9cdf08cd`.
// Returns the support user for the 'ask a question' feature.
//
// See https://core.telegram.org/method/help.getSupport for reference.
type HelpGetSupportRequest struct {
}

// HelpGetSupportRequestTypeID is TL type id of HelpGetSupportRequest.
const HelpGetSupportRequestTypeID = 0x9cdf08cd

func (g *HelpGetSupportRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetSupportRequest) String() string {
	if g == nil {
		return "HelpGetSupportRequest(nil)"
	}
	type Alias HelpGetSupportRequest
	return fmt.Sprintf("HelpGetSupportRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpGetSupportRequest) TypeID() uint32 {
	return HelpGetSupportRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpGetSupportRequest) TypeName() string {
	return "help.getSupport"
}

// TypeInfo returns info about TL type.
func (g *HelpGetSupportRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.getSupport",
		ID:   HelpGetSupportRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *HelpGetSupportRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode %s as nil", "help.getSupport#9cdf08cd")
	}
	b.PutID(HelpGetSupportRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *HelpGetSupportRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode %s as nil", "help.getSupport#9cdf08cd")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetSupportRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode %s to nil", "help.getSupport#9cdf08cd")
	}
	if err := b.ConsumeID(HelpGetSupportRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "help.getSupport#9cdf08cd", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *HelpGetSupportRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode %s to nil", "help.getSupport#9cdf08cd")
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetSupportRequest.
var (
	_ bin.Encoder     = &HelpGetSupportRequest{}
	_ bin.Decoder     = &HelpGetSupportRequest{}
	_ bin.BareEncoder = &HelpGetSupportRequest{}
	_ bin.BareDecoder = &HelpGetSupportRequest{}
)

// HelpGetSupport invokes method help.getSupport#9cdf08cd returning error if any.
// Returns the support user for the 'ask a question' feature.
//
// See https://core.telegram.org/method/help.getSupport for reference.
func (c *Client) HelpGetSupport(ctx context.Context) (*HelpSupport, error) {
	var result HelpSupport

	request := &HelpGetSupportRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
