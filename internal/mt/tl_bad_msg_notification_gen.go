// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// BadMsgNotification represents TL type `bad_msg_notification#a7eff811`.
type BadMsgNotification struct {
	// BadMsgID field of BadMsgNotification.
	BadMsgID int64
	// BadMsgSeqno field of BadMsgNotification.
	BadMsgSeqno int
	// ErrorCode field of BadMsgNotification.
	ErrorCode int
}

// BadMsgNotificationTypeID is TL type id of BadMsgNotification.
const BadMsgNotificationTypeID = 0xa7eff811

func (b *BadMsgNotification) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.BadMsgID == 0) {
		return false
	}
	if !(b.BadMsgSeqno == 0) {
		return false
	}
	if !(b.ErrorCode == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BadMsgNotification) String() string {
	if b == nil {
		return "BadMsgNotification(nil)"
	}
	type Alias BadMsgNotification
	return fmt.Sprintf("BadMsgNotification%+v", Alias(*b))
}

// FillFrom fills BadMsgNotification from given interface.
func (b *BadMsgNotification) FillFrom(from interface {
	GetBadMsgID() (value int64)
	GetBadMsgSeqno() (value int)
	GetErrorCode() (value int)
}) {
	b.BadMsgID = from.GetBadMsgID()
	b.BadMsgSeqno = from.GetBadMsgSeqno()
	b.ErrorCode = from.GetErrorCode()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BadMsgNotification) TypeID() uint32 {
	return BadMsgNotificationTypeID
}

// TypeName returns name of type in TL schema.
func (*BadMsgNotification) TypeName() string {
	return "bad_msg_notification"
}

// TypeInfo returns info about TL type.
func (b *BadMsgNotification) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bad_msg_notification",
		ID:   BadMsgNotificationTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BadMsgID",
			SchemaName: "bad_msg_id",
		},
		{
			Name:       "BadMsgSeqno",
			SchemaName: "bad_msg_seqno",
		},
		{
			Name:       "ErrorCode",
			SchemaName: "error_code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BadMsgNotification) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode %s as nil", "bad_msg_notification#a7eff811")
	}
	buf.PutID(BadMsgNotificationTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BadMsgNotification) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode %s as nil", "bad_msg_notification#a7eff811")
	}
	buf.PutLong(b.BadMsgID)
	buf.PutInt(b.BadMsgSeqno)
	buf.PutInt(b.ErrorCode)
	return nil
}

// GetBadMsgID returns value of BadMsgID field.
func (b *BadMsgNotification) GetBadMsgID() (value int64) {
	return b.BadMsgID
}

// GetBadMsgSeqno returns value of BadMsgSeqno field.
func (b *BadMsgNotification) GetBadMsgSeqno() (value int) {
	return b.BadMsgSeqno
}

// GetErrorCode returns value of ErrorCode field.
func (b *BadMsgNotification) GetErrorCode() (value int) {
	return b.ErrorCode
}

// Decode implements bin.Decoder.
func (b *BadMsgNotification) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode %s to nil", "bad_msg_notification#a7eff811")
	}
	if err := buf.ConsumeID(BadMsgNotificationTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "bad_msg_notification#a7eff811", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BadMsgNotification) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode %s to nil", "bad_msg_notification#a7eff811")
	}
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "bad_msg_notification#a7eff811", "bad_msg_id", err)
		}
		b.BadMsgID = value
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "bad_msg_notification#a7eff811", "bad_msg_seqno", err)
		}
		b.BadMsgSeqno = value
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "bad_msg_notification#a7eff811", "error_code", err)
		}
		b.ErrorCode = value
	}
	return nil
}

// construct implements constructor of BadMsgNotificationClass.
func (b BadMsgNotification) construct() BadMsgNotificationClass { return &b }

// Ensuring interfaces in compile-time for BadMsgNotification.
var (
	_ bin.Encoder     = &BadMsgNotification{}
	_ bin.Decoder     = &BadMsgNotification{}
	_ bin.BareEncoder = &BadMsgNotification{}
	_ bin.BareDecoder = &BadMsgNotification{}

	_ BadMsgNotificationClass = &BadMsgNotification{}
)

// BadServerSalt represents TL type `bad_server_salt#edab447b`.
type BadServerSalt struct {
	// BadMsgID field of BadServerSalt.
	BadMsgID int64
	// BadMsgSeqno field of BadServerSalt.
	BadMsgSeqno int
	// ErrorCode field of BadServerSalt.
	ErrorCode int
	// NewServerSalt field of BadServerSalt.
	NewServerSalt int64
}

// BadServerSaltTypeID is TL type id of BadServerSalt.
const BadServerSaltTypeID = 0xedab447b

