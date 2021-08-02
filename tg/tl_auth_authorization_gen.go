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

// AuthAuthorization represents TL type `auth.authorization#cd050916`.
// Contains user authorization info.
//
// See https://core.telegram.org/constructor/auth.authorization for reference.
type AuthAuthorization struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Temporary passport¹ sessions
	//
	// Links:
	//  1) https://core.telegram.org/passport
	//
	// Use SetTmpSessions and GetTmpSessions helpers.
	TmpSessions int
	// Info on authorized user
	User UserClass
}

// AuthAuthorizationTypeID is TL type id of AuthAuthorization.
const AuthAuthorizationTypeID = 0xcd050916

func (a *AuthAuthorization) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Flags.Zero()) {
		return false
	}
	if !(a.TmpSessions == 0) {
		return false
	}
	if !(a.User == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AuthAuthorization) String() string {
	if a == nil {
		return "AuthAuthorization(nil)"
	}
	type Alias AuthAuthorization
	return fmt.Sprintf("AuthAuthorization%+v", Alias(*a))
}

// FillFrom fills AuthAuthorization from given interface.
func (a *AuthAuthorization) FillFrom(from interface {
	GetTmpSessions() (value int, ok bool)
	GetUser() (value UserClass)
}) {
	if val, ok := from.GetTmpSessions(); ok {
		a.TmpSessions = val
	}

	a.User = from.GetUser()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthAuthorization) TypeID() uint32 {
	return AuthAuthorizationTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthAuthorization) TypeName() string {
	return "auth.authorization"
}

// TypeInfo returns info about TL type.
func (a *AuthAuthorization) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.authorization",
		ID:   AuthAuthorizationTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TmpSessions",
			SchemaName: "tmp_sessions",
			Null:       !a.Flags.Has(0),
		},
		{
			Name:       "User",
			SchemaName: "user",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AuthAuthorization) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode %s as nil", "auth.authorization#cd050916")
	}
	b.PutID(AuthAuthorizationTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AuthAuthorization) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode %s as nil", "auth.authorization#cd050916")
	}
	if !(a.TmpSessions == 0) {
		a.Flags.Set(0)
	}
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "auth.authorization#cd050916", "flags", err)
	}
	if a.Flags.Has(0) {
		b.PutInt(a.TmpSessions)
	}
	if a.User == nil {
		return fmt.Errorf("unable to encode %s: field %s is nil", "auth.authorization#cd050916", "user")
	}
	if err := a.User.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "auth.authorization#cd050916", "user", err)
	}
	return nil
}

// SetTmpSessions sets value of TmpSessions conditional field.
func (a *AuthAuthorization) SetTmpSessions(value int) {
	a.Flags.Set(0)
	a.TmpSessions = value
}

// GetTmpSessions returns value of TmpSessions conditional field and
// boolean which is true if field was set.
func (a *AuthAuthorization) GetTmpSessions() (value int, ok bool) {
	if !a.Flags.Has(0) {
		return value, false
	}
	return a.TmpSessions, true
}

// GetUser returns value of User field.
func (a *AuthAuthorization) GetUser() (value UserClass) {
	return a.User
}

// Decode implements bin.Decoder.
func (a *AuthAuthorization) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode %s to nil", "auth.authorization#cd050916")
	}
	if err := b.ConsumeID(AuthAuthorizationTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "auth.authorization#cd050916", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AuthAuthorization) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode %s to nil", "auth.authorization#cd050916")
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "auth.authorization#cd050916", "flags", err)
		}
	}
	if a.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "auth.authorization#cd050916", "tmp_sessions", err)
		}
		a.TmpSessions = value
	}
	{
		value, err := DecodeUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "auth.authorization#cd050916", "user", err)
		}
		a.User = value
	}
	return nil
}

// construct implements constructor of AuthAuthorizationClass.
func (a AuthAuthorization) construct() AuthAuthorizationClass { return &a }

// Ensuring interfaces in compile-time for AuthAuthorization.
var (
	_ bin.Encoder     = &AuthAuthorization{}
	_ bin.Decoder     = &AuthAuthorization{}
	_ bin.BareEncoder = &AuthAuthorization{}
	_ bin.BareDecoder = &AuthAuthorization{}

	_ AuthAuthorizationClass = &AuthAuthorization{}
)

// AuthAuthorizationSignUpRequired represents TL type `auth.authorizationSignUpRequired#44747e9a`.
// An account with this phone number doesn't exist on telegram: the user has to enter
// basic information and sign up¹
//
// Links:
//  1) https://core.telegram.org/api/auth
//
// See https://core.telegram.org/constructor/auth.authorizationSignUpRequired for reference.
type AuthAuthorizationSignUpRequired struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Telegram's terms of service: the user must read and accept the terms of service before
	// signing up to telegram
	//
	// Use SetTermsOfService and GetTermsOfService helpers.
	TermsOfService HelpTermsOfService
}

