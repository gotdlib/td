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

// StickerSetCovered represents TL type `stickerSetCovered#6410a5d2`.
// Stickerset, with a specific sticker as preview
//
// See https://core.telegram.org/constructor/stickerSetCovered for reference.
type StickerSetCovered struct {
	// Stickerset
	Set StickerSet
	// Preview
	Cover DocumentClass
}

// StickerSetCoveredTypeID is TL type id of StickerSetCovered.
const StickerSetCoveredTypeID = 0x6410a5d2

func (s *StickerSetCovered) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Set.Zero()) {
		return false
	}
	if !(s.Cover == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StickerSetCovered) String() string {
	if s == nil {
		return "StickerSetCovered(nil)"
	}
	type Alias StickerSetCovered
	return fmt.Sprintf("StickerSetCovered%+v", Alias(*s))
}

// FillFrom fills StickerSetCovered from given interface.
func (s *StickerSetCovered) FillFrom(from interface {
	GetSet() (value StickerSet)
	GetCover() (value DocumentClass)
}) {
	s.Set = from.GetSet()
	s.Cover = from.GetCover()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StickerSetCovered) TypeID() uint32 {
	return StickerSetCoveredTypeID
}

// TypeName returns name of type in TL schema.
func (*StickerSetCovered) TypeName() string {
	return "stickerSetCovered"
}

// TypeInfo returns info about TL type.
func (s *StickerSetCovered) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stickerSetCovered",
		ID:   StickerSetCoveredTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Set",
			SchemaName: "set",
		},
		{
			Name:       "Cover",
			SchemaName: "cover",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StickerSetCovered) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode %s as nil", "stickerSetCovered#6410a5d2")
	}
	b.PutID(StickerSetCoveredTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StickerSetCovered) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode %s as nil", "stickerSetCovered#6410a5d2")
	}
	if err := s.Set.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "stickerSetCovered#6410a5d2", "set", err)
	}
	if s.Cover == nil {
		return fmt.Errorf("unable to encode %s: field %s is nil", "stickerSetCovered#6410a5d2", "cover")
	}
	if err := s.Cover.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "stickerSetCovered#6410a5d2", "cover", err)
	}
	return nil
}

// GetSet returns value of Set field.
func (s *StickerSetCovered) GetSet() (value StickerSet) {
	return s.Set
}

// GetCover returns value of Cover field.
func (s *StickerSetCovered) GetCover() (value DocumentClass) {
	return s.Cover
}

// Decode implements bin.Decoder.
func (s *StickerSetCovered) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode %s to nil", "stickerSetCovered#6410a5d2")
	}
	if err := b.ConsumeID(StickerSetCoveredTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "stickerSetCovered#6410a5d2", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StickerSetCovered) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode %s to nil", "stickerSetCovered#6410a5d2")
	}
	{
		if err := s.Set.Decode(b); err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "stickerSetCovered#6410a5d2", "set", err)
		}
	}
	{
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "stickerSetCovered#6410a5d2", "cover", err)
		}
		s.Cover = value
	}
	return nil
}

// construct implements constructor of StickerSetCoveredClass.
func (s StickerSetCovered) construct() StickerSetCoveredClass { return &s }

// Ensuring interfaces in compile-time for StickerSetCovered.
var (
	_ bin.Encoder     = &StickerSetCovered{}
	_ bin.Decoder     = &StickerSetCovered{}
	_ bin.BareEncoder = &StickerSetCovered{}
	_ bin.BareDecoder = &StickerSetCovered{}

	_ StickerSetCoveredClass = &StickerSetCovered{}
)

// StickerSetMultiCovered represents TL type `stickerSetMultiCovered#3407e51b`.
// Stickerset, with a specific stickers as preview
//
// See https://core.telegram.org/constructor/stickerSetMultiCovered for reference.
type StickerSetMultiCovered struct {
	// Stickerset
	Set StickerSet
	// Preview stickers
	Covers []DocumentClass
}

// StickerSetMultiCoveredTypeID is TL type id of StickerSetMultiCovered.
const StickerSetMultiCoveredTypeID = 0x3407e51b

func (s *StickerSetMultiCovered) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Set.Zero()) {
		return false
	}
	if !(s.Covers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StickerSetMultiCovered) String() string {
	if s == nil {
		return "StickerSetMultiCovered(nil)"
	}
	type Alias StickerSetMultiCovered
	return fmt.Sprintf("StickerSetMultiCovered%+v", Alias(*s))
}