func (b *BadServerSalt) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.BadMsgID == 0) {
		return false
	}
	if !(b.BadMsgSeqno == 0) {
		return false
	}
	if !(b.ErrorCode == 0) {
		return false
	}
	if !(b.NewServerSalt == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BadServerSalt) String() string {
	if b == nil {
		return "BadServerSalt(nil)"
	}
	type Alias BadServerSalt
	return fmt.Sprintf("BadServerSalt%+v", Alias(*b))
}

// FillFrom fills BadServerSalt from given interface.
func (b *BadServerSalt) FillFrom(from interface {
	GetBadMsgID() (value int64)
	GetBadMsgSeqno() (value int)
	GetErrorCode() (value int)
	GetNewServerSalt() (value int64)
}) {
	b.BadMsgID = from.GetBadMsgID()
	b.BadMsgSeqno = from.GetBadMsgSeqno()
	b.ErrorCode = from.GetErrorCode()
	b.NewServerSalt = from.GetNewServerSalt()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BadServerSalt) TypeID() uint32 {
	return BadServerSaltTypeID
}

// TypeName returns name of type in TL schema.
func (*BadServerSalt) TypeName() string {
	return "bad_server_salt"
}

// TypeInfo returns info about TL type.
func (b *BadServerSalt) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bad_server_salt",
		ID:   BadServerSaltTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BadMsgID",
			SchemaName: "bad_msg_id",
		},
		{
			Name:       "BadMsgSeqno",
			SchemaName: "bad_msg_seqno",
		},
		{
			Name:       "ErrorCode",
			SchemaName: "error_code",
		},
		{
			Name:       "NewServerSalt",
			SchemaName: "new_server_salt",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BadServerSalt) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode %s as nil", "bad_server_salt#edab447b")
	}
	buf.PutID(BadServerSaltTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BadServerSalt) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode %s as nil", "bad_server_salt#edab447b")
	}
	buf.PutLong(b.BadMsgID)
	buf.PutInt(b.BadMsgSeqno)
	buf.PutInt(b.ErrorCode)
	buf.PutLong(b.NewServerSalt)
	return nil
}

// GetBadMsgID returns value of BadMsgID field.
func (b *BadServerSalt) GetBadMsgID() (value int64) {
	return b.BadMsgID
}

// GetBadMsgSeqno returns value of BadMsgSeqno field.
func (b *BadServerSalt) GetBadMsgSeqno() (value int) {
	return b.BadMsgSeqno
}

// GetErrorCode returns value of ErrorCode field.
func (b *BadServerSalt) GetErrorCode() (value int) {
	return b.ErrorCode
}

// GetNewServerSalt returns value of NewServerSalt field.
func (b *BadServerSalt) GetNewServerSalt() (value int64) {
	return b.NewServerSalt
}

// Decode implements bin.Decoder.
func (b *BadServerSalt) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode %s to nil", "bad_server_salt#edab447b")
	}
	if err := buf.ConsumeID(BadServerSaltTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "bad_server_salt#edab447b", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BadServerSalt) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode %s to nil", "bad_server_salt#edab447b")
	}
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "bad_server_salt#edab447b", "bad_msg_id", err)
		}
		b.BadMsgID = value
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "bad_server_salt#edab447b", "bad_msg_seqno", err)
		}
		b.BadMsgSeqno = value
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "bad_server_salt#edab447b", "error_code", err)
		}
		b.ErrorCode = value
	}
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "bad_server_salt#edab447b", "new_server_salt", err)
		}
		b.NewServerSalt = value
	}
	return nil
}

// construct implements constructor of BadMsgNotificationClass.
func (b BadServerSalt) construct() BadMsgNotificationClass { return &b }

// Ensuring interfaces in compile-time for BadServerSalt.
var (
	_ bin.Encoder     = &BadServerSalt{}
	_ bin.Decoder     = &BadServerSalt{}
	_ bin.BareEncoder = &BadServerSalt{}
	_ bin.BareDecoder = &BadServerSalt{}

	_ BadMsgNotificationClass = &BadServerSalt{}
)

// BadMsgNotificationClass represents BadMsgNotification generic type.
//
// Example:
//  g, err := mt.DecodeBadMsgNotification(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *mt.BadMsgNotification: // bad_msg_notification#a7eff811
//  case *mt.BadServerSalt: // bad_server_salt#edab447b
//  default: panic(v)
//  }
type BadMsgNotificationClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() BadMsgNotificationClass

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

	// BadMsgID field of BadMsgNotification.
	GetBadMsgID() (value int64)

	// BadMsgSeqno field of BadMsgNotification.
	GetBadMsgSeqno() (value int)

	// ErrorCode field of BadMsgNotification.
	GetErrorCode() (value int)
}

