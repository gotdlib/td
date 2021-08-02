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

// DraftMessageEmpty represents TL type `draftMessageEmpty#1b0c841a`.
// Empty draft
//
// See https://core.telegram.org/constructor/draftMessageEmpty for reference.
type DraftMessageEmpty struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// When was the draft last updated
	//
	// Use SetDate and GetDate helpers.
	Date int
}

// DraftMessageEmptyTypeID is TL type id of DraftMessageEmpty.
const DraftMessageEmptyTypeID = 0x1b0c841a

func (d *DraftMessageEmpty) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Flags.Zero()) {
		return false
	}
	if !(d.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DraftMessageEmpty) String() string {
	if d == nil {
		return "DraftMessageEmpty(nil)"
	}
	type Alias DraftMessageEmpty
	return fmt.Sprintf("DraftMessageEmpty%+v", Alias(*d))
}

// FillFrom fills DraftMessageEmpty from given interface.
func (d *DraftMessageEmpty) FillFrom(from interface {
	GetDate() (value int, ok bool)
}) {
	if val, ok := from.GetDate(); ok {
		d.Date = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DraftMessageEmpty) TypeID() uint32 {
	return DraftMessageEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*DraftMessageEmpty) TypeName() string {
	return "draftMessageEmpty"
}

// TypeInfo returns info about TL type.
func (d *DraftMessageEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "draftMessageEmpty",
		ID:   DraftMessageEmptyTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Date",
			SchemaName: "date",
			Null:       !d.Flags.Has(0),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DraftMessageEmpty) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode %s as nil", "draftMessageEmpty#1b0c841a")
	}
	b.PutID(DraftMessageEmptyTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DraftMessageEmpty) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode %s as nil", "draftMessageEmpty#1b0c841a")
	}
	if !(d.Date == 0) {
		d.Flags.Set(0)
	}
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "draftMessageEmpty#1b0c841a", "flags", err)
	}
	if d.Flags.Has(0) {
		b.PutInt(d.Date)
	}
	return nil
}

// SetDate sets value of Date conditional field.
func (d *DraftMessageEmpty) SetDate(value int) {
	d.Flags.Set(0)
	d.Date = value
}

// GetDate returns value of Date conditional field and
// boolean which is true if field was set.
func (d *DraftMessageEmpty) GetDate() (value int, ok bool) {
	if !d.Flags.Has(0) {
		return value, false
	}
	return d.Date, true
}

// Decode implements bin.Decoder.
func (d *DraftMessageEmpty) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode %s to nil", "draftMessageEmpty#1b0c841a")
	}
	if err := b.ConsumeID(DraftMessageEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "draftMessageEmpty#1b0c841a", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DraftMessageEmpty) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode %s to nil", "draftMessageEmpty#1b0c841a")
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "draftMessageEmpty#1b0c841a", "flags", err)
		}
	}
	if d.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "draftMessageEmpty#1b0c841a", "date", err)
		}
		d.Date = value
	}
	return nil
}

// construct implements constructor of DraftMessageClass.
func (d DraftMessageEmpty) construct() DraftMessageClass { return &d }

// Ensuring interfaces in compile-time for DraftMessageEmpty.
var (
	_ bin.Encoder     = &DraftMessageEmpty{}
	_ bin.Decoder     = &DraftMessageEmpty{}
	_ bin.BareEncoder = &DraftMessageEmpty{}
	_ bin.BareDecoder = &DraftMessageEmpty{}

	_ DraftMessageClass = &DraftMessageEmpty{}
)

// DraftMessage represents TL type `draftMessage#fd8e711f`.
// Represents a message draft¹.
//
// Links:
//  1) https://core.telegram.org/api/drafts
//
// See https://core.telegram.org/constructor/draftMessage for reference.
type DraftMessage struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether no webpage preview will be generated
	NoWebpage bool
	// The message this message will reply to
	//
	// Use SetReplyToMsgID and GetReplyToMsgID helpers.
	ReplyToMsgID int
	// The draft
	Message string
	// Message entities¹ for styled text.
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
	// Date of last update of the draft.
	Date int
}

// DraftMessageTypeID is TL type id of DraftMessage.
const DraftMessageTypeID = 0xfd8e711f

