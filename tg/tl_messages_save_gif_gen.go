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

// MessagesSaveGifRequest represents TL type `messages.saveGif#327a30cb`.
// Add GIF to saved gifs list
//
// See https://core.telegram.org/method/messages.saveGif for reference.
type MessagesSaveGifRequest struct {
	// GIF to save
	ID InputDocumentClass
	// Whether to remove GIF from saved gifs list
	Unsave bool
}

// MessagesSaveGifRequestTypeID is TL type id of MessagesSaveGifRequest.
const MessagesSaveGifRequestTypeID = 0x327a30cb

func (s *MessagesSaveGifRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ID == nil) {
		return false
	}
	if !(s.Unsave == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSaveGifRequest) String() string {
	if s == nil {
		return "MessagesSaveGifRequest(nil)"
	}
	type Alias MessagesSaveGifRequest
	return fmt.Sprintf("MessagesSaveGifRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSaveGifRequest from given interface.
func (s *MessagesSaveGifRequest) FillFrom(from interface {
	GetID() (value InputDocumentClass)
	GetUnsave() (value bool)
}) {
	s.ID = from.GetID()
	s.Unsave = from.GetUnsave()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSaveGifRequest) TypeID() uint32 {
	return MessagesSaveGifRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSaveGifRequest) TypeName() string {
	return "messages.saveGif"
}

// TypeInfo returns info about TL type.
func (s *MessagesSaveGifRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.saveGif",
		ID:   MessagesSaveGifRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Unsave",
			SchemaName: "unsave",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSaveGifRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode %s as nil", "messages.saveGif#327a30cb")
	}
	b.PutID(MessagesSaveGifRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSaveGifRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode %s as nil", "messages.saveGif#327a30cb")
	}
	if s.ID == nil {
		return fmt.Errorf("unable to encode %s: field %s is nil", "messages.saveGif#327a30cb", "id")
	}
	if err := s.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "messages.saveGif#327a30cb", "id", err)
	}
	b.PutBool(s.Unsave)
	return nil
}

// GetID returns value of ID field.
func (s *MessagesSaveGifRequest) GetID() (value InputDocumentClass) {
	return s.ID
}

// GetIDAsNotEmpty returns mapped value of ID field.
func (s *MessagesSaveGifRequest) GetIDAsNotEmpty() (*InputDocument, bool) {
	return s.ID.AsNotEmpty()
}

// GetUnsave returns value of Unsave field.
func (s *MessagesSaveGifRequest) GetUnsave() (value bool) {
	return s.Unsave
}

// Decode implements bin.Decoder.
func (s *MessagesSaveGifRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode %s to nil", "messages.saveGif#327a30cb")
	}
	if err := b.ConsumeID(MessagesSaveGifRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "messages.saveGif#327a30cb", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSaveGifRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode %s to nil", "messages.saveGif#327a30cb")
	}
	{
		value, err := DecodeInputDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.saveGif#327a30cb", "id", err)
		}
		s.ID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.saveGif#327a30cb", "unsave", err)
		}
		s.Unsave = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSaveGifRequest.
var (
	_ bin.Encoder     = &MessagesSaveGifRequest{}
	_ bin.Decoder     = &MessagesSaveGifRequest{}
	_ bin.BareEncoder = &MessagesSaveGifRequest{}
	_ bin.BareDecoder = &MessagesSaveGifRequest{}
)

// MessagesSaveGif invokes method messages.saveGif#327a30cb returning error if any.
// Add GIF to saved gifs list
//
// Possible errors:
//  400 GIF_ID_INVALID: The provided GIF ID is invalid
//
// See https://core.telegram.org/method/messages.saveGif for reference.
func (c *Client) MessagesSaveGif(ctx context.Context, request *MessagesSaveGifRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
