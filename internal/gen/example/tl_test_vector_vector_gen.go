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

// TestVectorVector represents TL type `testVectorVector#69e8846c`.
//
// See https://localhost:80/doc/constructor/testVectorVector for reference.
type TestVectorVector struct {
	// Value field of TestVectorVector.
	Value [][]string
}

// TestVectorVectorTypeID is TL type id of TestVectorVector.
const TestVectorVectorTypeID = 0x69e8846c

func (t *TestVectorVector) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Value == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TestVectorVector) String() string {
	if t == nil {
		return "TestVectorVector(nil)"
	}
	type Alias TestVectorVector
	return fmt.Sprintf("TestVectorVector%+v", Alias(*t))
}

// FillFrom fills TestVectorVector from given interface.
func (t *TestVectorVector) FillFrom(from interface {
	GetValue() (value [][]string)
}) {
	t.Value = from.GetValue()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TestVectorVector) TypeID() uint32 {
	return TestVectorVectorTypeID
}

// TypeName returns name of type in TL schema.
func (*TestVectorVector) TypeName() string {
	return "testVectorVector"
}

// TypeInfo returns info about TL type.
func (t *TestVectorVector) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "testVectorVector",
		ID:   TestVectorVectorTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Value",
			SchemaName: "value",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TestVectorVector) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode %s as nil", "testVectorVector#69e8846c")
	}
	b.PutID(TestVectorVectorTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TestVectorVector) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode %s as nil", "testVectorVector#69e8846c")
	}
	b.PutInt(len(t.Value))
	for _, row := range t.Value {
		b.PutVectorHeader(len(row))
		for _, v := range row {
			b.PutString(v)
		}
	}
	return nil
}

// GetValue returns value of Value field.
func (t *TestVectorVector) GetValue() (value [][]string) {
	return t.Value
}

// Decode implements bin.Decoder.
func (t *TestVectorVector) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode %s to nil", "testVectorVector#69e8846c")
	}
	if err := b.ConsumeID(TestVectorVectorTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "testVectorVector#69e8846c", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TestVectorVector) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode %s to nil", "testVectorVector#69e8846c")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "testVectorVector#69e8846c", "value", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			innerLen, err := b.VectorHeader()
			if err != nil {
				return fmt.Errorf("unable to decode %s: field %s: %w", "testVectorVector#69e8846c", "value", err)
			}
			var row []string
			for innerIndex := 0; innerIndex < innerLen; innerLen++ {
				value, err := b.String()
				if err != nil {
					return fmt.Errorf("unable to decode %s: field %s: %w", "testVectorVector#69e8846c", "value", err)
				}
				row = append(row, value)
			}
			t.Value = append(t.Value, row)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for TestVectorVector.
var (
	_ bin.Encoder     = &TestVectorVector{}
	_ bin.Decoder     = &TestVectorVector{}
	_ bin.BareEncoder = &TestVectorVector{}
	_ bin.BareDecoder = &TestVectorVector{}
)
