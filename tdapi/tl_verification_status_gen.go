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

// VerificationStatus represents TL type `verificationStatus#3ae6a18c`.
type VerificationStatus struct {
	// True, if the chat or the user is verified by Telegram
	IsVerified bool
	// True, if the chat or the user is marked as scam by Telegram
	IsScam bool
	// True, if the chat or the user is marked as fake by Telegram
	IsFake bool
	// Identifier of the custom emoji to be shown as verification sign provided by a bot for
	// the user; 0 if none
	BotVerificationIconCustomEmojiID int64
}

// VerificationStatusTypeID is TL type id of VerificationStatus.
const VerificationStatusTypeID = 0x3ae6a18c

// Ensuring interfaces in compile-time for VerificationStatus.
var (
	_ bin.Encoder     = &VerificationStatus{}
	_ bin.Decoder     = &VerificationStatus{}
	_ bin.BareEncoder = &VerificationStatus{}
	_ bin.BareDecoder = &VerificationStatus{}
)

func (v *VerificationStatus) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.IsVerified == false) {
		return false
	}
	if !(v.IsScam == false) {
		return false
	}
	if !(v.IsFake == false) {
		return false
	}
	if !(v.BotVerificationIconCustomEmojiID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *VerificationStatus) String() string {
	if v == nil {
		return "VerificationStatus(nil)"
	}
	type Alias VerificationStatus
	return fmt.Sprintf("VerificationStatus%+v", Alias(*v))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*VerificationStatus) TypeID() uint32 {
	return VerificationStatusTypeID
}

// TypeName returns name of type in TL schema.
func (*VerificationStatus) TypeName() string {
	return "verificationStatus"
}

// TypeInfo returns info about TL type.
func (v *VerificationStatus) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "verificationStatus",
		ID:   VerificationStatusTypeID,
	}
	if v == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "IsVerified",
			SchemaName: "is_verified",
		},
		{
			Name:       "IsScam",
			SchemaName: "is_scam",
		},
		{
			Name:       "IsFake",
			SchemaName: "is_fake",
		},
		{
			Name:       "BotVerificationIconCustomEmojiID",
			SchemaName: "bot_verification_icon_custom_emoji_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (v *VerificationStatus) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode verificationStatus#3ae6a18c as nil")
	}
	b.PutID(VerificationStatusTypeID)
	return v.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (v *VerificationStatus) EncodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode verificationStatus#3ae6a18c as nil")
	}
	b.PutBool(v.IsVerified)
	b.PutBool(v.IsScam)
	b.PutBool(v.IsFake)
	b.PutLong(v.BotVerificationIconCustomEmojiID)
	return nil
}

// Decode implements bin.Decoder.
func (v *VerificationStatus) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode verificationStatus#3ae6a18c to nil")
	}
	if err := b.ConsumeID(VerificationStatusTypeID); err != nil {
		return fmt.Errorf("unable to decode verificationStatus#3ae6a18c: %w", err)
	}
	return v.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (v *VerificationStatus) DecodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode verificationStatus#3ae6a18c to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode verificationStatus#3ae6a18c: field is_verified: %w", err)
		}
		v.IsVerified = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode verificationStatus#3ae6a18c: field is_scam: %w", err)
		}
		v.IsScam = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode verificationStatus#3ae6a18c: field is_fake: %w", err)
		}
		v.IsFake = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode verificationStatus#3ae6a18c: field bot_verification_icon_custom_emoji_id: %w", err)
		}
		v.BotVerificationIconCustomEmojiID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (v *VerificationStatus) EncodeTDLibJSON(b tdjson.Encoder) error {
	if v == nil {
		return fmt.Errorf("can't encode verificationStatus#3ae6a18c as nil")
	}
	b.ObjStart()
	b.PutID("verificationStatus")
	b.Comma()
	b.FieldStart("is_verified")
	b.PutBool(v.IsVerified)
	b.Comma()
	b.FieldStart("is_scam")
	b.PutBool(v.IsScam)
	b.Comma()
	b.FieldStart("is_fake")
	b.PutBool(v.IsFake)
	b.Comma()
	b.FieldStart("bot_verification_icon_custom_emoji_id")
	b.PutLong(v.BotVerificationIconCustomEmojiID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (v *VerificationStatus) DecodeTDLibJSON(b tdjson.Decoder) error {
	if v == nil {
		return fmt.Errorf("can't decode verificationStatus#3ae6a18c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("verificationStatus"); err != nil {
				return fmt.Errorf("unable to decode verificationStatus#3ae6a18c: %w", err)
			}
		case "is_verified":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode verificationStatus#3ae6a18c: field is_verified: %w", err)
			}
			v.IsVerified = value
		case "is_scam":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode verificationStatus#3ae6a18c: field is_scam: %w", err)
			}
			v.IsScam = value
		case "is_fake":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode verificationStatus#3ae6a18c: field is_fake: %w", err)
			}
			v.IsFake = value
		case "bot_verification_icon_custom_emoji_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode verificationStatus#3ae6a18c: field bot_verification_icon_custom_emoji_id: %w", err)
			}
			v.BotVerificationIconCustomEmojiID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetIsVerified returns value of IsVerified field.
func (v *VerificationStatus) GetIsVerified() (value bool) {
	if v == nil {
		return
	}
	return v.IsVerified
}

// GetIsScam returns value of IsScam field.
func (v *VerificationStatus) GetIsScam() (value bool) {
	if v == nil {
		return
	}
	return v.IsScam
}

// GetIsFake returns value of IsFake field.
func (v *VerificationStatus) GetIsFake() (value bool) {
	if v == nil {
		return
	}
	return v.IsFake
}

// GetBotVerificationIconCustomEmojiID returns value of BotVerificationIconCustomEmojiID field.
func (v *VerificationStatus) GetBotVerificationIconCustomEmojiID() (value int64) {
	if v == nil {
		return
	}
	return v.BotVerificationIconCustomEmojiID
}
