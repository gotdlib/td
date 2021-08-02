// Code generated by gotdgen, DO NOT EDIT.

package e2e

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

// InputStickerSetShortName represents TL type `inputStickerSetShortName#861cc8a0`.
// Stickerset by short name, from tg://addstickers?set=short_name
//
// See https://core.telegram.org/constructor/inputStickerSetShortName for reference.
type InputStickerSetShortName struct {
	// From tg://addstickers?set=short_name
	ShortName string
}

// InputStickerSetShortNameTypeID is TL type id of InputStickerSetShortName.
const InputStickerSetShortNameTypeID = 0x861cc8a0

func (i *InputStickerSetShortName) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ShortName == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStickerSetShortName) String() string {
	if i == nil {
		return "InputStickerSetShortName(nil)"
	}
	type Alias InputStickerSetShortName
	return fmt.Sprintf("InputStickerSetShortName%+v", Alias(*i))
}

// FillFrom fills InputStickerSetShortName from given interface.
func (i *InputStickerSetShortName) FillFrom(from interface {
	GetShortName() (value string)
}) {
	i.ShortName = from.GetShortName()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputStickerSetShortName) TypeID() uint32 {
	return InputStickerSetShortNameTypeID
}

// TypeName returns name of type in TL schema.
func (*InputStickerSetShortName) TypeName() string {
	return "inputStickerSetShortName"
}

// TypeInfo returns info about TL type.
func (i *InputStickerSetShortName) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputStickerSetShortName",
		ID:   InputStickerSetShortNameTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ShortName",
			SchemaName: "short_name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputStickerSetShortName) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode %s as nil", "inputStickerSetShortName#861cc8a0")
	}
	b.PutID(InputStickerSetShortNameTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputStickerSetShortName) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode %s as nil", "inputStickerSetShortName#861cc8a0")
	}
	b.PutString(i.ShortName)
	return nil
}

// GetShortName returns value of ShortName field.
func (i *InputStickerSetShortName) GetShortName() (value string) {
	return i.ShortName
}

// Decode implements bin.Decoder.
func (i *InputStickerSetShortName) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode %s to nil", "inputStickerSetShortName#861cc8a0")
	}
	if err := b.ConsumeID(InputStickerSetShortNameTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "inputStickerSetShortName#861cc8a0", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputStickerSetShortName) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode %s to nil", "inputStickerSetShortName#861cc8a0")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "inputStickerSetShortName#861cc8a0", "short_name", err)
		}
		i.ShortName = value
	}
	return nil
}

// construct implements constructor of InputStickerSetClass.
func (i InputStickerSetShortName) construct() InputStickerSetClass { return &i }

// Ensuring interfaces in compile-time for InputStickerSetShortName.
var (
	_ bin.Encoder     = &InputStickerSetShortName{}
	_ bin.Decoder     = &InputStickerSetShortName{}
	_ bin.BareEncoder = &InputStickerSetShortName{}
	_ bin.BareDecoder = &InputStickerSetShortName{}

	_ InputStickerSetClass = &InputStickerSetShortName{}
)

// InputStickerSetEmpty represents TL type `inputStickerSetEmpty#ffb62b95`.
// Empty constructor
//
// See https://core.telegram.org/constructor/inputStickerSetEmpty for reference.
type InputStickerSetEmpty struct {
}

// InputStickerSetEmptyTypeID is TL type id of InputStickerSetEmpty.
const InputStickerSetEmptyTypeID = 0xffb62b95

func (i *InputStickerSetEmpty) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStickerSetEmpty) String() string {
	if i == nil {
		return "InputStickerSetEmpty(nil)"
	}
	type Alias InputStickerSetEmpty
	return fmt.Sprintf("InputStickerSetEmpty%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputStickerSetEmpty) TypeID() uint32 {
	return InputStickerSetEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*InputStickerSetEmpty) TypeName() string {
	return "inputStickerSetEmpty"
}

// TypeInfo returns info about TL type.
func (i *InputStickerSetEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputStickerSetEmpty",
		ID:   InputStickerSetEmptyTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputStickerSetEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode %s as nil", "inputStickerSetEmpty#ffb62b95")
	}
	b.PutID(InputStickerSetEmptyTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputStickerSetEmpty) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode %s as nil", "inputStickerSetEmpty#ffb62b95")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStickerSetEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode %s to nil", "inputStickerSetEmpty#ffb62b95")
	}
	if err := b.ConsumeID(InputStickerSetEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "inputStickerSetEmpty#ffb62b95", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputStickerSetEmpty) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode %s to nil", "inputStickerSetEmpty#ffb62b95")
	}
	return nil
}

// construct implements constructor of InputStickerSetClass.
func (i InputStickerSetEmpty) construct() InputStickerSetClass { return &i }

// Ensuring interfaces in compile-time for InputStickerSetEmpty.
var (
	_ bin.Encoder     = &InputStickerSetEmpty{}
	_ bin.Decoder     = &InputStickerSetEmpty{}
	_ bin.BareEncoder = &InputStickerSetEmpty{}
	_ bin.BareDecoder = &InputStickerSetEmpty{}

	_ InputStickerSetClass = &InputStickerSetEmpty{}
)

