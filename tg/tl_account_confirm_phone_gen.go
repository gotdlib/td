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

// AccountConfirmPhoneRequest represents TL type `account.confirmPhone#5f2178c3`.
// Confirm a phone number to cancel account deletion, for more info click here »¹
//
// Links:
//  1) https://core.telegram.org/api/account-deletion
//
// See https://core.telegram.org/method/account.confirmPhone for reference.
type AccountConfirmPhoneRequest struct {
	// Phone code hash, for more info click here »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/account-deletion
	PhoneCodeHash string
	// SMS code, for more info click here »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/account-deletion
	PhoneCode string
}

// AccountConfirmPhoneRequestTypeID is TL type id of AccountConfirmPhoneRequest.
const AccountConfirmPhoneRequestTypeID = 0x5f2178c3

func (c *AccountConfirmPhoneRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.PhoneCodeHash == "") {
		return false
	}
	if !(c.PhoneCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *AccountConfirmPhoneRequest) String() string {
	if c == nil {
		return "AccountConfirmPhoneRequest(nil)"
	}
	type Alias AccountConfirmPhoneRequest
	return fmt.Sprintf("AccountConfirmPhoneRequest%+v", Alias(*c))
}

// FillFrom fills AccountConfirmPhoneRequest from given interface.
func (c *AccountConfirmPhoneRequest) FillFrom(from interface {
	GetPhoneCodeHash() (value string)
	GetPhoneCode() (value string)
}) {
	c.PhoneCodeHash = from.GetPhoneCodeHash()
	c.PhoneCode = from.GetPhoneCode()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountConfirmPhoneRequest) TypeID() uint32 {
	return AccountConfirmPhoneRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountConfirmPhoneRequest) TypeName() string {
	return "account.confirmPhone"
}

// TypeInfo returns info about TL type.
func (c *AccountConfirmPhoneRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.confirmPhone",
		ID:   AccountConfirmPhoneRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PhoneCodeHash",
			SchemaName: "phone_code_hash",
		},
		{
			Name:       "PhoneCode",
			SchemaName: "phone_code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *AccountConfirmPhoneRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode %s as nil", "account.confirmPhone#5f2178c3")
	}
	b.PutID(AccountConfirmPhoneRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *AccountConfirmPhoneRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode %s as nil", "account.confirmPhone#5f2178c3")
	}
	b.PutString(c.PhoneCodeHash)
	b.PutString(c.PhoneCode)
	return nil
}

// GetPhoneCodeHash returns value of PhoneCodeHash field.
func (c *AccountConfirmPhoneRequest) GetPhoneCodeHash() (value string) {
	return c.PhoneCodeHash
}

// GetPhoneCode returns value of PhoneCode field.
func (c *AccountConfirmPhoneRequest) GetPhoneCode() (value string) {
	return c.PhoneCode
}

// Decode implements bin.Decoder.
func (c *AccountConfirmPhoneRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode %s to nil", "account.confirmPhone#5f2178c3")
	}
	if err := b.ConsumeID(AccountConfirmPhoneRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "account.confirmPhone#5f2178c3", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *AccountConfirmPhoneRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode %s to nil", "account.confirmPhone#5f2178c3")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "account.confirmPhone#5f2178c3", "phone_code_hash", err)
		}
		c.PhoneCodeHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "account.confirmPhone#5f2178c3", "phone_code", err)
		}
		c.PhoneCode = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountConfirmPhoneRequest.
var (
	_ bin.Encoder     = &AccountConfirmPhoneRequest{}
	_ bin.Decoder     = &AccountConfirmPhoneRequest{}
	_ bin.BareEncoder = &AccountConfirmPhoneRequest{}
	_ bin.BareDecoder = &AccountConfirmPhoneRequest{}
)

// AccountConfirmPhone invokes method account.confirmPhone#5f2178c3 returning error if any.
// Confirm a phone number to cancel account deletion, for more info click here »¹
//
// Links:
//  1) https://core.telegram.org/api/account-deletion
//
// Possible errors:
//  400 CODE_HASH_INVALID: Code hash invalid
//  400 PHONE_CODE_EMPTY: phone_code is missing
//
// See https://core.telegram.org/method/account.confirmPhone for reference.
func (c *Client) AccountConfirmPhone(ctx context.Context, request *AccountConfirmPhoneRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
