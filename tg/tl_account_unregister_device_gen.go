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

// AccountUnregisterDeviceRequest represents TL type `account.unregisterDevice#3076c4bf`.
// Deletes a device by its token, stops sending PUSH-notifications to it.
//
// See https://core.telegram.org/method/account.unregisterDevice for reference.
type AccountUnregisterDeviceRequest struct {
	// Device token type.Possible values:1 - APNS (device token for apple push)2 - FCM
	// (firebase token for google firebase)3 - MPNS (channel URI for microsoft push)4 -
	// Simple push (endpoint for firefox's simple push API)5 - Ubuntu phone (token for ubuntu
	// push)6 - Blackberry (token for blackberry push)7 - Unused8 - WNS (windows push)9 -
	// APNS VoIP (token for apple push VoIP)10 - Web push (web push, see below)11 - MPNS VoIP
	// (token for microsoft push VoIP)12 - Tizen (token for tizen push)For 10 web push, the
	// token must be a JSON-encoded object containing the keys described in PUSH updates¹
	//
	// Links:
	//  1) https://core.telegram.org/api/push-updates
	TokenType int
	// Device token
	Token string
	// List of user identifiers of other users currently using the client
	OtherUIDs []int
}

// AccountUnregisterDeviceRequestTypeID is TL type id of AccountUnregisterDeviceRequest.
const AccountUnregisterDeviceRequestTypeID = 0x3076c4bf

func (u *AccountUnregisterDeviceRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.TokenType == 0) {
		return false
	}
	if !(u.Token == "") {
		return false
	}
	if !(u.OtherUIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUnregisterDeviceRequest) String() string {
	if u == nil {
		return "AccountUnregisterDeviceRequest(nil)"
	}
	type Alias AccountUnregisterDeviceRequest
	return fmt.Sprintf("AccountUnregisterDeviceRequest%+v", Alias(*u))
}

// FillFrom fills AccountUnregisterDeviceRequest from given interface.
func (u *AccountUnregisterDeviceRequest) FillFrom(from interface {
	GetTokenType() (value int)
	GetToken() (value string)
	GetOtherUIDs() (value []int)
}) {
	u.TokenType = from.GetTokenType()
	u.Token = from.GetToken()
	u.OtherUIDs = from.GetOtherUIDs()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUnregisterDeviceRequest) TypeID() uint32 {
	return AccountUnregisterDeviceRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUnregisterDeviceRequest) TypeName() string {
	return "account.unregisterDevice"
}

// TypeInfo returns info about TL type.
func (u *AccountUnregisterDeviceRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.unregisterDevice",
		ID:   AccountUnregisterDeviceRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TokenType",
			SchemaName: "token_type",
		},
		{
			Name:       "Token",
			SchemaName: "token",
		},
		{
			Name:       "OtherUIDs",
			SchemaName: "other_uids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *AccountUnregisterDeviceRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode %s as nil", "account.unregisterDevice#3076c4bf")
	}
	b.PutID(AccountUnregisterDeviceRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *AccountUnregisterDeviceRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode %s as nil", "account.unregisterDevice#3076c4bf")
	}
	b.PutInt(u.TokenType)
	b.PutString(u.Token)
	b.PutVectorHeader(len(u.OtherUIDs))
	for _, v := range u.OtherUIDs {
		b.PutInt(v)
	}
	return nil
}

// GetTokenType returns value of TokenType field.
func (u *AccountUnregisterDeviceRequest) GetTokenType() (value int) {
	return u.TokenType
}

// GetToken returns value of Token field.
func (u *AccountUnregisterDeviceRequest) GetToken() (value string) {
	return u.Token
}

// GetOtherUIDs returns value of OtherUIDs field.
func (u *AccountUnregisterDeviceRequest) GetOtherUIDs() (value []int) {
	return u.OtherUIDs
}

// Decode implements bin.Decoder.
func (u *AccountUnregisterDeviceRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode %s to nil", "account.unregisterDevice#3076c4bf")
	}
	if err := b.ConsumeID(AccountUnregisterDeviceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "account.unregisterDevice#3076c4bf", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *AccountUnregisterDeviceRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode %s to nil", "account.unregisterDevice#3076c4bf")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "account.unregisterDevice#3076c4bf", "token_type", err)
		}
		u.TokenType = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "account.unregisterDevice#3076c4bf", "token", err)
		}
		u.Token = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "account.unregisterDevice#3076c4bf", "other_uids", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode %s: field %s: %w", "account.unregisterDevice#3076c4bf", "other_uids", err)
			}
			u.OtherUIDs = append(u.OtherUIDs, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountUnregisterDeviceRequest.
var (
	_ bin.Encoder     = &AccountUnregisterDeviceRequest{}
	_ bin.Decoder     = &AccountUnregisterDeviceRequest{}
	_ bin.BareEncoder = &AccountUnregisterDeviceRequest{}
	_ bin.BareDecoder = &AccountUnregisterDeviceRequest{}
)

// AccountUnregisterDevice invokes method account.unregisterDevice#3076c4bf returning error if any.
// Deletes a device by its token, stops sending PUSH-notifications to it.
//
// Possible errors:
//  400 TOKEN_INVALID: The provided token is invalid
//
// See https://core.telegram.org/method/account.unregisterDevice for reference.
func (c *Client) AccountUnregisterDevice(ctx context.Context, request *AccountUnregisterDeviceRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
