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

// MessagesSetHistoryTTLRequest represents TL type `messages.setHistoryTTL#b80e5fe4`.
//
// See https://core.telegram.org/method/messages.setHistoryTTL for reference.
type MessagesSetHistoryTTLRequest struct {
	// Peer field of MessagesSetHistoryTTLRequest.
	Peer InputPeerClass
	// Period field of MessagesSetHistoryTTLRequest.
	Period int
}

// MessagesSetHistoryTTLRequestTypeID is TL type id of MessagesSetHistoryTTLRequest.
const MessagesSetHistoryTTLRequestTypeID = 0xb80e5fe4

func (s *MessagesSetHistoryTTLRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.Period == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSetHistoryTTLRequest) String() string {
	if s == nil {
		return "MessagesSetHistoryTTLRequest(nil)"
	}
	type Alias MessagesSetHistoryTTLRequest
	return fmt.Sprintf("MessagesSetHistoryTTLRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSetHistoryTTLRequest from given interface.
func (s *MessagesSetHistoryTTLRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetPeriod() (value int)
}) {
	s.Peer = from.GetPeer()
	s.Period = from.GetPeriod()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSetHistoryTTLRequest) TypeID() uint32 {
	return MessagesSetHistoryTTLRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSetHistoryTTLRequest) TypeName() string {
	return "messages.setHistoryTTL"
}

// TypeInfo returns info about TL type.
func (s *MessagesSetHistoryTTLRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.setHistoryTTL",
		ID:   MessagesSetHistoryTTLRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Period",
			SchemaName: "period",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSetHistoryTTLRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode %s as nil", "messages.setHistoryTTL#b80e5fe4")
	}
	b.PutID(MessagesSetHistoryTTLRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSetHistoryTTLRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode %s as nil", "messages.setHistoryTTL#b80e5fe4")
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode %s: field %s is nil", "messages.setHistoryTTL#b80e5fe4", "peer")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "messages.setHistoryTTL#b80e5fe4", "peer", err)
	}
	b.PutInt(s.Period)
	return nil
}

// GetPeer returns value of Peer field.
func (s *MessagesSetHistoryTTLRequest) GetPeer() (value InputPeerClass) {
	return s.Peer
}

// GetPeriod returns value of Period field.
func (s *MessagesSetHistoryTTLRequest) GetPeriod() (value int) {
	return s.Period
}

// Decode implements bin.Decoder.
func (s *MessagesSetHistoryTTLRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode %s to nil", "messages.setHistoryTTL#b80e5fe4")
	}
	if err := b.ConsumeID(MessagesSetHistoryTTLRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "messages.setHistoryTTL#b80e5fe4", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSetHistoryTTLRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode %s to nil", "messages.setHistoryTTL#b80e5fe4")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.setHistoryTTL#b80e5fe4", "peer", err)
		}
		s.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.setHistoryTTL#b80e5fe4", "period", err)
		}
		s.Period = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSetHistoryTTLRequest.
var (
	_ bin.Encoder     = &MessagesSetHistoryTTLRequest{}
	_ bin.Decoder     = &MessagesSetHistoryTTLRequest{}
	_ bin.BareEncoder = &MessagesSetHistoryTTLRequest{}
	_ bin.BareDecoder = &MessagesSetHistoryTTLRequest{}
)

// MessagesSetHistoryTTL invokes method messages.setHistoryTTL#b80e5fe4 returning error if any.
//
// See https://core.telegram.org/method/messages.setHistoryTTL for reference.
func (c *Client) MessagesSetHistoryTTL(ctx context.Context, request *MessagesSetHistoryTTLRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
