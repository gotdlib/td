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

// AuthPasswordRecovery represents TL type `auth.passwordRecovery#137948a5`.
// Recovery info of a 2FA password¹, only for accounts with a recovery email
// configured².
//
// Links:
//  1) https://core.telegram.org/api/srp
//  2) https://core.telegram.org/api/srp#email-verification
//
// See https://core.telegram.org/constructor/auth.passwordRecovery for reference.
type AuthPasswordRecovery struct {
	// The email to which the recovery code was sent must match this pattern¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/pattern
	EmailPattern string
}

// AuthPasswordRecoveryTypeID is TL type id of AuthPasswordRecovery.
const AuthPasswordRecoveryTypeID = 0x137948a5

func (p *AuthPasswordRecovery) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.EmailPattern == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *AuthPasswordRecovery) String() string {
	if p == nil {
		return "AuthPasswordRecovery(nil)"
	}
	type Alias AuthPasswordRecovery
	return fmt.Sprintf("AuthPasswordRecovery%+v", Alias(*p))
}

// FillFrom fills AuthPasswordRecovery from given interface.
func (p *AuthPasswordRecovery) FillFrom(from interface {
	GetEmailPattern() (value string)
}) {
	p.EmailPattern = from.GetEmailPattern()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthPasswordRecovery) TypeID() uint32 {
	return AuthPasswordRecoveryTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthPasswordRecovery) TypeName() string {
	return "auth.passwordRecovery"
}

// TypeInfo returns info about TL type.
func (p *AuthPasswordRecovery) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.passwordRecovery",
		ID:   AuthPasswordRecoveryTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "EmailPattern",
			SchemaName: "email_pattern",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *AuthPasswordRecovery) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode %s as nil", "auth.passwordRecovery#137948a5")
	}
	b.PutID(AuthPasswordRecoveryTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *AuthPasswordRecovery) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode %s as nil", "auth.passwordRecovery#137948a5")
	}
	b.PutString(p.EmailPattern)
	return nil
}

// GetEmailPattern returns value of EmailPattern field.
func (p *AuthPasswordRecovery) GetEmailPattern() (value string) {
	return p.EmailPattern
}

// Decode implements bin.Decoder.
func (p *AuthPasswordRecovery) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode %s to nil", "auth.passwordRecovery#137948a5")
	}
	if err := b.ConsumeID(AuthPasswordRecoveryTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "auth.passwordRecovery#137948a5", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *AuthPasswordRecovery) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode %s to nil", "auth.passwordRecovery#137948a5")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "auth.passwordRecovery#137948a5", "email_pattern", err)
		}
		p.EmailPattern = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthPasswordRecovery.
var (
	_ bin.Encoder     = &AuthPasswordRecovery{}
	_ bin.Decoder     = &AuthPasswordRecovery{}
	_ bin.BareEncoder = &AuthPasswordRecovery{}
	_ bin.BareDecoder = &AuthPasswordRecovery{}
)
