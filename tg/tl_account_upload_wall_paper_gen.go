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

// AccountUploadWallPaperRequest represents TL type `account.uploadWallPaper#dd853661`.
// Create and upload a new wallpaper
//
// See https://core.telegram.org/method/account.uploadWallPaper for reference.
type AccountUploadWallPaperRequest struct {
	// The JPG/PNG wallpaper
	File InputFileClass
	// MIME type of uploaded wallpaper
	MimeType string
	// Wallpaper settings
	Settings WallPaperSettings
}

// AccountUploadWallPaperRequestTypeID is TL type id of AccountUploadWallPaperRequest.
const AccountUploadWallPaperRequestTypeID = 0xdd853661

func (u *AccountUploadWallPaperRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.File == nil) {
		return false
	}
	if !(u.MimeType == "") {
		return false
	}
	if !(u.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUploadWallPaperRequest) String() string {
	if u == nil {
		return "AccountUploadWallPaperRequest(nil)"
	}
	type Alias AccountUploadWallPaperRequest
	return fmt.Sprintf("AccountUploadWallPaperRequest%+v", Alias(*u))
}

// FillFrom fills AccountUploadWallPaperRequest from given interface.
func (u *AccountUploadWallPaperRequest) FillFrom(from interface {
	GetFile() (value InputFileClass)
	GetMimeType() (value string)
	GetSettings() (value WallPaperSettings)
}) {
	u.File = from.GetFile()
	u.MimeType = from.GetMimeType()
	u.Settings = from.GetSettings()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUploadWallPaperRequest) TypeID() uint32 {
	return AccountUploadWallPaperRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUploadWallPaperRequest) TypeName() string {
	return "account.uploadWallPaper"
}

// TypeInfo returns info about TL type.
func (u *AccountUploadWallPaperRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.uploadWallPaper",
		ID:   AccountUploadWallPaperRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "File",
			SchemaName: "file",
		},
		{
			Name:       "MimeType",
			SchemaName: "mime_type",
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *AccountUploadWallPaperRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode %s as nil", "account.uploadWallPaper#dd853661")
	}
	b.PutID(AccountUploadWallPaperRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *AccountUploadWallPaperRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode %s as nil", "account.uploadWallPaper#dd853661")
	}
	if u.File == nil {
		return fmt.Errorf("unable to encode %s: field %s is nil", "account.uploadWallPaper#dd853661", "file")
	}
	if err := u.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "account.uploadWallPaper#dd853661", "file", err)
	}
	b.PutString(u.MimeType)
	if err := u.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "account.uploadWallPaper#dd853661", "settings", err)
	}
	return nil
}

// GetFile returns value of File field.
func (u *AccountUploadWallPaperRequest) GetFile() (value InputFileClass) {
	return u.File
}

// GetMimeType returns value of MimeType field.
func (u *AccountUploadWallPaperRequest) GetMimeType() (value string) {
	return u.MimeType
}

// GetSettings returns value of Settings field.
func (u *AccountUploadWallPaperRequest) GetSettings() (value WallPaperSettings) {
	return u.Settings
}

// Decode implements bin.Decoder.
func (u *AccountUploadWallPaperRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode %s to nil", "account.uploadWallPaper#dd853661")
	}
	if err := b.ConsumeID(AccountUploadWallPaperRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "account.uploadWallPaper#dd853661", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *AccountUploadWallPaperRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode %s to nil", "account.uploadWallPaper#dd853661")
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "account.uploadWallPaper#dd853661", "file", err)
		}
		u.File = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "account.uploadWallPaper#dd853661", "mime_type", err)
		}
		u.MimeType = value
	}
	{
		if err := u.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "account.uploadWallPaper#dd853661", "settings", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountUploadWallPaperRequest.
var (
	_ bin.Encoder     = &AccountUploadWallPaperRequest{}
	_ bin.Decoder     = &AccountUploadWallPaperRequest{}
	_ bin.BareEncoder = &AccountUploadWallPaperRequest{}
	_ bin.BareDecoder = &AccountUploadWallPaperRequest{}
)

// AccountUploadWallPaper invokes method account.uploadWallPaper#dd853661 returning error if any.
// Create and upload a new wallpaper
//
// See https://core.telegram.org/method/account.uploadWallPaper for reference.
func (c *Client) AccountUploadWallPaper(ctx context.Context, request *AccountUploadWallPaperRequest) (WallPaperClass, error) {
	var result WallPaperBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.WallPaper, nil
}