func (d *DraftMessage) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Flags.Zero()) {
		return false
	}
	if !(d.NoWebpage == false) {
		return false
	}
	if !(d.ReplyToMsgID == 0) {
		return false
	}
	if !(d.Message == "") {
		return false
	}
	if !(d.Entities == nil) {
		return false
	}
	if !(d.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DraftMessage) String() string {
	if d == nil {
		return "DraftMessage(nil)"
	}
	type Alias DraftMessage
	return fmt.Sprintf("DraftMessage%+v", Alias(*d))
}

// FillFrom fills DraftMessage from given interface.
func (d *DraftMessage) FillFrom(from interface {
	GetNoWebpage() (value bool)
	GetReplyToMsgID() (value int, ok bool)
	GetMessage() (value string)
	GetEntities() (value []MessageEntityClass, ok bool)
	GetDate() (value int)
}) {
	d.NoWebpage = from.GetNoWebpage()
	if val, ok := from.GetReplyToMsgID(); ok {
		d.ReplyToMsgID = val
	}

	d.Message = from.GetMessage()
	if val, ok := from.GetEntities(); ok {
		d.Entities = val
	}

	d.Date = from.GetDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DraftMessage) TypeID() uint32 {
	return DraftMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*DraftMessage) TypeName() string {
	return "draftMessage"
}

// TypeInfo returns info about TL type.
func (d *DraftMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "draftMessage",
		ID:   DraftMessageTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "NoWebpage",
			SchemaName: "no_webpage",
			Null:       !d.Flags.Has(1),
		},
		{
			Name:       "ReplyToMsgID",
			SchemaName: "reply_to_msg_id",
			Null:       !d.Flags.Has(0),
		},
		{
			Name:       "Message",
			SchemaName: "message",
		},
		{
			Name:       "Entities",
			SchemaName: "entities",
			Null:       !d.Flags.Has(3),
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DraftMessage) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode %s as nil", "draftMessage#fd8e711f")
	}
	b.PutID(DraftMessageTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DraftMessage) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode %s as nil", "draftMessage#fd8e711f")
	}
	if !(d.NoWebpage == false) {
		d.Flags.Set(1)
	}
	if !(d.ReplyToMsgID == 0) {
		d.Flags.Set(0)
	}
	if !(d.Entities == nil) {
		d.Flags.Set(3)
	}
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "draftMessage#fd8e711f", "flags", err)
	}
	if d.Flags.Has(0) {
		b.PutInt(d.ReplyToMsgID)
	}
	b.PutString(d.Message)
	if d.Flags.Has(3) {
		b.PutVectorHeader(len(d.Entities))
		for idx, v := range d.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode %s: field %s element with index %d is nil", "draftMessage#fd8e711f", "entities", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode %s: field %s element with index %d: %w", "draftMessage#fd8e711f", "entities", idx, err)
			}
		}
	}
	b.PutInt(d.Date)
	return nil
}

// SetNoWebpage sets value of NoWebpage conditional field.
func (d *DraftMessage) SetNoWebpage(value bool) {
	if value {
		d.Flags.Set(1)
		d.NoWebpage = true
	} else {
		d.Flags.Unset(1)
		d.NoWebpage = false
	}
}

// GetNoWebpage returns value of NoWebpage conditional field.
func (d *DraftMessage) GetNoWebpage() (value bool) {
	return d.Flags.Has(1)
}

// SetReplyToMsgID sets value of ReplyToMsgID conditional field.
func (d *DraftMessage) SetReplyToMsgID(value int) {
	d.Flags.Set(0)
	d.ReplyToMsgID = value
}

// GetReplyToMsgID returns value of ReplyToMsgID conditional field and
// boolean which is true if field was set.
func (d *DraftMessage) GetReplyToMsgID() (value int, ok bool) {
	if !d.Flags.Has(0) {
		return value, false
	}
	return d.ReplyToMsgID, true
}

// GetMessage returns value of Message field.
func (d *DraftMessage) GetMessage() (value string) {
	return d.Message
}

