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

// LangPackLanguageVector is a box for Vector<LangPackLanguage>
type LangPackLanguageVector struct {
	// Elements of Vector<LangPackLanguage>
	Elems []LangPackLanguage
}

// LangPackLanguageVectorTypeID is TL type id of LangPackLanguageVector.
const LangPackLanguageVectorTypeID = bin.TypeVector

func (vec *LangPackLanguageVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *LangPackLanguageVector) String() string {
	if vec == nil {
		return "LangPackLanguageVector(nil)"
	}
	type Alias LangPackLanguageVector
	return fmt.Sprintf("LangPackLanguageVector%+v", Alias(*vec))
}

// FillFrom fills LangPackLanguageVector from given interface.
func (vec *LangPackLanguageVector) FillFrom(from interface {
	GetElems() (value []LangPackLanguage)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LangPackLanguageVector) TypeID() uint32 {
	return LangPackLanguageVectorTypeID
}

// TypeName returns name of type in TL schema.
func (*LangPackLanguageVector) TypeName() string {
	return ""
}

// TypeInfo returns info about TL type.
func (vec *LangPackLanguageVector) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "",
		ID:   LangPackLanguageVectorTypeID,
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
func (vec *LangPackLanguageVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode %s as nil", "Vector<LangPackLanguage>")
	}

	return vec.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (vec *LangPackLanguageVector) EncodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode %s as nil", "Vector<LangPackLanguage>")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode %s: field %s element with index %d: %w", "Vector<LangPackLanguage>", "Elems", idx, err)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *LangPackLanguageVector) GetElems() (value []LangPackLanguage) {
	return vec.Elems
}

// Decode implements bin.Decoder.
func (vec *LangPackLanguageVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode %s to nil", "Vector<LangPackLanguage>")
	}

	return vec.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (vec *LangPackLanguageVector) DecodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode %s to nil", "Vector<LangPackLanguage>")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "Vector<LangPackLanguage>", "Elems", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value LangPackLanguage
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode %s: field %s: %w", "Vector<LangPackLanguage>", "Elems", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for LangPackLanguageVector.
var (
	_ bin.Encoder     = &LangPackLanguageVector{}
	_ bin.Decoder     = &LangPackLanguageVector{}
	_ bin.BareEncoder = &LangPackLanguageVector{}
	_ bin.BareDecoder = &LangPackLanguageVector{}
)
