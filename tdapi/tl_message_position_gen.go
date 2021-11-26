// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// MessagePosition represents TL type `messagePosition#4d0540ef`.
type MessagePosition struct {
	// 0-based message position in the full list of suitable messages
	Position int32
	// Message identifier
	MessageID int64
	// Point in time (Unix timestamp) when the message was sent
	Date int32
}

// MessagePositionTypeID is TL type id of MessagePosition.
const MessagePositionTypeID = 0x4d0540ef

// Ensuring interfaces in compile-time for MessagePosition.
var (
	_ bin.Encoder     = &MessagePosition{}
	_ bin.Decoder     = &MessagePosition{}
	_ bin.BareEncoder = &MessagePosition{}
	_ bin.BareDecoder = &MessagePosition{}
)

func (m *MessagePosition) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Position == 0) {
		return false
	}
	if !(m.MessageID == 0) {
		return false
	}
	if !(m.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessagePosition) String() string {
	if m == nil {
		return "MessagePosition(nil)"
	}
	type Alias MessagePosition
	return fmt.Sprintf("MessagePosition%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagePosition) TypeID() uint32 {
	return MessagePositionTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagePosition) TypeName() string {
	return "messagePosition"
}

// TypeInfo returns info about TL type.
func (m *MessagePosition) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messagePosition",
		ID:   MessagePositionTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Position",
			SchemaName: "position",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessagePosition) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messagePosition#4d0540ef as nil")
	}
	b.PutID(MessagePositionTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessagePosition) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messagePosition#4d0540ef as nil")
	}
	b.PutInt32(m.Position)
	b.PutLong(m.MessageID)
	b.PutInt32(m.Date)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessagePosition) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messagePosition#4d0540ef to nil")
	}
	if err := b.ConsumeID(MessagePositionTypeID); err != nil {
		return fmt.Errorf("unable to decode messagePosition#4d0540ef: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessagePosition) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messagePosition#4d0540ef to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode messagePosition#4d0540ef: field position: %w", err)
		}
		m.Position = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messagePosition#4d0540ef: field message_id: %w", err)
		}
		m.MessageID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode messagePosition#4d0540ef: field date: %w", err)
		}
		m.Date = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessagePosition) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messagePosition#4d0540ef as nil")
	}
	b.ObjStart()
	b.PutID("messagePosition")
	b.FieldStart("position")
	b.PutInt32(m.Position)
	b.FieldStart("message_id")
	b.PutLong(m.MessageID)
	b.FieldStart("date")
	b.PutInt32(m.Date)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *MessagePosition) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messagePosition#4d0540ef to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messagePosition"); err != nil {
				return fmt.Errorf("unable to decode messagePosition#4d0540ef: %w", err)
			}
		case "position":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode messagePosition#4d0540ef: field position: %w", err)
			}
			m.Position = value
		case "message_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode messagePosition#4d0540ef: field message_id: %w", err)
			}
			m.MessageID = value
		case "date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode messagePosition#4d0540ef: field date: %w", err)
			}
			m.Date = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPosition returns value of Position field.
func (m *MessagePosition) GetPosition() (value int32) {
	return m.Position
}

// GetMessageID returns value of MessageID field.
func (m *MessagePosition) GetMessageID() (value int64) {
	return m.MessageID
}

// GetDate returns value of Date field.
func (m *MessagePosition) GetDate() (value int32) {
	return m.Date
}