// SetEntities sets value of Entities conditional field.
func (d *DraftMessage) SetEntities(value []MessageEntityClass) {
	d.Flags.Set(3)
	d.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (d *DraftMessage) GetEntities() (value []MessageEntityClass, ok bool) {
	if !d.Flags.Has(3) {
		return value, false
	}
	return d.Entities, true
}

// MapEntities returns field Entities wrapped in MessageEntityClassArray helper.
func (d *DraftMessage) MapEntities() (value MessageEntityClassArray, ok bool) {
	if !d.Flags.Has(3) {
		return value, false
	}
	return MessageEntityClassArray(d.Entities), true
}

// GetDate returns value of Date field.
func (d *DraftMessage) GetDate() (value int) {
	return d.Date
}

// Decode implements bin.Decoder.
func (d *DraftMessage) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode %s to nil", "draftMessage#fd8e711f")
	}
	if err := b.ConsumeID(DraftMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "draftMessage#fd8e711f", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DraftMessage) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode %s to nil", "draftMessage#fd8e711f")
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "draftMessage#fd8e711f", "flags", err)
		}
	}
	d.NoWebpage = d.Flags.Has(1)
	if d.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "draftMessage#fd8e711f", "reply_to_msg_id", err)
		}
		d.ReplyToMsgID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "draftMessage#fd8e711f", "message", err)
		}
		d.Message = value
	}
	if d.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "draftMessage#fd8e711f", "entities", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode %s: field %s: %w", "draftMessage#fd8e711f", "entities", err)
			}
			d.Entities = append(d.Entities, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "draftMessage#fd8e711f", "date", err)
		}
		d.Date = value
	}
	return nil
}

// construct implements constructor of DraftMessageClass.
func (d DraftMessage) construct() DraftMessageClass { return &d }

// Ensuring interfaces in compile-time for DraftMessage.
var (
	_ bin.Encoder     = &DraftMessage{}
	_ bin.Decoder     = &DraftMessage{}
	_ bin.BareEncoder = &DraftMessage{}
	_ bin.BareDecoder = &DraftMessage{}

	_ DraftMessageClass = &DraftMessage{}
)

// DraftMessageClass represents DraftMessage generic type.
//
// See https://core.telegram.org/type/DraftMessage for reference.
//
// Example:
//  g, err := tg.DecodeDraftMessage(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.DraftMessageEmpty: // draftMessageEmpty#1b0c841a
//  case *tg.DraftMessage: // draftMessage#fd8e711f
//  default: panic(v)
//  }
type DraftMessageClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() DraftMessageClass

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

	// AsNotEmpty tries to map DraftMessageClass to DraftMessage.
	AsNotEmpty() (*DraftMessage, bool)
}

// AsNotEmpty tries to map DraftMessageEmpty to DraftMessage.
func (d *DraftMessageEmpty) AsNotEmpty() (*DraftMessage, bool) {
	return nil, false
}

// AsNotEmpty tries to map DraftMessage to DraftMessage.
func (d *DraftMessage) AsNotEmpty() (*DraftMessage, bool) {
	return d, true
}

// DecodeDraftMessage implements binary de-serialization for DraftMessageClass.
func DecodeDraftMessage(buf *bin.Buffer) (DraftMessageClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case DraftMessageEmptyTypeID:
		// Decoding draftMessageEmpty#1b0c841a.
		v := DraftMessageEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode %s: %w", "DraftMessageClass", err)
		}
		return &v, nil
	case DraftMessageTypeID:
		// Decoding draftMessage#fd8e711f.
		v := DraftMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode %s: %w", "DraftMessageClass", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode %s: %w", "DraftMessageClass", bin.NewUnexpectedID(id))
	}
}

// DraftMessage boxes the DraftMessageClass providing a helper.
type DraftMessageBox struct {
	DraftMessage DraftMessageClass
}

// Decode implements bin.Decoder for DraftMessageBox.
func (b *DraftMessageBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode %sBox to nil", "DraftMessage")
	}
	v, err := DecodeDraftMessage(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.DraftMessage = v
	return nil
}

// Encode implements bin.Encode for DraftMessageBox.
func (b *DraftMessageBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.DraftMessage == nil {
		return fmt.Errorf("unable to encode %s as nil", "DraftMessageClass")
	}
	return b.DraftMessage.Encode(buf)
}

