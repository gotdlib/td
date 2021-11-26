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

// AddStickerToSetRequest represents TL type `addStickerToSet#3b0b81`.
type AddStickerToSetRequest struct {
	// Sticker set owner
	UserID int64
	// Sticker set name
	Name string
	// Sticker to add to the set
	Sticker InputStickerClass
}

// AddStickerToSetRequestTypeID is TL type id of AddStickerToSetRequest.
const AddStickerToSetRequestTypeID = 0x3b0b81

// Ensuring interfaces in compile-time for AddStickerToSetRequest.
var (
	_ bin.Encoder     = &AddStickerToSetRequest{}
	_ bin.Decoder     = &AddStickerToSetRequest{}
	_ bin.BareEncoder = &AddStickerToSetRequest{}
	_ bin.BareDecoder = &AddStickerToSetRequest{}
)

func (a *AddStickerToSetRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.UserID == 0) {
		return false
	}
	if !(a.Name == "") {
		return false
	}
	if !(a.Sticker == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AddStickerToSetRequest) String() string {
	if a == nil {
		return "AddStickerToSetRequest(nil)"
	}
	type Alias AddStickerToSetRequest
	return fmt.Sprintf("AddStickerToSetRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AddStickerToSetRequest) TypeID() uint32 {
	return AddStickerToSetRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AddStickerToSetRequest) TypeName() string {
	return "addStickerToSet"
}

// TypeInfo returns info about TL type.
func (a *AddStickerToSetRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "addStickerToSet",
		ID:   AddStickerToSetRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "Sticker",
			SchemaName: "sticker",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AddStickerToSetRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addStickerToSet#3b0b81 as nil")
	}
	b.PutID(AddStickerToSetRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AddStickerToSetRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addStickerToSet#3b0b81 as nil")
	}
	b.PutLong(a.UserID)
	b.PutString(a.Name)
	if a.Sticker == nil {
		return fmt.Errorf("unable to encode addStickerToSet#3b0b81: field sticker is nil")
	}
	if err := a.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode addStickerToSet#3b0b81: field sticker: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AddStickerToSetRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addStickerToSet#3b0b81 to nil")
	}
	if err := b.ConsumeID(AddStickerToSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode addStickerToSet#3b0b81: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AddStickerToSetRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addStickerToSet#3b0b81 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode addStickerToSet#3b0b81: field user_id: %w", err)
		}
		a.UserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode addStickerToSet#3b0b81: field name: %w", err)
		}
		a.Name = value
	}
	{
		value, err := DecodeInputSticker(b)
		if err != nil {
			return fmt.Errorf("unable to decode addStickerToSet#3b0b81: field sticker: %w", err)
		}
		a.Sticker = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AddStickerToSetRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode addStickerToSet#3b0b81 as nil")
	}
	b.ObjStart()
	b.PutID("addStickerToSet")
	b.FieldStart("user_id")
	b.PutLong(a.UserID)
	b.FieldStart("name")
	b.PutString(a.Name)
	b.FieldStart("sticker")
	if a.Sticker == nil {
		return fmt.Errorf("unable to encode addStickerToSet#3b0b81: field sticker is nil")
	}
	if err := a.Sticker.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode addStickerToSet#3b0b81: field sticker: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AddStickerToSetRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode addStickerToSet#3b0b81 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("addStickerToSet"); err != nil {
				return fmt.Errorf("unable to decode addStickerToSet#3b0b81: %w", err)
			}
		case "user_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode addStickerToSet#3b0b81: field user_id: %w", err)
			}
			a.UserID = value
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode addStickerToSet#3b0b81: field name: %w", err)
			}
			a.Name = value
		case "sticker":
			value, err := DecodeTDLibJSONInputSticker(b)
			if err != nil {
				return fmt.Errorf("unable to decode addStickerToSet#3b0b81: field sticker: %w", err)
			}
			a.Sticker = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (a *AddStickerToSetRequest) GetUserID() (value int64) {
	return a.UserID
}

// GetName returns value of Name field.
func (a *AddStickerToSetRequest) GetName() (value string) {
	return a.Name
}

// GetSticker returns value of Sticker field.
func (a *AddStickerToSetRequest) GetSticker() (value InputStickerClass) {
	return a.Sticker
}

// AddStickerToSet invokes method addStickerToSet#3b0b81 returning error if any.
func (c *Client) AddStickerToSet(ctx context.Context, request *AddStickerToSetRequest) (*StickerSet, error) {
	var result StickerSet

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
