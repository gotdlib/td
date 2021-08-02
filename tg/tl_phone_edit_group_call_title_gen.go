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

// PhoneEditGroupCallTitleRequest represents TL type `phone.editGroupCallTitle#1ca6ac0a`.
//
// See https://core.telegram.org/method/phone.editGroupCallTitle for reference.
type PhoneEditGroupCallTitleRequest struct {
	// Call field of PhoneEditGroupCallTitleRequest.
	Call InputGroupCall
	// Title field of PhoneEditGroupCallTitleRequest.
	Title string
}

// PhoneEditGroupCallTitleRequestTypeID is TL type id of PhoneEditGroupCallTitleRequest.
const PhoneEditGroupCallTitleRequestTypeID = 0x1ca6ac0a

func (e *PhoneEditGroupCallTitleRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Call.Zero()) {
		return false
	}
	if !(e.Title == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *PhoneEditGroupCallTitleRequest) String() string {
	if e == nil {
		return "PhoneEditGroupCallTitleRequest(nil)"
	}
	type Alias PhoneEditGroupCallTitleRequest
	return fmt.Sprintf("PhoneEditGroupCallTitleRequest%+v", Alias(*e))
}

// FillFrom fills PhoneEditGroupCallTitleRequest from given interface.
func (e *PhoneEditGroupCallTitleRequest) FillFrom(from interface {
	GetCall() (value InputGroupCall)
	GetTitle() (value string)
}) {
	e.Call = from.GetCall()
	e.Title = from.GetTitle()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneEditGroupCallTitleRequest) TypeID() uint32 {
	return PhoneEditGroupCallTitleRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneEditGroupCallTitleRequest) TypeName() string {
	return "phone.editGroupCallTitle"
}

// TypeInfo returns info about TL type.
func (e *PhoneEditGroupCallTitleRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.editGroupCallTitle",
		ID:   PhoneEditGroupCallTitleRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Call",
			SchemaName: "call",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *PhoneEditGroupCallTitleRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode %s as nil", "phone.editGroupCallTitle#1ca6ac0a")
	}
	b.PutID(PhoneEditGroupCallTitleRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *PhoneEditGroupCallTitleRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode %s as nil", "phone.editGroupCallTitle#1ca6ac0a")
	}
	if err := e.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "phone.editGroupCallTitle#1ca6ac0a", "call", err)
	}
	b.PutString(e.Title)
	return nil
}

// GetCall returns value of Call field.
func (e *PhoneEditGroupCallTitleRequest) GetCall() (value InputGroupCall) {
	return e.Call
}

// GetTitle returns value of Title field.
func (e *PhoneEditGroupCallTitleRequest) GetTitle() (value string) {
	return e.Title
}

// Decode implements bin.Decoder.
func (e *PhoneEditGroupCallTitleRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode %s to nil", "phone.editGroupCallTitle#1ca6ac0a")
	}
	if err := b.ConsumeID(PhoneEditGroupCallTitleRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "phone.editGroupCallTitle#1ca6ac0a", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *PhoneEditGroupCallTitleRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode %s to nil", "phone.editGroupCallTitle#1ca6ac0a")
	}
	{
		if err := e.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "phone.editGroupCallTitle#1ca6ac0a", "call", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "phone.editGroupCallTitle#1ca6ac0a", "title", err)
		}
		e.Title = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneEditGroupCallTitleRequest.
var (
	_ bin.Encoder     = &PhoneEditGroupCallTitleRequest{}
	_ bin.Decoder     = &PhoneEditGroupCallTitleRequest{}
	_ bin.BareEncoder = &PhoneEditGroupCallTitleRequest{}
	_ bin.BareDecoder = &PhoneEditGroupCallTitleRequest{}
)

// PhoneEditGroupCallTitle invokes method phone.editGroupCallTitle#1ca6ac0a returning error if any.
//
// See https://core.telegram.org/method/phone.editGroupCallTitle for reference.
func (c *Client) PhoneEditGroupCallTitle(ctx context.Context, request *PhoneEditGroupCallTitleRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