// InputStickerSetClass represents InputStickerSet generic type.
//
// See https://core.telegram.org/type/InputStickerSet for reference.
//
// Example:
//  g, err := e2e.DecodeInputStickerSet(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *e2e.InputStickerSetShortName: // inputStickerSetShortName#861cc8a0
//  case *e2e.InputStickerSetEmpty: // inputStickerSetEmpty#ffb62b95
//  default: panic(v)
//  }
type InputStickerSetClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputStickerSetClass

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

	// AsNotEmpty tries to map InputStickerSetClass to InputStickerSetShortName.
	AsNotEmpty() (*InputStickerSetShortName, bool)
}

// AsNotEmpty tries to map InputStickerSetShortName to InputStickerSetShortName.
func (i *InputStickerSetShortName) AsNotEmpty() (*InputStickerSetShortName, bool) {
	return i, true
}

// AsNotEmpty tries to map InputStickerSetEmpty to InputStickerSetShortName.
func (i *InputStickerSetEmpty) AsNotEmpty() (*InputStickerSetShortName, bool) {
	return nil, false
}

// DecodeInputStickerSet implements binary de-serialization for InputStickerSetClass.
func DecodeInputStickerSet(buf *bin.Buffer) (InputStickerSetClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputStickerSetShortNameTypeID:
		// Decoding inputStickerSetShortName#861cc8a0.
		v := InputStickerSetShortName{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode %s: %w", "InputStickerSetClass", err)
		}
		return &v, nil
	case InputStickerSetEmptyTypeID:
		// Decoding inputStickerSetEmpty#ffb62b95.
		v := InputStickerSetEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode %s: %w", "InputStickerSetClass", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode %s: %w", "InputStickerSetClass", bin.NewUnexpectedID(id))
	}
}

// InputStickerSet boxes the InputStickerSetClass providing a helper.
type InputStickerSetBox struct {
	InputStickerSet InputStickerSetClass
}

// Decode implements bin.Decoder for InputStickerSetBox.
func (b *InputStickerSetBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode %sBox to nil", "InputStickerSet")
	}
	v, err := DecodeInputStickerSet(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputStickerSet = v
	return nil
}

// Encode implements bin.Encode for InputStickerSetBox.
func (b *InputStickerSetBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputStickerSet == nil {
		return fmt.Errorf("unable to encode %s as nil", "InputStickerSetClass")
	}
	return b.InputStickerSet.Encode(buf)
}

// InputStickerSetClassArray is adapter for slice of InputStickerSetClass.
type InputStickerSetClassArray []InputStickerSetClass

// Sort sorts slice of InputStickerSetClass.
func (s InputStickerSetClassArray) Sort(less func(a, b InputStickerSetClass) bool) InputStickerSetClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputStickerSetClass.
func (s InputStickerSetClassArray) SortStable(less func(a, b InputStickerSetClass) bool) InputStickerSetClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputStickerSetClass.
func (s InputStickerSetClassArray) Retain(keep func(x InputStickerSetClass) bool) InputStickerSetClassArray {
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
func (s InputStickerSetClassArray) First() (v InputStickerSetClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputStickerSetClassArray) Last() (v InputStickerSetClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputStickerSetClassArray) PopFirst() (v InputStickerSetClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputStickerSetClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputStickerSetClassArray) Pop() (v InputStickerSetClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsInputStickerSetShortName returns copy with only InputStickerSetShortName constructors.
func (s InputStickerSetClassArray) AsInputStickerSetShortName() (to InputStickerSetShortNameArray) {
	for _, elem := range s {
		value, ok := elem.(*InputStickerSetShortName)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s InputStickerSetClassArray) AppendOnlyNotEmpty(to []*InputStickerSetShortName) []*InputStickerSetShortName {
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
func (s InputStickerSetClassArray) AsNotEmpty() (to []*InputStickerSetShortName) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s InputStickerSetClassArray) FirstAsNotEmpty() (v *InputStickerSetShortName, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s InputStickerSetClassArray) LastAsNotEmpty() (v *InputStickerSetShortName, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *InputStickerSetClassArray) PopFirstAsNotEmpty() (v *InputStickerSetShortName, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *InputStickerSetClassArray) PopAsNotEmpty() (v *InputStickerSetShortName, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// InputStickerSetShortNameArray is adapter for slice of InputStickerSetShortName.
type InputStickerSetShortNameArray []InputStickerSetShortName

// Sort sorts slice of InputStickerSetShortName.
func (s InputStickerSetShortNameArray) Sort(less func(a, b InputStickerSetShortName) bool) InputStickerSetShortNameArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputStickerSetShortName.
func (s InputStickerSetShortNameArray) SortStable(less func(a, b InputStickerSetShortName) bool) InputStickerSetShortNameArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputStickerSetShortName.
func (s InputStickerSetShortNameArray) Retain(keep func(x InputStickerSetShortName) bool) InputStickerSetShortNameArray {
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
func (s InputStickerSetShortNameArray) First() (v InputStickerSetShortName, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputStickerSetShortNameArray) Last() (v InputStickerSetShortName, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputStickerSetShortNameArray) PopFirst() (v InputStickerSetShortName, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputStickerSetShortName
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputStickerSetShortNameArray) Pop() (v InputStickerSetShortName, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
