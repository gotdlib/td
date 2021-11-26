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

// EditMessageReplyMarkupRequest represents TL type `editMessageReplyMarkup#13cbde89`.
type EditMessageReplyMarkupRequest struct {
	// The chat the message belongs to
	ChatID int64
	// Identifier of the message
	MessageID int64
	// The new message reply markup; pass null if none
	ReplyMarkup ReplyMarkupClass
}

// EditMessageReplyMarkupRequestTypeID is TL type id of EditMessageReplyMarkupRequest.
const EditMessageReplyMarkupRequestTypeID = 0x13cbde89

// Ensuring interfaces in compile-time for EditMessageReplyMarkupRequest.
var (
	_ bin.Encoder     = &EditMessageReplyMarkupRequest{}
	_ bin.Decoder     = &EditMessageReplyMarkupRequest{}
	_ bin.BareEncoder = &EditMessageReplyMarkupRequest{}
	_ bin.BareDecoder = &EditMessageReplyMarkupRequest{}
)

func (e *EditMessageReplyMarkupRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ChatID == 0) {
		return false
	}
	if !(e.MessageID == 0) {
		return false
	}
	if !(e.ReplyMarkup == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EditMessageReplyMarkupRequest) String() string {
	if e == nil {
		return "EditMessageReplyMarkupRequest(nil)"
	}
	type Alias EditMessageReplyMarkupRequest
	return fmt.Sprintf("EditMessageReplyMarkupRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EditMessageReplyMarkupRequest) TypeID() uint32 {
	return EditMessageReplyMarkupRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EditMessageReplyMarkupRequest) TypeName() string {
	return "editMessageReplyMarkup"
}

// TypeInfo returns info about TL type.
func (e *EditMessageReplyMarkupRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "editMessageReplyMarkup",
		ID:   EditMessageReplyMarkupRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "ReplyMarkup",
			SchemaName: "reply_markup",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EditMessageReplyMarkupRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editMessageReplyMarkup#13cbde89 as nil")
	}
	b.PutID(EditMessageReplyMarkupRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EditMessageReplyMarkupRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editMessageReplyMarkup#13cbde89 as nil")
	}
	b.PutLong(e.ChatID)
	b.PutLong(e.MessageID)
	if e.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode editMessageReplyMarkup#13cbde89: field reply_markup is nil")
	}
	if err := e.ReplyMarkup.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editMessageReplyMarkup#13cbde89: field reply_markup: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EditMessageReplyMarkupRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editMessageReplyMarkup#13cbde89 to nil")
	}
	if err := b.ConsumeID(EditMessageReplyMarkupRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode editMessageReplyMarkup#13cbde89: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EditMessageReplyMarkupRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editMessageReplyMarkup#13cbde89 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode editMessageReplyMarkup#13cbde89: field chat_id: %w", err)
		}
		e.ChatID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode editMessageReplyMarkup#13cbde89: field message_id: %w", err)
		}
		e.MessageID = value
	}
	{
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode editMessageReplyMarkup#13cbde89: field reply_markup: %w", err)
		}
		e.ReplyMarkup = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EditMessageReplyMarkupRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode editMessageReplyMarkup#13cbde89 as nil")
	}
	b.ObjStart()
	b.PutID("editMessageReplyMarkup")
	b.FieldStart("chat_id")
	b.PutLong(e.ChatID)
	b.FieldStart("message_id")
	b.PutLong(e.MessageID)
	b.FieldStart("reply_markup")
	if e.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode editMessageReplyMarkup#13cbde89: field reply_markup is nil")
	}
	if err := e.ReplyMarkup.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editMessageReplyMarkup#13cbde89: field reply_markup: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EditMessageReplyMarkupRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode editMessageReplyMarkup#13cbde89 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("editMessageReplyMarkup"); err != nil {
				return fmt.Errorf("unable to decode editMessageReplyMarkup#13cbde89: %w", err)
			}
		case "chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode editMessageReplyMarkup#13cbde89: field chat_id: %w", err)
			}
			e.ChatID = value
		case "message_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode editMessageReplyMarkup#13cbde89: field message_id: %w", err)
			}
			e.MessageID = value
		case "reply_markup":
			value, err := DecodeTDLibJSONReplyMarkup(b)
			if err != nil {
				return fmt.Errorf("unable to decode editMessageReplyMarkup#13cbde89: field reply_markup: %w", err)
			}
			e.ReplyMarkup = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (e *EditMessageReplyMarkupRequest) GetChatID() (value int64) {
	return e.ChatID
}

// GetMessageID returns value of MessageID field.
func (e *EditMessageReplyMarkupRequest) GetMessageID() (value int64) {
	return e.MessageID
}

// GetReplyMarkup returns value of ReplyMarkup field.
func (e *EditMessageReplyMarkupRequest) GetReplyMarkup() (value ReplyMarkupClass) {
	return e.ReplyMarkup
}

// EditMessageReplyMarkup invokes method editMessageReplyMarkup#13cbde89 returning error if any.
func (c *Client) EditMessageReplyMarkup(ctx context.Context, request *EditMessageReplyMarkupRequest) (*Message, error) {
	var result Message

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