// DraftMessageClassArray is adapter for slice of DraftMessageClass.
type DraftMessageClassArray []DraftMessageClass

// Sort sorts slice of DraftMessageClass.
func (s DraftMessageClassArray) Sort(less func(a, b DraftMessageClass) bool) DraftMessageClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of DraftMessageClass.
func (s DraftMessageClassArray) SortStable(less func(a, b DraftMessageClass) bool) DraftMessageClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of DraftMessageClass.
func (s DraftMessageClassArray) Retain(keep func(x DraftMessageClass) bool) DraftMessageClassArray {
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
func (s DraftMessageClassArray) First() (v DraftMessageClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DraftMessageClassArray) Last() (v DraftMessageClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DraftMessageClassArray) PopFirst() (v DraftMessageClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero DraftMessageClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DraftMessageClassArray) Pop() (v DraftMessageClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsDraftMessageEmpty returns copy with only DraftMessageEmpty constructors.
func (s DraftMessageClassArray) AsDraftMessageEmpty() (to DraftMessageEmptyArray) {
	for _, elem := range s {
		value, ok := elem.(*DraftMessageEmpty)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsDraftMessage returns copy with only DraftMessage constructors.
func (s DraftMessageClassArray) AsDraftMessage() (to DraftMessageArray) {
	for _, elem := range s {
		value, ok := elem.(*DraftMessage)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s DraftMessageClassArray) AppendOnlyNotEmpty(to []*DraftMessage) []*DraftMessage {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotEmpty returns copy with only NotEmpty constructors.
func (s DraftMessageClassArray) AsNotEmpty() (to []*DraftMessage) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s DraftMessageClassArray) FirstAsNotEmpty() (v *DraftMessage, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s DraftMessageClassArray) LastAsNotEmpty() (v *DraftMessage, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *DraftMessageClassArray) PopFirstAsNotEmpty() (v *DraftMessage, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *DraftMessageClassArray) PopAsNotEmpty() (v *DraftMessage, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// DraftMessageEmptyArray is adapter for slice of DraftMessageEmpty.
type DraftMessageEmptyArray []DraftMessageEmpty

// Sort sorts slice of DraftMessageEmpty.
func (s DraftMessageEmptyArray) Sort(less func(a, b DraftMessageEmpty) bool) DraftMessageEmptyArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of DraftMessageEmpty.
func (s DraftMessageEmptyArray) SortStable(less func(a, b DraftMessageEmpty) bool) DraftMessageEmptyArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of DraftMessageEmpty.
func (s DraftMessageEmptyArray) Retain(keep func(x DraftMessageEmpty) bool) DraftMessageEmptyArray {
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
func (s DraftMessageEmptyArray) First() (v DraftMessageEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DraftMessageEmptyArray) Last() (v DraftMessageEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DraftMessageEmptyArray) PopFirst() (v DraftMessageEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero DraftMessageEmpty
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DraftMessageEmptyArray) Pop() (v DraftMessageEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// DraftMessageArray is adapter for slice of DraftMessage.
type DraftMessageArray []DraftMessage

// Sort sorts slice of DraftMessage.
func (s DraftMessageArray) Sort(less func(a, b DraftMessage) bool) DraftMessageArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of DraftMessage.
func (s DraftMessageArray) SortStable(less func(a, b DraftMessage) bool) DraftMessageArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of DraftMessage.
func (s DraftMessageArray) Retain(keep func(x DraftMessage) bool) DraftMessageArray {
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
func (s DraftMessageArray) First() (v DraftMessage, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DraftMessageArray) Last() (v DraftMessage, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DraftMessageArray) PopFirst() (v DraftMessage, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero DraftMessage
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DraftMessageArray) Pop() (v DraftMessage, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByDate sorts slice of DraftMessage by Date.
func (s DraftMessageArray) SortByDate() DraftMessageArray {
	return s.Sort(func(a, b DraftMessage) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of DraftMessage by Date.
func (s DraftMessageArray) SortStableByDate() DraftMessageArray {
	return s.SortStable(func(a, b DraftMessage) bool {
		return a.GetDate() < b.GetDate()
	})
}