// FillFrom fills StickerSetMultiCovered from given interface.
func (s *StickerSetMultiCovered) FillFrom(from interface {
	GetSet() (value StickerSet)
	GetCovers() (value []DocumentClass)
}) {
	s.Set = from.GetSet()
	s.Covers = from.GetCovers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StickerSetMultiCovered) TypeID() uint32 {
	return StickerSetMultiCoveredTypeID
}

// TypeName returns name of type in TL schema.
func (*StickerSetMultiCovered) TypeName() string {
	return "stickerSetMultiCovered"
}

// TypeInfo returns info about TL type.
func (s *StickerSetMultiCovered) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stickerSetMultiCovered",
		ID:   StickerSetMultiCoveredTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Set",
			SchemaName: "set",
		},
		{
			Name:       "Covers",
			SchemaName: "covers",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StickerSetMultiCovered) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode %s as nil", "stickerSetMultiCovered#3407e51b")
	}
	b.PutID(StickerSetMultiCoveredTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StickerSetMultiCovered) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode %s as nil", "stickerSetMultiCovered#3407e51b")
	}
	if err := s.Set.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "stickerSetMultiCovered#3407e51b", "set", err)
	}
	b.PutVectorHeader(len(s.Covers))
	for idx, v := range s.Covers {
		if v == nil {
			return fmt.Errorf("unable to encode %s: field %s element with index %d is nil", "stickerSetMultiCovered#3407e51b", "covers", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode %s: field %s element with index %d: %w", "stickerSetMultiCovered#3407e51b", "covers", idx, err)
		}
	}
	return nil
}

// GetSet returns value of Set field.
func (s *StickerSetMultiCovered) GetSet() (value StickerSet) {
	return s.Set
}

// GetCovers returns value of Covers field.
func (s *StickerSetMultiCovered) GetCovers() (value []DocumentClass) {
	return s.Covers
}

// MapCovers returns field Covers wrapped in DocumentClassArray helper.
func (s *StickerSetMultiCovered) MapCovers() (value DocumentClassArray) {
	return DocumentClassArray(s.Covers)
}

// Decode implements bin.Decoder.
func (s *StickerSetMultiCovered) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode %s to nil", "stickerSetMultiCovered#3407e51b")
	}
	if err := b.ConsumeID(StickerSetMultiCoveredTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "stickerSetMultiCovered#3407e51b", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StickerSetMultiCovered) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode %s to nil", "stickerSetMultiCovered#3407e51b")
	}
	{
		if err := s.Set.Decode(b); err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "stickerSetMultiCovered#3407e51b", "set", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "stickerSetMultiCovered#3407e51b", "covers", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocument(b)
			if err != nil {
				return fmt.Errorf("unable to decode %s: field %s: %w", "stickerSetMultiCovered#3407e51b", "covers", err)
			}
			s.Covers = append(s.Covers, value)
		}
	}
	return nil
}

// construct implements constructor of StickerSetCoveredClass.
func (s StickerSetMultiCovered) construct() StickerSetCoveredClass { return &s }

// Ensuring interfaces in compile-time for StickerSetMultiCovered.
var (
	_ bin.Encoder     = &StickerSetMultiCovered{}
	_ bin.Decoder     = &StickerSetMultiCovered{}
	_ bin.BareEncoder = &StickerSetMultiCovered{}
	_ bin.BareDecoder = &StickerSetMultiCovered{}

	_ StickerSetCoveredClass = &StickerSetMultiCovered{}
)

// StickerSetCoveredClass represents StickerSetCovered generic type.
//
// See https://core.telegram.org/type/StickerSetCovered for reference.
//
// Example:
//  g, err := tg.DecodeStickerSetCovered(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.StickerSetCovered: // stickerSetCovered#6410a5d2
//  case *tg.StickerSetMultiCovered: // stickerSetMultiCovered#3407e51b
//  default: panic(v)
//  }
type StickerSetCoveredClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() StickerSetCoveredClass

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

	// Stickerset
	GetSet() (value StickerSet)
}

// DecodeStickerSetCovered implements binary de-serialization for StickerSetCoveredClass.
func DecodeStickerSetCovered(buf *bin.Buffer) (StickerSetCoveredClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case StickerSetCoveredTypeID:
		// Decoding stickerSetCovered#6410a5d2.
		v := StickerSetCovered{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode %s: %w", "StickerSetCoveredClass", err)
		}
		return &v, nil
	case StickerSetMultiCoveredTypeID:
		// Decoding stickerSetMultiCovered#3407e51b.
		v := StickerSetMultiCovered{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode %s: %w", "StickerSetCoveredClass", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode %s: %w", "StickerSetCoveredClass", bin.NewUnexpectedID(id))
	}
}

