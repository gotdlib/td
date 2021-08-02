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

// MaskCoords represents TL type `maskCoords#aed6dbb2`.
// Position on a photo where a mask should be placed
// The n position indicates where the mask should be placed:
//
// See https://core.telegram.org/constructor/maskCoords for reference.
type MaskCoords struct {
	// Part of the face, relative to which the mask should be placed
	N int
	// Shift by X-axis measured in widths of the mask scaled to the face size, from left to
	// right. (For example, -1.0 will place the mask just to the left of the default mask
	// position)
	X float64
	// Shift by Y-axis measured in widths of the mask scaled to the face size, from left to
	// right. (For example, -1.0 will place the mask just to the left of the default mask
	// position)
	Y float64
	// Mask scaling coefficient. (For example, 2.0 means a doubled size)
	Zoom float64
}

// MaskCoordsTypeID is TL type id of MaskCoords.
const MaskCoordsTypeID = 0xaed6dbb2

func (m *MaskCoords) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.N == 0) {
		return false
	}
	if !(m.X == 0) {
		return false
	}
	if !(m.Y == 0) {
		return false
	}
	if !(m.Zoom == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MaskCoords) String() string {
	if m == nil {
		return "MaskCoords(nil)"
	}
	type Alias MaskCoords
	return fmt.Sprintf("MaskCoords%+v", Alias(*m))
}

// FillFrom fills MaskCoords from given interface.
func (m *MaskCoords) FillFrom(from interface {
	GetN() (value int)
	GetX() (value float64)
	GetY() (value float64)
	GetZoom() (value float64)
}) {
	m.N = from.GetN()
	m.X = from.GetX()
	m.Y = from.GetY()
	m.Zoom = from.GetZoom()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MaskCoords) TypeID() uint32 {
	return MaskCoordsTypeID
}

// TypeName returns name of type in TL schema.
func (*MaskCoords) TypeName() string {
	return "maskCoords"
}

// TypeInfo returns info about TL type.
func (m *MaskCoords) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "maskCoords",
		ID:   MaskCoordsTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "N",
			SchemaName: "n",
		},
		{
			Name:       "X",
			SchemaName: "x",
		},
		{
			Name:       "Y",
			SchemaName: "y",
		},
		{
			Name:       "Zoom",
			SchemaName: "zoom",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MaskCoords) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode %s as nil", "maskCoords#aed6dbb2")
	}
	b.PutID(MaskCoordsTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MaskCoords) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode %s as nil", "maskCoords#aed6dbb2")
	}
	b.PutInt(m.N)
	b.PutDouble(m.X)
	b.PutDouble(m.Y)
	b.PutDouble(m.Zoom)
	return nil
}

// GetN returns value of N field.
func (m *MaskCoords) GetN() (value int) {
	return m.N
}

// GetX returns value of X field.
func (m *MaskCoords) GetX() (value float64) {
	return m.X
}

// GetY returns value of Y field.
func (m *MaskCoords) GetY() (value float64) {
	return m.Y
}

// GetZoom returns value of Zoom field.
func (m *MaskCoords) GetZoom() (value float64) {
	return m.Zoom
}

// Decode implements bin.Decoder.
func (m *MaskCoords) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode %s to nil", "maskCoords#aed6dbb2")
	}
	if err := b.ConsumeID(MaskCoordsTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "maskCoords#aed6dbb2", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MaskCoords) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode %s to nil", "maskCoords#aed6dbb2")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "maskCoords#aed6dbb2", "n", err)
		}
		m.N = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "maskCoords#aed6dbb2", "x", err)
		}
		m.X = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "maskCoords#aed6dbb2", "y", err)
		}
		m.Y = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "maskCoords#aed6dbb2", "zoom", err)
		}
		m.Zoom = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MaskCoords.
var (
	_ bin.Encoder     = &MaskCoords{}
	_ bin.Decoder     = &MaskCoords{}
	_ bin.BareEncoder = &MaskCoords{}
	_ bin.BareDecoder = &MaskCoords{}
)
