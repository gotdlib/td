// Code generated by gotdgen, DO NOT EDIT.

package td

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

// InvokeWithLayer represents TL type `invokeWithLayer#da9b0d0d`.
//
// See https://localhost:80/doc/constructor/invokeWithLayer for reference.
type InvokeWithLayer struct {
	// Layer field of InvokeWithLayer.
	Layer int
	// Query field of InvokeWithLayer.
	Query bin.Object
}

// InvokeWithLayerTypeID is TL type id of InvokeWithLayer.
const InvokeWithLayerTypeID = 0xda9b0d0d

func (i *InvokeWithLayer) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Layer == 0) {
		return false
	}
	if !(i.Query == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InvokeWithLayer) String() string {
	if i == nil {
		return "InvokeWithLayer(nil)"
	}
	type Alias InvokeWithLayer
	return fmt.Sprintf("InvokeWithLayer%+v", Alias(*i))
}

// FillFrom fills InvokeWithLayer from given interface.
func (i *InvokeWithLayer) FillFrom(from interface {
	GetLayer() (value int)
	GetQuery() (value bin.Object)
}) {
	i.Layer = from.GetLayer()
	i.Query = from.GetQuery()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InvokeWithLayer) TypeID() uint32 {
	return InvokeWithLayerTypeID
}

// TypeName returns name of type in TL schema.
func (*InvokeWithLayer) TypeName() string {
	return "invokeWithLayer"
}

// TypeInfo returns info about TL type.
func (i *InvokeWithLayer) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "invokeWithLayer",
		ID:   InvokeWithLayerTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Layer",
			SchemaName: "layer",
		},
		{
			Name:       "Query",
			SchemaName: "query",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InvokeWithLayer) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode %s as nil", "invokeWithLayer#da9b0d0d")
	}
	b.PutID(InvokeWithLayerTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InvokeWithLayer) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode %s as nil", "invokeWithLayer#da9b0d0d")
	}
	b.PutInt(i.Layer)
	if err := i.Query.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "invokeWithLayer#da9b0d0d", "query", err)
	}
	return nil
}

// GetLayer returns value of Layer field.
func (i *InvokeWithLayer) GetLayer() (value int) {
	return i.Layer
}

// GetQuery returns value of Query field.
func (i *InvokeWithLayer) GetQuery() (value bin.Object) {
	return i.Query
}

// Decode implements bin.Decoder.
func (i *InvokeWithLayer) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode %s to nil", "invokeWithLayer#da9b0d0d")
	}
	if err := b.ConsumeID(InvokeWithLayerTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "invokeWithLayer#da9b0d0d", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InvokeWithLayer) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode %s to nil", "invokeWithLayer#da9b0d0d")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "invokeWithLayer#da9b0d0d", "layer", err)
		}
		i.Layer = value
	}
	{
		if err := i.Query.Decode(b); err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "invokeWithLayer#da9b0d0d", "query", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for InvokeWithLayer.
var (
	_ bin.Encoder     = &InvokeWithLayer{}
	_ bin.Decoder     = &InvokeWithLayer{}
	_ bin.BareEncoder = &InvokeWithLayer{}
	_ bin.BareDecoder = &InvokeWithLayer{}
)
