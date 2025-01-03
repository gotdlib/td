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

// UpgradeGiftResult represents TL type `upgradeGiftResult#3d9e8331`.
type UpgradeGiftResult struct {
	// The upgraded gift
	Gift UpgradedGift
	// True, if the gift is displayed on the user's profile page
	IsSaved bool
	// True, if the gift can be transferred to another user
	CanBeTransferred bool
	// Number of Telegram Stars that must be paid to transfer the upgraded gift
	TransferStarCount int64
	// Point in time (Unix timestamp) when the gift can be transferred to TON blockchain as
	// an NFT
	ExportDate int32
}

// UpgradeGiftResultTypeID is TL type id of UpgradeGiftResult.
const UpgradeGiftResultTypeID = 0x3d9e8331

// Ensuring interfaces in compile-time for UpgradeGiftResult.
var (
	_ bin.Encoder     = &UpgradeGiftResult{}
	_ bin.Decoder     = &UpgradeGiftResult{}
	_ bin.BareEncoder = &UpgradeGiftResult{}
	_ bin.BareDecoder = &UpgradeGiftResult{}
)

func (u *UpgradeGiftResult) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Gift.Zero()) {
		return false
	}
	if !(u.IsSaved == false) {
		return false
	}
	if !(u.CanBeTransferred == false) {
		return false
	}
	if !(u.TransferStarCount == 0) {
		return false
	}
	if !(u.ExportDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UpgradeGiftResult) String() string {
	if u == nil {
		return "UpgradeGiftResult(nil)"
	}
	type Alias UpgradeGiftResult
	return fmt.Sprintf("UpgradeGiftResult%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UpgradeGiftResult) TypeID() uint32 {
	return UpgradeGiftResultTypeID
}

// TypeName returns name of type in TL schema.
func (*UpgradeGiftResult) TypeName() string {
	return "upgradeGiftResult"
}

// TypeInfo returns info about TL type.
func (u *UpgradeGiftResult) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "upgradeGiftResult",
		ID:   UpgradeGiftResultTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Gift",
			SchemaName: "gift",
		},
		{
			Name:       "IsSaved",
			SchemaName: "is_saved",
		},
		{
			Name:       "CanBeTransferred",
			SchemaName: "can_be_transferred",
		},
		{
			Name:       "TransferStarCount",
			SchemaName: "transfer_star_count",
		},
		{
			Name:       "ExportDate",
			SchemaName: "export_date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *UpgradeGiftResult) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode upgradeGiftResult#3d9e8331 as nil")
	}
	b.PutID(UpgradeGiftResultTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UpgradeGiftResult) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode upgradeGiftResult#3d9e8331 as nil")
	}
	if err := u.Gift.Encode(b); err != nil {
		return fmt.Errorf("unable to encode upgradeGiftResult#3d9e8331: field gift: %w", err)
	}
	b.PutBool(u.IsSaved)
	b.PutBool(u.CanBeTransferred)
	b.PutInt53(u.TransferStarCount)
	b.PutInt32(u.ExportDate)
	return nil
}

// Decode implements bin.Decoder.
func (u *UpgradeGiftResult) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode upgradeGiftResult#3d9e8331 to nil")
	}
	if err := b.ConsumeID(UpgradeGiftResultTypeID); err != nil {
		return fmt.Errorf("unable to decode upgradeGiftResult#3d9e8331: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UpgradeGiftResult) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode upgradeGiftResult#3d9e8331 to nil")
	}
	{
		if err := u.Gift.Decode(b); err != nil {
			return fmt.Errorf("unable to decode upgradeGiftResult#3d9e8331: field gift: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode upgradeGiftResult#3d9e8331: field is_saved: %w", err)
		}
		u.IsSaved = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode upgradeGiftResult#3d9e8331: field can_be_transferred: %w", err)
		}
		u.CanBeTransferred = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode upgradeGiftResult#3d9e8331: field transfer_star_count: %w", err)
		}
		u.TransferStarCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode upgradeGiftResult#3d9e8331: field export_date: %w", err)
		}
		u.ExportDate = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (u *UpgradeGiftResult) EncodeTDLibJSON(b tdjson.Encoder) error {
	if u == nil {
		return fmt.Errorf("can't encode upgradeGiftResult#3d9e8331 as nil")
	}
	b.ObjStart()
	b.PutID("upgradeGiftResult")
	b.Comma()
	b.FieldStart("gift")
	if err := u.Gift.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode upgradeGiftResult#3d9e8331: field gift: %w", err)
	}
	b.Comma()
	b.FieldStart("is_saved")
	b.PutBool(u.IsSaved)
	b.Comma()
	b.FieldStart("can_be_transferred")
	b.PutBool(u.CanBeTransferred)
	b.Comma()
	b.FieldStart("transfer_star_count")
	b.PutInt53(u.TransferStarCount)
	b.Comma()
	b.FieldStart("export_date")
	b.PutInt32(u.ExportDate)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (u *UpgradeGiftResult) DecodeTDLibJSON(b tdjson.Decoder) error {
	if u == nil {
		return fmt.Errorf("can't decode upgradeGiftResult#3d9e8331 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("upgradeGiftResult"); err != nil {
				return fmt.Errorf("unable to decode upgradeGiftResult#3d9e8331: %w", err)
			}
		case "gift":
			if err := u.Gift.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode upgradeGiftResult#3d9e8331: field gift: %w", err)
			}
		case "is_saved":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode upgradeGiftResult#3d9e8331: field is_saved: %w", err)
			}
			u.IsSaved = value
		case "can_be_transferred":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode upgradeGiftResult#3d9e8331: field can_be_transferred: %w", err)
			}
			u.CanBeTransferred = value
		case "transfer_star_count":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode upgradeGiftResult#3d9e8331: field transfer_star_count: %w", err)
			}
			u.TransferStarCount = value
		case "export_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode upgradeGiftResult#3d9e8331: field export_date: %w", err)
			}
			u.ExportDate = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetGift returns value of Gift field.
func (u *UpgradeGiftResult) GetGift() (value UpgradedGift) {
	if u == nil {
		return
	}
	return u.Gift
}

// GetIsSaved returns value of IsSaved field.
func (u *UpgradeGiftResult) GetIsSaved() (value bool) {
	if u == nil {
		return
	}
	return u.IsSaved
}

// GetCanBeTransferred returns value of CanBeTransferred field.
func (u *UpgradeGiftResult) GetCanBeTransferred() (value bool) {
	if u == nil {
		return
	}
	return u.CanBeTransferred
}

// GetTransferStarCount returns value of TransferStarCount field.
func (u *UpgradeGiftResult) GetTransferStarCount() (value int64) {
	if u == nil {
		return
	}
	return u.TransferStarCount
}

// GetExportDate returns value of ExportDate field.
func (u *UpgradeGiftResult) GetExportDate() (value int32) {
	if u == nil {
		return
	}
	return u.ExportDate
}
