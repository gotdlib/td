// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// PingRequest represents TL type `ping#7abe77ec`.
type PingRequest struct {
	// PingID field of PingRequest.
	PingID int64
}

// PingRequestTypeID is TL type id of PingRequest.
const PingRequestTypeID = 0x7abe77ec

func (p *PingRequest) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.PingID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PingRequest) String() string {
	if p == nil {
		return "PingRequest(nil)"
	}
	type Alias PingRequest
	return fmt.Sprintf("PingRequest%+v", Alias(*p))
}

// FillFrom fills PingRequest from given interface.
func (p *PingRequest) FillFrom(from interface {
	GetPingID() (value int64)
}) {
	p.PingID = from.GetPingID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PingRequest) TypeID() uint32 {
	return PingRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PingRequest) TypeName() string {
	return "ping"
}

// TypeInfo returns info about TL type.
func (p *PingRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "ping",
		ID:   PingRequestTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PingID",
			SchemaName: "ping_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PingRequest) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode %s as nil", "ping#7abe77ec")
	}
	b.PutID(PingRequestTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PingRequest) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode %s as nil", "ping#7abe77ec")
	}
	b.PutLong(p.PingID)
	return nil
}

// GetPingID returns value of PingID field.
func (p *PingRequest) GetPingID() (value int64) {
	return p.PingID
}

// Decode implements bin.Decoder.
func (p *PingRequest) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode %s to nil", "ping#7abe77ec")
	}
	if err := b.ConsumeID(PingRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "ping#7abe77ec", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PingRequest) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode %s to nil", "ping#7abe77ec")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "ping#7abe77ec", "ping_id", err)
		}
		p.PingID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PingRequest.
var (
	_ bin.Encoder     = &PingRequest{}
	_ bin.Decoder     = &PingRequest{}
	_ bin.BareEncoder = &PingRequest{}
	_ bin.BareDecoder = &PingRequest{}
)

// Ping invokes method ping#7abe77ec returning error if any.
func (c *Client) Ping(ctx context.Context, pingid int64) (*Pong, error) {
	var result Pong

	request := &PingRequest{
		PingID: pingid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
