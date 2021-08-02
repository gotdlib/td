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

// MessagesSentEncryptedMessage represents TL type `messages.sentEncryptedMessage#560f8935`.
// Message without file attachemts sent to an encrypted file.
//
// See https://core.telegram.org/constructor/messages.sentEncryptedMessage for reference.
type MessagesSentEncryptedMessage struct {
	// Date of sending
	Date int
}

// MessagesSentEncryptedMessageTypeID is TL type id of MessagesSentEncryptedMessage.
const MessagesSentEncryptedMessageTypeID = 0x560f8935

func (s *MessagesSentEncryptedMessage) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSentEncryptedMessage) String() string {
	if s == nil {
		return "MessagesSentEncryptedMessage(nil)"
	}
	type Alias MessagesSentEncryptedMessage
	return fmt.Sprintf("MessagesSentEncryptedMessage%+v", Alias(*s))
}

// FillFrom fills MessagesSentEncryptedMessage from given interface.
func (s *MessagesSentEncryptedMessage) FillFrom(from interface {
	GetDate() (value int)
}) {
	s.Date = from.GetDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSentEncryptedMessage) TypeID() uint32 {
	return MessagesSentEncryptedMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSentEncryptedMessage) TypeName() string {
	return "messages.sentEncryptedMessage"
}

// TypeInfo returns info about TL type.
func (s *MessagesSentEncryptedMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.sentEncryptedMessage",
		ID:   MessagesSentEncryptedMessageTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSentEncryptedMessage) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode %s as nil", "messages.sentEncryptedMessage#560f8935")
	}
	b.PutID(MessagesSentEncryptedMessageTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSentEncryptedMessage) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode %s as nil", "messages.sentEncryptedMessage#560f8935")
	}
	b.PutInt(s.Date)
	return nil
}

// GetDate returns value of Date field.
func (s *MessagesSentEncryptedMessage) GetDate() (value int) {
	return s.Date
}

// Decode implements bin.Decoder.
func (s *MessagesSentEncryptedMessage) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode %s to nil", "messages.sentEncryptedMessage#560f8935")
	}
	if err := b.ConsumeID(MessagesSentEncryptedMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "messages.sentEncryptedMessage#560f8935", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSentEncryptedMessage) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode %s to nil", "messages.sentEncryptedMessage#560f8935")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.sentEncryptedMessage#560f8935", "date", err)
		}
		s.Date = value
	}
	return nil
}

// construct implements constructor of MessagesSentEncryptedMessageClass.
func (s MessagesSentEncryptedMessage) construct() MessagesSentEncryptedMessageClass { return &s }

// Ensuring interfaces in compile-time for MessagesSentEncryptedMessage.
var (
	_ bin.Encoder     = &MessagesSentEncryptedMessage{}
	_ bin.Decoder     = &MessagesSentEncryptedMessage{}
	_ bin.BareEncoder = &MessagesSentEncryptedMessage{}
	_ bin.BareDecoder = &MessagesSentEncryptedMessage{}

	_ MessagesSentEncryptedMessageClass = &MessagesSentEncryptedMessage{}
)

// MessagesSentEncryptedFile represents TL type `messages.sentEncryptedFile#9493ff32`.
// Message with a file enclosure sent to a protected chat
//
// See https://core.telegram.org/constructor/messages.sentEncryptedFile for reference.
type MessagesSentEncryptedFile struct {
	// Sending date
	Date int
	// Attached file
	File EncryptedFileClass
}

// MessagesSentEncryptedFileTypeID is TL type id of MessagesSentEncryptedFile.
const MessagesSentEncryptedFileTypeID = 0x9493ff32

func (s *MessagesSentEncryptedFile) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Date == 0) {
		return false
	}
	if !(s.File == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSentEncryptedFile) String() string {
	if s == nil {
		return "MessagesSentEncryptedFile(nil)"
	}
	type Alias MessagesSentEncryptedFile
	return fmt.Sprintf("MessagesSentEncryptedFile%+v", Alias(*s))
}

// FillFrom fills MessagesSentEncryptedFile from given interface.
func (s *MessagesSentEncryptedFile) FillFrom(from interface {
	GetDate() (value int)
	GetFile() (value EncryptedFileClass)
}) {
	s.Date = from.GetDate()
	s.File = from.GetFile()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSentEncryptedFile) TypeID() uint32 {
	return MessagesSentEncryptedFileTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSentEncryptedFile) TypeName() string {
	return "messages.sentEncryptedFile"
}

// TypeInfo returns info about TL type.
func (s *MessagesSentEncryptedFile) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.sentEncryptedFile",
		ID:   MessagesSentEncryptedFileTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "File",
			SchemaName: "file",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSentEncryptedFile) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode %s as nil", "messages.sentEncryptedFile#9493ff32")
	}
	b.PutID(MessagesSentEncryptedFileTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSentEncryptedFile) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode %s as nil", "messages.sentEncryptedFile#9493ff32")
	}
	b.PutInt(s.Date)
	if s.File == nil {
		return fmt.Errorf("unable to encode %s: field %s is nil", "messages.sentEncryptedFile#9493ff32", "file")
	}
	if err := s.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "messages.sentEncryptedFile#9493ff32", "file", err)
	}
	return nil
}

