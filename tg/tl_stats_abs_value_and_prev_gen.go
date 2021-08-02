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

// StatsAbsValueAndPrev represents TL type `statsAbsValueAndPrev#cb43acde`.
// Statistics value couple; intial and final value for period of time currently in
// consideration
//
// See https://core.telegram.org/constructor/statsAbsValueAndPrev for reference.
type StatsAbsValueAndPrev struct {
	// Current value
	Current float64
	// Previous value
	Previous float64
}

// StatsAbsValueAndPrevTypeID is TL type id of StatsAbsValueAndPrev.
const StatsAbsValueAndPrevTypeID = 0xcb43acde

func (s *StatsAbsValueAndPrev) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Current == 0) {
		return false
	}
	if !(s.Previous == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StatsAbsValueAndPrev) String() string {
	if s == nil {
		return "StatsAbsValueAndPrev(nil)"
	}
	type Alias StatsAbsValueAndPrev
	return fmt.Sprintf("StatsAbsValueAndPrev%+v", Alias(*s))
}

// FillFrom fills StatsAbsValueAndPrev from given interface.
func (s *StatsAbsValueAndPrev) FillFrom(from interface {
	GetCurrent() (value float64)
	GetPrevious() (value float64)
}) {
	s.Current = from.GetCurrent()
	s.Previous = from.GetPrevious()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsAbsValueAndPrev) TypeID() uint32 {
	return StatsAbsValueAndPrevTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsAbsValueAndPrev) TypeName() string {
	return "statsAbsValueAndPrev"
}

// TypeInfo returns info about TL type.
func (s *StatsAbsValueAndPrev) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "statsAbsValueAndPrev",
		ID:   StatsAbsValueAndPrevTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Current",
			SchemaName: "current",
		},
		{
			Name:       "Previous",
			SchemaName: "previous",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StatsAbsValueAndPrev) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode %s as nil", "statsAbsValueAndPrev#cb43acde")
	}
	b.PutID(StatsAbsValueAndPrevTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StatsAbsValueAndPrev) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode %s as nil", "statsAbsValueAndPrev#cb43acde")
	}
	b.PutDouble(s.Current)
	b.PutDouble(s.Previous)
	return nil
}

// GetCurrent returns value of Current field.
func (s *StatsAbsValueAndPrev) GetCurrent() (value float64) {
	return s.Current
}

// GetPrevious returns value of Previous field.
func (s *StatsAbsValueAndPrev) GetPrevious() (value float64) {
	return s.Previous
}

// Decode implements bin.Decoder.
func (s *StatsAbsValueAndPrev) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode %s to nil", "statsAbsValueAndPrev#cb43acde")
	}
	if err := b.ConsumeID(StatsAbsValueAndPrevTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "statsAbsValueAndPrev#cb43acde", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StatsAbsValueAndPrev) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode %s to nil", "statsAbsValueAndPrev#cb43acde")
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "statsAbsValueAndPrev#cb43acde", "current", err)
		}
		s.Current = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "statsAbsValueAndPrev#cb43acde", "previous", err)
		}
		s.Previous = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsAbsValueAndPrev.
var (
	_ bin.Encoder     = &StatsAbsValueAndPrev{}
	_ bin.Decoder     = &StatsAbsValueAndPrev{}
	_ bin.BareEncoder = &StatsAbsValueAndPrev{}
	_ bin.BareDecoder = &StatsAbsValueAndPrev{}
)