// StickerSetCovered boxes the StickerSetCoveredClass providing a helper.
type StickerSetCoveredBox struct {
	StickerSetCovered StickerSetCoveredClass
}

// Decode implements bin.Decoder for StickerSetCoveredBox.
func (b *StickerSetCoveredBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode %sBox to nil", "StickerSetCovered")
	}
	v, err := DecodeStickerSetCovered(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StickerSetCovered = v
	return nil
}

// Encode implements bin.Encode for StickerSetCoveredBox.
func (b *StickerSetCoveredBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.StickerSetCovered == nil {
		return fmt.Errorf("unable to encode %s as nil", "StickerSetCoveredClass")
	}
	return b.StickerSetCovered.Encode(buf)
}

// StickerSetCoveredClassArray is adapter for slice of StickerSetCoveredClass.
type StickerSetCoveredClassArray []StickerSetCoveredClass

// Sort sorts slice of StickerSetCoveredClass.
func (s StickerSetCoveredClassArray) Sort(less func(a, b StickerSetCoveredClass) bool) StickerSetCoveredClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of StickerSetCoveredClass.
func (s StickerSetCoveredClassArray) SortStable(less func(a, b StickerSetCoveredClass) bool) StickerSetCoveredClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of StickerSetCoveredClass.
func (s StickerSetCoveredClassArray) Retain(keep func(x StickerSetCoveredClass) bool) StickerSetCoveredClassArray {
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
func (s StickerSetCoveredClassArray) First() (v StickerSetCoveredClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s StickerSetCoveredClassArray) Last() (v StickerSetCoveredClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *StickerSetCoveredClassArray) PopFirst() (v StickerSetCoveredClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero StickerSetCoveredClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *StickerSetCoveredClassArray) Pop() (v StickerSetCoveredClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsStickerSetCovered returns copy with only StickerSetCovered constructors.
func (s StickerSetCoveredClassArray) AsStickerSetCovered() (to StickerSetCoveredArray) {
	for _, elem := range s {
		value, ok := elem.(*StickerSetCovered)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsStickerSetMultiCovered returns copy with only StickerSetMultiCovered constructors.
func (s StickerSetCoveredClassArray) AsStickerSetMultiCovered() (to StickerSetMultiCoveredArray) {
	for _, elem := range s {
		value, ok := elem.(*StickerSetMultiCovered)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// StickerSetCoveredArray is adapter for slice of StickerSetCovered.
type StickerSetCoveredArray []StickerSetCovered

// Sort sorts slice of StickerSetCovered.
func (s StickerSetCoveredArray) Sort(less func(a, b StickerSetCovered) bool) StickerSetCoveredArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of StickerSetCovered.
func (s StickerSetCoveredArray) SortStable(less func(a, b StickerSetCovered) bool) StickerSetCoveredArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of StickerSetCovered.
func (s StickerSetCoveredArray) Retain(keep func(x StickerSetCovered) bool) StickerSetCoveredArray {
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
func (s StickerSetCoveredArray) First() (v StickerSetCovered, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s StickerSetCoveredArray) Last() (v StickerSetCovered, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *StickerSetCoveredArray) PopFirst() (v StickerSetCovered, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero StickerSetCovered
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *StickerSetCoveredArray) Pop() (v StickerSetCovered, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// StickerSetMultiCoveredArray is adapter for slice of StickerSetMultiCovered.
type StickerSetMultiCoveredArray []StickerSetMultiCovered

// Sort sorts slice of StickerSetMultiCovered.
func (s StickerSetMultiCoveredArray) Sort(less func(a, b StickerSetMultiCovered) bool) StickerSetMultiCoveredArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of StickerSetMultiCovered.
func (s StickerSetMultiCoveredArray) SortStable(less func(a, b StickerSetMultiCovered) bool) StickerSetMultiCoveredArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of StickerSetMultiCovered.
func (s StickerSetMultiCoveredArray) Retain(keep func(x StickerSetMultiCovered) bool) StickerSetMultiCoveredArray {
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
func (s StickerSetMultiCoveredArray) First() (v StickerSetMultiCovered, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s StickerSetMultiCoveredArray) Last() (v StickerSetMultiCovered, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *StickerSetMultiCoveredArray) PopFirst() (v StickerSetMultiCovered, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero StickerSetMultiCovered
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *StickerSetMultiCoveredArray) Pop() (v StickerSetMultiCovered, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