// GetDate returns value of Date field.
func (s *MessagesSentEncryptedFile) GetDate() (value int) {
	return s.Date
}

// GetFile returns value of File field.
func (s *MessagesSentEncryptedFile) GetFile() (value EncryptedFileClass) {
	return s.File
}

// Decode implements bin.Decoder.
func (s *MessagesSentEncryptedFile) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode %s to nil", "messages.sentEncryptedFile#9493ff32")
	}
	if err := b.ConsumeID(MessagesSentEncryptedFileTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "messages.sentEncryptedFile#9493ff32", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSentEncryptedFile) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode %s to nil", "messages.sentEncryptedFile#9493ff32")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.sentEncryptedFile#9493ff32", "date", err)
		}
		s.Date = value
	}
	{
		value, err := DecodeEncryptedFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.sentEncryptedFile#9493ff32", "file", err)
		}
		s.File = value
	}
	return nil
}

// construct implements constructor of MessagesSentEncryptedMessageClass.
func (s MessagesSentEncryptedFile) construct() MessagesSentEncryptedMessageClass { return &s }

// Ensuring interfaces in compile-time for MessagesSentEncryptedFile.
var (
	_ bin.Encoder     = &MessagesSentEncryptedFile{}
	_ bin.Decoder     = &MessagesSentEncryptedFile{}
	_ bin.BareEncoder = &MessagesSentEncryptedFile{}
	_ bin.BareDecoder = &MessagesSentEncryptedFile{}

	_ MessagesSentEncryptedMessageClass = &MessagesSentEncryptedFile{}
)

// MessagesSentEncryptedMessageClass represents messages.SentEncryptedMessage generic type.
//
// See https://core.telegram.org/type/messages.SentEncryptedMessage for reference.
//
// Example:
//  g, err := tg.DecodeMessagesSentEncryptedMessage(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.MessagesSentEncryptedMessage: // messages.sentEncryptedMessage#560f8935
//  case *tg.MessagesSentEncryptedFile: // messages.sentEncryptedFile#9493ff32
//  default: panic(v)
//  }
type MessagesSentEncryptedMessageClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessagesSentEncryptedMessageClass

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

	// Date of sending
	GetDate() (value int)
}

// DecodeMessagesSentEncryptedMessage implements binary de-serialization for MessagesSentEncryptedMessageClass.
func DecodeMessagesSentEncryptedMessage(buf *bin.Buffer) (MessagesSentEncryptedMessageClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesSentEncryptedMessageTypeID:
		// Decoding messages.sentEncryptedMessage#560f8935.
		v := MessagesSentEncryptedMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode %s: %w", "MessagesSentEncryptedMessageClass", err)
		}
		return &v, nil
	case MessagesSentEncryptedFileTypeID:
		// Decoding messages.sentEncryptedFile#9493ff32.
		v := MessagesSentEncryptedFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode %s: %w", "MessagesSentEncryptedMessageClass", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode %s: %w", "MessagesSentEncryptedMessageClass", bin.NewUnexpectedID(id))
	}
}

// MessagesSentEncryptedMessage boxes the MessagesSentEncryptedMessageClass providing a helper.
type MessagesSentEncryptedMessageBox struct {
	SentEncryptedMessage MessagesSentEncryptedMessageClass
}

// Decode implements bin.Decoder for MessagesSentEncryptedMessageBox.
func (b *MessagesSentEncryptedMessageBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode %sBox to nil", "MessagesSentEncryptedMessage")
	}
	v, err := DecodeMessagesSentEncryptedMessage(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SentEncryptedMessage = v
	return nil
}

// Encode implements bin.Encode for MessagesSentEncryptedMessageBox.
func (b *MessagesSentEncryptedMessageBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SentEncryptedMessage == nil {
		return fmt.Errorf("unable to encode %s as nil", "MessagesSentEncryptedMessageClass")
	}
	return b.SentEncryptedMessage.Encode(buf)
}

// MessagesSentEncryptedMessageClassArray is adapter for slice of MessagesSentEncryptedMessageClass.
type MessagesSentEncryptedMessageClassArray []MessagesSentEncryptedMessageClass

