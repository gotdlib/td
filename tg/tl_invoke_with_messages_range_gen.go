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

// InvokeWithMessagesRangeRequest represents TL type `invokeWithMessagesRange#365275f2`.
// Invoke with the given message range
//
// See https://core.telegram.org/constructor/invokeWithMessagesRange for reference.
type InvokeWithMessagesRangeRequest struct {
	// Message range
	Range MessageRange
	// Query
	Query bin.Object
}

// InvokeWithMessagesRangeRequestTypeID is TL type id of InvokeWithMessagesRangeRequest.
const InvokeWithMessagesRangeRequestTypeID = 0x365275f2

func (i *InvokeWithMessagesRangeRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Range.Zero()) {
		return false
	}
	if !(i.Query == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InvokeWithMessagesRangeRequest) String() string {
	if i == nil {
		return "InvokeWithMessagesRangeRequest(nil)"
	}
	type Alias InvokeWithMessagesRangeRequest
	return fmt.Sprintf("InvokeWithMessagesRangeRequest%+v", Alias(*i))
}

// FillFrom fills InvokeWithMessagesRangeRequest from given interface.
func (i *InvokeWithMessagesRangeRequest) FillFrom(from interface {
	GetRange() (value MessageRange)
	GetQuery() (value bin.Object)
}) {
	i.Range = from.GetRange()
	i.Query = from.GetQuery()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InvokeWithMessagesRangeRequest) TypeID() uint32 {
	return InvokeWithMessagesRangeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*InvokeWithMessagesRangeRequest) TypeName() string {
	return "invokeWithMessagesRange"
}

// TypeInfo returns info about TL type.
func (i *InvokeWithMessagesRangeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "invokeWithMessagesRange",
		ID:   InvokeWithMessagesRangeRequestTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Range",
			SchemaName: "range",
		},
		{
			Name:       "Query",
			SchemaName: "query",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InvokeWithMessagesRangeRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode %s as nil", "invokeWithMessagesRange#365275f2")
	}
	b.PutID(InvokeWithMessagesRangeRequestTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InvokeWithMessagesRangeRequest) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode %s as nil", "invokeWithMessagesRange#365275f2")
	}
	if err := i.Range.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "invokeWithMessagesRange#365275f2", "range", err)
	}
	if err := i.Query.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "invokeWithMessagesRange#365275f2", "query", err)
	}
	return nil
}

// GetRange returns value of Range field.
func (i *InvokeWithMessagesRangeRequest) GetRange() (value MessageRange) {
	return i.Range
}

// GetQuery returns value of Query field.
func (i *InvokeWithMessagesRangeRequest) GetQuery() (value bin.Object) {
	return i.Query
}

// Decode implements bin.Decoder.
func (i *InvokeWithMessagesRangeRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode %s to nil", "invokeWithMessagesRange#365275f2")
	}
	if err := b.ConsumeID(InvokeWithMessagesRangeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "invokeWithMessagesRange#365275f2", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InvokeWithMessagesRangeRequest) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode %s to nil", "invokeWithMessagesRange#365275f2")
	}
	{
		if err := i.Range.Decode(b); err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "invokeWithMessagesRange#365275f2", "range", err)
		}
	}
	{
		if err := i.Query.Decode(b); err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "invokeWithMessagesRange#365275f2", "query", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for InvokeWithMessagesRangeRequest.
var (
	_ bin.Encoder     = &InvokeWithMessagesRangeRequest{}
	_ bin.Decoder     = &InvokeWithMessagesRangeRequest{}
	_ bin.BareEncoder = &InvokeWithMessagesRangeRequest{}
	_ bin.BareDecoder = &InvokeWithMessagesRangeRequest{}
)
