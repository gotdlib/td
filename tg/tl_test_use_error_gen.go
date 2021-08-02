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

// TestUseErrorRequest represents TL type `test.useError#ee75af01`.
//
// See https://core.telegram.org/method/test.useError for reference.
type TestUseErrorRequest struct {
}

// TestUseErrorRequestTypeID is TL type id of TestUseErrorRequest.
const TestUseErrorRequestTypeID = 0xee75af01

func (u *TestUseErrorRequest) Zero() bool {
	if u == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (u *TestUseErrorRequest) String() string {
	if u == nil {
		return "TestUseErrorRequest(nil)"
	}
	type Alias TestUseErrorRequest
	return fmt.Sprintf("TestUseErrorRequest%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TestUseErrorRequest) TypeID() uint32 {
	return TestUseErrorRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*TestUseErrorRequest) TypeName() string {
	return "test.useError"
}

// TypeInfo returns info about TL type.
func (u *TestUseErrorRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "test.useError",
		ID:   TestUseErrorRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (u *TestUseErrorRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode %s as nil", "test.useError#ee75af01")
	}
	b.PutID(TestUseErrorRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *TestUseErrorRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode %s as nil", "test.useError#ee75af01")
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *TestUseErrorRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode %s to nil", "test.useError#ee75af01")
	}
	if err := b.ConsumeID(TestUseErrorRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "test.useError#ee75af01", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *TestUseErrorRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode %s to nil", "test.useError#ee75af01")
	}
	return nil
}

// Ensuring interfaces in compile-time for TestUseErrorRequest.
var (
	_ bin.Encoder     = &TestUseErrorRequest{}
	_ bin.Decoder     = &TestUseErrorRequest{}
	_ bin.BareEncoder = &TestUseErrorRequest{}
	_ bin.BareDecoder = &TestUseErrorRequest{}
)

// TestUseError invokes method test.useError#ee75af01 returning error if any.
//
// See https://core.telegram.org/method/test.useError for reference.
func (c *Client) TestUseError(ctx context.Context) (*Error, error) {
	var result Error

	request := &TestUseErrorRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