// Sort sorts slice of MessagesSentEncryptedMessageClass.
func (s MessagesSentEncryptedMessageClassArray) Sort(less func(a, b MessagesSentEncryptedMessageClass) bool) MessagesSentEncryptedMessageClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesSentEncryptedMessageClass.
func (s MessagesSentEncryptedMessageClassArray) SortStable(less func(a, b MessagesSentEncryptedMessageClass) bool) MessagesSentEncryptedMessageClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesSentEncryptedMessageClass.
func (s MessagesSentEncryptedMessageClassArray) Retain(keep func(x MessagesSentEncryptedMessageClass) bool) MessagesSentEncryptedMessageClassArray {
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
func (s MessagesSentEncryptedMessageClassArray) First() (v MessagesSentEncryptedMessageClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesSentEncryptedMessageClassArray) Last() (v MessagesSentEncryptedMessageClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesSentEncryptedMessageClassArray) PopFirst() (v MessagesSentEncryptedMessageClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesSentEncryptedMessageClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesSentEncryptedMessageClassArray) Pop() (v MessagesSentEncryptedMessageClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByDate sorts slice of MessagesSentEncryptedMessageClass by Date.
func (s MessagesSentEncryptedMessageClassArray) SortByDate() MessagesSentEncryptedMessageClassArray {
	return s.Sort(func(a, b MessagesSentEncryptedMessageClass) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of MessagesSentEncryptedMessageClass by Date.
func (s MessagesSentEncryptedMessageClassArray) SortStableByDate() MessagesSentEncryptedMessageClassArray {
	return s.SortStable(func(a, b MessagesSentEncryptedMessageClass) bool {
		return a.GetDate() < b.GetDate()
	})
}

// AsMessagesSentEncryptedMessage returns copy with only MessagesSentEncryptedMessage constructors.
func (s MessagesSentEncryptedMessageClassArray) AsMessagesSentEncryptedMessage() (to MessagesSentEncryptedMessageArray) {
	for _, elem := range s {
		value, ok := elem.(*MessagesSentEncryptedMessage)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMessagesSentEncryptedFile returns copy with only MessagesSentEncryptedFile constructors.
func (s MessagesSentEncryptedMessageClassArray) AsMessagesSentEncryptedFile() (to MessagesSentEncryptedFileArray) {
	for _, elem := range s {
		value, ok := elem.(*MessagesSentEncryptedFile)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// MessagesSentEncryptedMessageArray is adapter for slice of MessagesSentEncryptedMessage.
type MessagesSentEncryptedMessageArray []MessagesSentEncryptedMessage

// Sort sorts slice of MessagesSentEncryptedMessage.
func (s MessagesSentEncryptedMessageArray) Sort(less func(a, b MessagesSentEncryptedMessage) bool) MessagesSentEncryptedMessageArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesSentEncryptedMessage.
func (s MessagesSentEncryptedMessageArray) SortStable(less func(a, b MessagesSentEncryptedMessage) bool) MessagesSentEncryptedMessageArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesSentEncryptedMessage.
func (s MessagesSentEncryptedMessageArray) Retain(keep func(x MessagesSentEncryptedMessage) bool) MessagesSentEncryptedMessageArray {
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
func (s MessagesSentEncryptedMessageArray) First() (v MessagesSentEncryptedMessage, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesSentEncryptedMessageArray) Last() (v MessagesSentEncryptedMessage, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesSentEncryptedMessageArray) PopFirst() (v MessagesSentEncryptedMessage, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesSentEncryptedMessage
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesSentEncryptedMessageArray) Pop() (v MessagesSentEncryptedMessage, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByDate sorts slice of MessagesSentEncryptedMessage by Date.
func (s MessagesSentEncryptedMessageArray) SortByDate() MessagesSentEncryptedMessageArray {
	return s.Sort(func(a, b MessagesSentEncryptedMessage) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of MessagesSentEncryptedMessage by Date.
func (s MessagesSentEncryptedMessageArray) SortStableByDate() MessagesSentEncryptedMessageArray {
	return s.SortStable(func(a, b MessagesSentEncryptedMessage) bool {
		return a.GetDate() < b.GetDate()
	})
}

// MessagesSentEncryptedFileArray is adapter for slice of MessagesSentEncryptedFile.
type MessagesSentEncryptedFileArray []MessagesSentEncryptedFile

// Sort sorts slice of MessagesSentEncryptedFile.
func (s MessagesSentEncryptedFileArray) Sort(less func(a, b MessagesSentEncryptedFile) bool) MessagesSentEncryptedFileArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesSentEncryptedFile.
func (s MessagesSentEncryptedFileArray) SortStable(less func(a, b MessagesSentEncryptedFile) bool) MessagesSentEncryptedFileArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesSentEncryptedFile.
func (s MessagesSentEncryptedFileArray) Retain(keep func(x MessagesSentEncryptedFile) bool) MessagesSentEncryptedFileArray {
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
func (s MessagesSentEncryptedFileArray) First() (v MessagesSentEncryptedFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesSentEncryptedFileArray) Last() (v MessagesSentEncryptedFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesSentEncryptedFileArray) PopFirst() (v MessagesSentEncryptedFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesSentEncryptedFile
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesSentEncryptedFileArray) Pop() (v MessagesSentEncryptedFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByDate sorts slice of MessagesSentEncryptedFile by Date.
func (s MessagesSentEncryptedFileArray) SortByDate() MessagesSentEncryptedFileArray {
	return s.Sort(func(a, b MessagesSentEncryptedFile) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of MessagesSentEncryptedFile by Date.
func (s MessagesSentEncryptedFileArray) SortStableByDate() MessagesSentEncryptedFileArray {
	return s.SortStable(func(a, b MessagesSentEncryptedFile) bool {
		return a.GetDate() < b.GetDate()
	})
}