// AuthAuthorizationSignUpRequiredTypeID is TL type id of AuthAuthorizationSignUpRequired.
const AuthAuthorizationSignUpRequiredTypeID = 0x44747e9a

func (a *AuthAuthorizationSignUpRequired) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Flags.Zero()) {
		return false
	}
	if !(a.TermsOfService.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AuthAuthorizationSignUpRequired) String() string {
	if a == nil {
		return "AuthAuthorizationSignUpRequired(nil)"
	}
	type Alias AuthAuthorizationSignUpRequired
	return fmt.Sprintf("AuthAuthorizationSignUpRequired%+v", Alias(*a))
}

// FillFrom fills AuthAuthorizationSignUpRequired from given interface.
func (a *AuthAuthorizationSignUpRequired) FillFrom(from interface {
	GetTermsOfService() (value HelpTermsOfService, ok bool)
}) {
	if val, ok := from.GetTermsOfService(); ok {
		a.TermsOfService = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthAuthorizationSignUpRequired) TypeID() uint32 {
	return AuthAuthorizationSignUpRequiredTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthAuthorizationSignUpRequired) TypeName() string {
	return "auth.authorizationSignUpRequired"
}

// TypeInfo returns info about TL type.
func (a *AuthAuthorizationSignUpRequired) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.authorizationSignUpRequired",
		ID:   AuthAuthorizationSignUpRequiredTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TermsOfService",
			SchemaName: "terms_of_service",
			Null:       !a.Flags.Has(0),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AuthAuthorizationSignUpRequired) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode %s as nil", "auth.authorizationSignUpRequired#44747e9a")
	}
	b.PutID(AuthAuthorizationSignUpRequiredTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AuthAuthorizationSignUpRequired) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode %s as nil", "auth.authorizationSignUpRequired#44747e9a")
	}
	if !(a.TermsOfService.Zero()) {
		a.Flags.Set(0)
	}
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "auth.authorizationSignUpRequired#44747e9a", "flags", err)
	}
	if a.Flags.Has(0) {
		if err := a.TermsOfService.Encode(b); err != nil {
			return fmt.Errorf("unable to encode %s: field %s: %w", "auth.authorizationSignUpRequired#44747e9a", "terms_of_service", err)
		}
	}
	return nil
}

// SetTermsOfService sets value of TermsOfService conditional field.
func (a *AuthAuthorizationSignUpRequired) SetTermsOfService(value HelpTermsOfService) {
	a.Flags.Set(0)
	a.TermsOfService = value
}

// GetTermsOfService returns value of TermsOfService conditional field and
// boolean which is true if field was set.
func (a *AuthAuthorizationSignUpRequired) GetTermsOfService() (value HelpTermsOfService, ok bool) {
	if !a.Flags.Has(0) {
		return value, false
	}
	return a.TermsOfService, true
}

// Decode implements bin.Decoder.
func (a *AuthAuthorizationSignUpRequired) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode %s to nil", "auth.authorizationSignUpRequired#44747e9a")
	}
	if err := b.ConsumeID(AuthAuthorizationSignUpRequiredTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "auth.authorizationSignUpRequired#44747e9a", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AuthAuthorizationSignUpRequired) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode %s to nil", "auth.authorizationSignUpRequired#44747e9a")
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "auth.authorizationSignUpRequired#44747e9a", "flags", err)
		}
	}
	if a.Flags.Has(0) {
		if err := a.TermsOfService.Decode(b); err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "auth.authorizationSignUpRequired#44747e9a", "terms_of_service", err)
		}
	}
	return nil
}

// construct implements constructor of AuthAuthorizationClass.
func (a AuthAuthorizationSignUpRequired) construct() AuthAuthorizationClass { return &a }

// Ensuring interfaces in compile-time for AuthAuthorizationSignUpRequired.
var (
	_ bin.Encoder     = &AuthAuthorizationSignUpRequired{}
	_ bin.Decoder     = &AuthAuthorizationSignUpRequired{}
	_ bin.BareEncoder = &AuthAuthorizationSignUpRequired{}
	_ bin.BareDecoder = &AuthAuthorizationSignUpRequired{}

	_ AuthAuthorizationClass = &AuthAuthorizationSignUpRequired{}
)

// AuthAuthorizationClass represents auth.Authorization generic type.
//
// See https://core.telegram.org/type/auth.Authorization for reference.
//
// Example:
//  g, err := tg.DecodeAuthAuthorization(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.AuthAuthorization: // auth.authorization#cd050916
//  case *tg.AuthAuthorizationSignUpRequired: // auth.authorizationSignUpRequired#44747e9a
//  default: panic(v)
//  }
type AuthAuthorizationClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() AuthAuthorizationClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeAuthAuthorization implements binary de-serialization for AuthAuthorizationClass.
func DecodeAuthAuthorization(buf *bin.Buffer) (AuthAuthorizationClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AuthAuthorizationTypeID:
		// Decoding auth.authorization#cd050916.
		v := AuthAuthorization{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode %s: %w", "AuthAuthorizationClass", err)
		}
		return &v, nil
	case AuthAuthorizationSignUpRequiredTypeID:
		// Decoding auth.authorizationSignUpRequired#44747e9a.
		v := AuthAuthorizationSignUpRequired{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode %s: %w", "AuthAuthorizationClass", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode %s: %w", "AuthAuthorizationClass", bin.NewUnexpectedID(id))
	}
}

