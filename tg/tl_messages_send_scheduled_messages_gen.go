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

// MessagesSendScheduledMessagesRequest represents TL type `messages.sendScheduledMessages#bd38850a`.
// Send scheduled messages right away
//
// See https://core.telegram.org/method/messages.sendScheduledMessages for reference.
type MessagesSendScheduledMessagesRequest struct {
	// Peer
	Peer InputPeerClass
	// Scheduled message IDs
	ID []int
}

// MessagesSendScheduledMessagesRequestTypeID is TL type id of MessagesSendScheduledMessagesRequest.
const MessagesSendScheduledMessagesRequestTypeID = 0xbd38850a

func (s *MessagesSendScheduledMessagesRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSendScheduledMessagesRequest) String() string {
	if s == nil {
		return "MessagesSendScheduledMessagesRequest(nil)"
	}
	type Alias MessagesSendScheduledMessagesRequest
	return fmt.Sprintf("MessagesSendScheduledMessagesRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSendScheduledMessagesRequest from given interface.
func (s *MessagesSendScheduledMessagesRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetID() (value []int)
}) {
	s.Peer = from.GetPeer()
	s.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSendScheduledMessagesRequest) TypeID() uint32 {
	return MessagesSendScheduledMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSendScheduledMessagesRequest) TypeName() string {
	return "messages.sendScheduledMessages"
}

// TypeInfo returns info about TL type.
func (s *MessagesSendScheduledMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.sendScheduledMessages",
		ID:   MessagesSendScheduledMessagesRequestTypeID,
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
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSendScheduledMessagesRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode %s as nil", "messages.sendScheduledMessages#bd38850a")
	}
	b.PutID(MessagesSendScheduledMessagesRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSendScheduledMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode %s as nil", "messages.sendScheduledMessages#bd38850a")
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode %s: field %s is nil", "messages.sendScheduledMessages#bd38850a", "peer")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "messages.sendScheduledMessages#bd38850a", "peer", err)
	}
	b.PutVectorHeader(len(s.ID))
	for _, v := range s.ID {
		b.PutInt(v)
	}
	return nil
}

// GetPeer returns value of Peer field.
func (s *MessagesSendScheduledMessagesRequest) GetPeer() (value InputPeerClass) {
	return s.Peer
}

// GetID returns value of ID field.
func (s *MessagesSendScheduledMessagesRequest) GetID() (value []int) {
	return s.ID
}

// Decode implements bin.Decoder.
func (s *MessagesSendScheduledMessagesRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode %s to nil", "messages.sendScheduledMessages#bd38850a")
	}
	if err := b.ConsumeID(MessagesSendScheduledMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "messages.sendScheduledMessages#bd38850a", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSendScheduledMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode %s to nil", "messages.sendScheduledMessages#bd38850a")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.sendScheduledMessages#bd38850a", "peer", err)
		}
		s.Peer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.sendScheduledMessages#bd38850a", "id", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode %s: field %s: %w", "messages.sendScheduledMessages#bd38850a", "id", err)
			}
			s.ID = append(s.ID, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSendScheduledMessagesRequest.
var (
	_ bin.Encoder     = &MessagesSendScheduledMessagesRequest{}
	_ bin.Decoder     = &MessagesSendScheduledMessagesRequest{}
	_ bin.BareEncoder = &MessagesSendScheduledMessagesRequest{}
	_ bin.BareDecoder = &MessagesSendScheduledMessagesRequest{}
)

// MessagesSendScheduledMessages invokes method messages.sendScheduledMessages#bd38850a returning error if any.
// Send scheduled messages right away
//
// Possible errors:
//  400 MESSAGE_ID_INVALID: The provided message id is invalid
//
// See https://core.telegram.org/method/messages.sendScheduledMessages for reference.
func (c *Client) MessagesSendScheduledMessages(ctx context.Context, request *MessagesSendScheduledMessagesRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