// DecodeBadMsgNotification implements binary de-serialization for BadMsgNotificationClass.
func DecodeBadMsgNotification(buf *bin.Buffer) (BadMsgNotificationClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case BadMsgNotificationTypeID:
		// Decoding bad_msg_notification#a7eff811.
		v := BadMsgNotification{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode %s: %w", "BadMsgNotificationClass", err)
		}
		return &v, nil
	case BadServerSaltTypeID:
		// Decoding bad_server_salt#edab447b.
		v := BadServerSalt{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode %s: %w", "BadMsgNotificationClass", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode %s: %w", "BadMsgNotificationClass", bin.NewUnexpectedID(id))
	}
}

// BadMsgNotification boxes the BadMsgNotificationClass providing a helper.
type BadMsgNotificationBox struct {
	BadMsgNotification BadMsgNotificationClass
}

// Decode implements bin.Decoder for BadMsgNotificationBox.
func (b *BadMsgNotificationBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode %sBox to nil", "BadMsgNotification")
	}
	v, err := DecodeBadMsgNotification(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.BadMsgNotification = v
	return nil
}

// Encode implements bin.Encode for BadMsgNotificationBox.
func (b *BadMsgNotificationBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.BadMsgNotification == nil {
		return fmt.Errorf("unable to encode %s as nil", "BadMsgNotificationClass")
	}
	return b.BadMsgNotification.Encode(buf)
}

// BadMsgNotificationClassArray is adapter for slice of BadMsgNotificationClass.
type BadMsgNotificationClassArray []BadMsgNotificationClass

// Sort sorts slice of BadMsgNotificationClass.
func (s BadMsgNotificationClassArray) Sort(less func(a, b BadMsgNotificationClass) bool) BadMsgNotificationClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of BadMsgNotificationClass.
func (s BadMsgNotificationClassArray) SortStable(less func(a, b BadMsgNotificationClass) bool) BadMsgNotificationClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of BadMsgNotificationClass.
func (s BadMsgNotificationClassArray) Retain(keep func(x BadMsgNotificationClass) bool) BadMsgNotificationClassArray {
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
func (s BadMsgNotificationClassArray) First() (v BadMsgNotificationClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s BadMsgNotificationClassArray) Last() (v BadMsgNotificationClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *BadMsgNotificationClassArray) PopFirst() (v BadMsgNotificationClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero BadMsgNotificationClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *BadMsgNotificationClassArray) Pop() (v BadMsgNotificationClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsBadMsgNotification returns copy with only BadMsgNotification constructors.
func (s BadMsgNotificationClassArray) AsBadMsgNotification() (to BadMsgNotificationArray) {
	for _, elem := range s {
		value, ok := elem.(*BadMsgNotification)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsBadServerSalt returns copy with only BadServerSalt constructors.
func (s BadMsgNotificationClassArray) AsBadServerSalt() (to BadServerSaltArray) {
	for _, elem := range s {
		value, ok := elem.(*BadServerSalt)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// BadMsgNotificationArray is adapter for slice of BadMsgNotification.
type BadMsgNotificationArray []BadMsgNotification

// Sort sorts slice of BadMsgNotification.
func (s BadMsgNotificationArray) Sort(less func(a, b BadMsgNotification) bool) BadMsgNotificationArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of BadMsgNotification.
func (s BadMsgNotificationArray) SortStable(less func(a, b BadMsgNotification) bool) BadMsgNotificationArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of BadMsgNotification.
func (s BadMsgNotificationArray) Retain(keep func(x BadMsgNotification) bool) BadMsgNotificationArray {
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
func (s BadMsgNotificationArray) First() (v BadMsgNotification, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s BadMsgNotificationArray) Last() (v BadMsgNotification, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *BadMsgNotificationArray) PopFirst() (v BadMsgNotification, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero BadMsgNotification
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *BadMsgNotificationArray) Pop() (v BadMsgNotification, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// BadServerSaltArray is adapter for slice of BadServerSalt.
type BadServerSaltArray []BadServerSalt

// Sort sorts slice of BadServerSalt.
func (s BadServerSaltArray) Sort(less func(a, b BadServerSalt) bool) BadServerSaltArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of BadServerSalt.
func (s BadServerSaltArray) SortStable(less func(a, b BadServerSalt) bool) BadServerSaltArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of BadServerSalt.
func (s BadServerSaltArray) Retain(keep func(x BadServerSalt) bool) BadServerSaltArray {
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
func (s BadServerSaltArray) First() (v BadServerSalt, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s BadServerSaltArray) Last() (v BadServerSalt, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *BadServerSaltArray) PopFirst() (v BadServerSalt, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero BadServerSalt
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *BadServerSaltArray) Pop() (v BadServerSalt, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