// AuthAuthorization boxes the AuthAuthorizationClass providing a helper.
type AuthAuthorizationBox struct {
	Authorization AuthAuthorizationClass
}

// Decode implements bin.Decoder for AuthAuthorizationBox.
func (b *AuthAuthorizationBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode %sBox to nil", "AuthAuthorization")
	}
	v, err := DecodeAuthAuthorization(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Authorization = v
	return nil
}

// Encode implements bin.Encode for AuthAuthorizationBox.
func (b *AuthAuthorizationBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Authorization == nil {
		return fmt.Errorf("unable to encode %s as nil", "AuthAuthorizationClass")
	}
	return b.Authorization.Encode(buf)
}

// AuthAuthorizationClassArray is adapter for slice of AuthAuthorizationClass.
type AuthAuthorizationClassArray []AuthAuthorizationClass

// Sort sorts slice of AuthAuthorizationClass.
func (s AuthAuthorizationClassArray) Sort(less func(a, b AuthAuthorizationClass) bool) AuthAuthorizationClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AuthAuthorizationClass.
func (s AuthAuthorizationClassArray) SortStable(less func(a, b AuthAuthorizationClass) bool) AuthAuthorizationClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AuthAuthorizationClass.
func (s AuthAuthorizationClassArray) Retain(keep func(x AuthAuthorizationClass) bool) AuthAuthorizationClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s AuthAuthorizationClassArray) First() (v AuthAuthorizationClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AuthAuthorizationClassArray) Last() (v AuthAuthorizationClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AuthAuthorizationClassArray) PopFirst() (v AuthAuthorizationClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AuthAuthorizationClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AuthAuthorizationClassArray) Pop() (v AuthAuthorizationClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsAuthAuthorization returns copy with only AuthAuthorization constructors.
func (s AuthAuthorizationClassArray) AsAuthAuthorization() (to AuthAuthorizationArray) {
	for _, elem := range s {
		value, ok := elem.(*AuthAuthorization)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsAuthAuthorizationSignUpRequired returns copy with only AuthAuthorizationSignUpRequired constructors.
func (s AuthAuthorizationClassArray) AsAuthAuthorizationSignUpRequired() (to AuthAuthorizationSignUpRequiredArray) {
	for _, elem := range s {
		value, ok := elem.(*AuthAuthorizationSignUpRequired)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AuthAuthorizationArray is adapter for slice of AuthAuthorization.
type AuthAuthorizationArray []AuthAuthorization

// Sort sorts slice of AuthAuthorization.
func (s AuthAuthorizationArray) Sort(less func(a, b AuthAuthorization) bool) AuthAuthorizationArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AuthAuthorization.
func (s AuthAuthorizationArray) SortStable(less func(a, b AuthAuthorization) bool) AuthAuthorizationArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AuthAuthorization.
func (s AuthAuthorizationArray) Retain(keep func(x AuthAuthorization) bool) AuthAuthorizationArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s AuthAuthorizationArray) First() (v AuthAuthorization, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AuthAuthorizationArray) Last() (v AuthAuthorization, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AuthAuthorizationArray) PopFirst() (v AuthAuthorization, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AuthAuthorization
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AuthAuthorizationArray) Pop() (v AuthAuthorization, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AuthAuthorizationSignUpRequiredArray is adapter for slice of AuthAuthorizationSignUpRequired.
type AuthAuthorizationSignUpRequiredArray []AuthAuthorizationSignUpRequired

// Sort sorts slice of AuthAuthorizationSignUpRequired.
func (s AuthAuthorizationSignUpRequiredArray) Sort(less func(a, b AuthAuthorizationSignUpRequired) bool) AuthAuthorizationSignUpRequiredArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AuthAuthorizationSignUpRequired.
func (s AuthAuthorizationSignUpRequiredArray) SortStable(less func(a, b AuthAuthorizationSignUpRequired) bool) AuthAuthorizationSignUpRequiredArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AuthAuthorizationSignUpRequired.
func (s AuthAuthorizationSignUpRequiredArray) Retain(keep func(x AuthAuthorizationSignUpRequired) bool) AuthAuthorizationSignUpRequiredArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s AuthAuthorizationSignUpRequiredArray) First() (v AuthAuthorizationSignUpRequired, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AuthAuthorizationSignUpRequiredArray) Last() (v AuthAuthorizationSignUpRequired, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AuthAuthorizationSignUpRequiredArray) PopFirst() (v AuthAuthorizationSignUpRequired, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AuthAuthorizationSignUpRequired
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AuthAuthorizationSignUpRequiredArray) Pop() (v AuthAuthorizationSignUpRequired, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
