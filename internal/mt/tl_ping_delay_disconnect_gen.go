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

// PingDelayDisconnectRequest represents TL type `ping_delay_disconnect#f3427b8c`.
type PingDelayDisconnectRequest struct {
	// PingID field of PingDelayDisconnectRequest.
	PingID int64
	// DisconnectDelay field of PingDelayDisconnectRequest.
	DisconnectDelay int
}

// PingDelayDisconnectRequestTypeID is TL type id of PingDelayDisconnectRequest.
const PingDelayDisconnectRequestTypeID = 0xf3427b8c

func (p *PingDelayDisconnectRequest) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.PingID == 0) {
		return false
	}
	if !(p.DisconnectDelay == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PingDelayDisconnectRequest) String() string {
	if p == nil {
		return "PingDelayDisconnectRequest(nil)"
	}
	type Alias PingDelayDisconnectRequest
	return fmt.Sprintf("PingDelayDisconnectRequest%+v", Alias(*p))
}

// FillFrom fills PingDelayDisconnectRequest from given interface.
func (p *PingDelayDisconnectRequest) FillFrom(from interface {
	GetPingID() (value int64)
	GetDisconnectDelay() (value int)
}) {
	p.PingID = from.GetPingID()
	p.DisconnectDelay = from.GetDisconnectDelay()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PingDelayDisconnectRequest) TypeID() uint32 {
	return PingDelayDisconnectRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PingDelayDisconnectRequest) TypeName() string {
	return "ping_delay_disconnect"
}

// TypeInfo returns info about TL type.
func (p *PingDelayDisconnectRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "ping_delay_disconnect",
		ID:   PingDelayDisconnectRequestTypeID,
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
		{
			Name:       "DisconnectDelay",
			SchemaName: "disconnect_delay",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PingDelayDisconnectRequest) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode %s as nil", "ping_delay_disconnect#f3427b8c")
	}
	b.PutID(PingDelayDisconnectRequestTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PingDelayDisconnectRequest) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode %s as nil", "ping_delay_disconnect#f3427b8c")
	}
	b.PutLong(p.PingID)
	b.PutInt(p.DisconnectDelay)
	return nil
}

// GetPingID returns value of PingID field.
func (p *PingDelayDisconnectRequest) GetPingID() (value int64) {
	return p.PingID
}

// GetDisconnectDelay returns value of DisconnectDelay field.
func (p *PingDelayDisconnectRequest) GetDisconnectDelay() (value int) {
	return p.DisconnectDelay
}

// Decode implements bin.Decoder.
func (p *PingDelayDisconnectRequest) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode %s to nil", "ping_delay_disconnect#f3427b8c")
	}
	if err := b.ConsumeID(PingDelayDisconnectRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "ping_delay_disconnect#f3427b8c", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PingDelayDisconnectRequest) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode %s to nil", "ping_delay_disconnect#f3427b8c")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "ping_delay_disconnect#f3427b8c", "ping_id", err)
		}
		p.PingID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "ping_delay_disconnect#f3427b8c", "disconnect_delay", err)
		}
		p.DisconnectDelay = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PingDelayDisconnectRequest.
var (
	_ bin.Encoder     = &PingDelayDisconnectRequest{}
	_ bin.Decoder     = &PingDelayDisconnectRequest{}
	_ bin.BareEncoder = &PingDelayDisconnectRequest{}
	_ bin.BareDecoder = &PingDelayDisconnectRequest{}
)

// PingDelayDisconnect invokes method ping_delay_disconnect#f3427b8c returning error if any.
func (c *Client) PingDelayDisconnect(ctx context.Context, request *PingDelayDisconnectRequest) (*Pong, error) {
	var result Pong

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
