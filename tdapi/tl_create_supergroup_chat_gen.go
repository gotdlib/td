// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
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
	_ = tdjson.Encoder{}
)

// CreateSupergroupChatRequest represents TL type `createSupergroupChat#46c770eb`.
type CreateSupergroupChatRequest struct {
	// Supergroup or channel identifier
	SupergroupID int64
	// If true, the chat will be created without network request. In this case all
	// information about the chat except its type, title and photo can be incorrect
	Force bool
}

// CreateSupergroupChatRequestTypeID is TL type id of CreateSupergroupChatRequest.
const CreateSupergroupChatRequestTypeID = 0x46c770eb

// Ensuring interfaces in compile-time for CreateSupergroupChatRequest.
var (
	_ bin.Encoder     = &CreateSupergroupChatRequest{}
	_ bin.Decoder     = &CreateSupergroupChatRequest{}
	_ bin.BareEncoder = &CreateSupergroupChatRequest{}
	_ bin.BareDecoder = &CreateSupergroupChatRequest{}
)

func (c *CreateSupergroupChatRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.SupergroupID == 0) {
		return false
	}
	if !(c.Force == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CreateSupergroupChatRequest) String() string {
	if c == nil {
		return "CreateSupergroupChatRequest(nil)"
	}
	type Alias CreateSupergroupChatRequest
	return fmt.Sprintf("CreateSupergroupChatRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CreateSupergroupChatRequest) TypeID() uint32 {
	return CreateSupergroupChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CreateSupergroupChatRequest) TypeName() string {
	return "createSupergroupChat"
}

// TypeInfo returns info about TL type.
func (c *CreateSupergroupChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "createSupergroupChat",
		ID:   CreateSupergroupChatRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SupergroupID",
			SchemaName: "supergroup_id",
		},
		{
			Name:       "Force",
			SchemaName: "force",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CreateSupergroupChatRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode createSupergroupChat#46c770eb as nil")
	}
	b.PutID(CreateSupergroupChatRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CreateSupergroupChatRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode createSupergroupChat#46c770eb as nil")
	}
	b.PutLong(c.SupergroupID)
	b.PutBool(c.Force)
	return nil
}

// Decode implements bin.Decoder.
func (c *CreateSupergroupChatRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode createSupergroupChat#46c770eb to nil")
	}
	if err := b.ConsumeID(CreateSupergroupChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode createSupergroupChat#46c770eb: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CreateSupergroupChatRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode createSupergroupChat#46c770eb to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode createSupergroupChat#46c770eb: field supergroup_id: %w", err)
		}
		c.SupergroupID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode createSupergroupChat#46c770eb: field force: %w", err)
		}
		c.Force = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CreateSupergroupChatRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode createSupergroupChat#46c770eb as nil")
	}
	b.ObjStart()
	b.PutID("createSupergroupChat")
	b.FieldStart("supergroup_id")
	b.PutLong(c.SupergroupID)
	b.FieldStart("force")
	b.PutBool(c.Force)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CreateSupergroupChatRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode createSupergroupChat#46c770eb to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("createSupergroupChat"); err != nil {
				return fmt.Errorf("unable to decode createSupergroupChat#46c770eb: %w", err)
			}
		case "supergroup_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode createSupergroupChat#46c770eb: field supergroup_id: %w", err)
			}
			c.SupergroupID = value
		case "force":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode createSupergroupChat#46c770eb: field force: %w", err)
			}
			c.Force = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSupergroupID returns value of SupergroupID field.
func (c *CreateSupergroupChatRequest) GetSupergroupID() (value int64) {
	return c.SupergroupID
}

// GetForce returns value of Force field.
func (c *CreateSupergroupChatRequest) GetForce() (value bool) {
	return c.Force
}

// CreateSupergroupChat invokes method createSupergroupChat#46c770eb returning error if any.
func (c *Client) CreateSupergroupChat(ctx context.Context, request *CreateSupergroupChatRequest) (*Chat, error) {
	var result Chat

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
