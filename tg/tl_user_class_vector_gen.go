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

// UserClassVector is a box for Vector<User>
type UserClassVector struct {
	// Elements of Vector<User>
	Elems []UserClass
}

// UserClassVectorTypeID is TL type id of UserClassVector.
const UserClassVectorTypeID = bin.TypeVector

func (vec *UserClassVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *UserClassVector) String() string {
	if vec == nil {
		return "UserClassVector(nil)"
	}
	type Alias UserClassVector
	return fmt.Sprintf("UserClassVector%+v", Alias(*vec))
}

// FillFrom fills UserClassVector from given interface.
func (vec *UserClassVector) FillFrom(from interface {
	GetElems() (value []UserClass)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UserClassVector) TypeID() uint32 {
	return UserClassVectorTypeID
}

// TypeName returns name of type in TL schema.
func (*UserClassVector) TypeName() string {
	return ""
}

// TypeInfo returns info about TL type.
func (vec *UserClassVector) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "",
		ID:   UserClassVectorTypeID,
	}
	if vec == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Elems",
			SchemaName: "Elems",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (vec *UserClassVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode %s as nil", "Vector<User>")
	}

	return vec.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (vec *UserClassVector) EncodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode %s as nil", "Vector<User>")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if v == nil {
			return fmt.Errorf("unable to encode %s: field %s element with index %d is nil", "Vector<User>", "Elems", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode %s: field %s element with index %d: %w", "Vector<User>", "Elems", idx, err)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *UserClassVector) GetElems() (value []UserClass) {
	return vec.Elems
}

// MapElems returns field Elems wrapped in UserClassArray helper.
func (vec *UserClassVector) MapElems() (value UserClassArray) {
	return UserClassArray(vec.Elems)
}

// Decode implements bin.Decoder.
func (vec *UserClassVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode %s to nil", "Vector<User>")
	}

	return vec.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (vec *UserClassVector) DecodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode %s to nil", "Vector<User>")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "Vector<User>", "Elems", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode %s: field %s: %w", "Vector<User>", "Elems", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for UserClassVector.
var (
	_ bin.Encoder     = &UserClassVector{}
	_ bin.Decoder     = &UserClassVector{}
	_ bin.BareEncoder = &UserClassVector{}
	_ bin.BareDecoder = &UserClassVector{}
)
