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

// GetChatJoinRequestsRequest represents TL type `getChatJoinRequests#e8d90ea2`.
type GetChatJoinRequestsRequest struct {
	// Chat identifier
	ChatID int64
	// Invite link for which to return join requests. If empty, all join requests will be
	// returned. Requires administrator privileges and can_invite_users right in the chat for
	// own links and owner privileges for other links
	InviteLink string
	// A query to search for in the first names, last names and usernames of the users to
	// return
	Query string
	// A chat join request from which to return next requests; pass null to get results from
	// the beginning
	OffsetRequest ChatJoinRequest
	// The maximum number of chat join requests to return
	Limit int32
}

// GetChatJoinRequestsRequestTypeID is TL type id of GetChatJoinRequestsRequest.
const GetChatJoinRequestsRequestTypeID = 0xe8d90ea2

// Ensuring interfaces in compile-time for GetChatJoinRequestsRequest.
var (
	_ bin.Encoder     = &GetChatJoinRequestsRequest{}
	_ bin.Decoder     = &GetChatJoinRequestsRequest{}
	_ bin.BareEncoder = &GetChatJoinRequestsRequest{}
	_ bin.BareDecoder = &GetChatJoinRequestsRequest{}
)

func (g *GetChatJoinRequestsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.InviteLink == "") {
		return false
	}
	if !(g.Query == "") {
		return false
	}
	if !(g.OffsetRequest.Zero()) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatJoinRequestsRequest) String() string {
	if g == nil {
		return "GetChatJoinRequestsRequest(nil)"
	}
	type Alias GetChatJoinRequestsRequest
	return fmt.Sprintf("GetChatJoinRequestsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatJoinRequestsRequest) TypeID() uint32 {
	return GetChatJoinRequestsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatJoinRequestsRequest) TypeName() string {
	return "getChatJoinRequests"
}

// TypeInfo returns info about TL type.
func (g *GetChatJoinRequestsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatJoinRequests",
		ID:   GetChatJoinRequestsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "InviteLink",
			SchemaName: "invite_link",
		},
		{
			Name:       "Query",
			SchemaName: "query",
		},
		{
			Name:       "OffsetRequest",
			SchemaName: "offset_request",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatJoinRequestsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatJoinRequests#e8d90ea2 as nil")
	}
	b.PutID(GetChatJoinRequestsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatJoinRequestsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatJoinRequests#e8d90ea2 as nil")
	}
	b.PutLong(g.ChatID)
	b.PutString(g.InviteLink)
	b.PutString(g.Query)
	if err := g.OffsetRequest.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getChatJoinRequests#e8d90ea2: field offset_request: %w", err)
	}
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatJoinRequestsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatJoinRequests#e8d90ea2 to nil")
	}
	if err := b.ConsumeID(GetChatJoinRequestsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatJoinRequests#e8d90ea2: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatJoinRequestsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatJoinRequests#e8d90ea2 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getChatJoinRequests#e8d90ea2: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getChatJoinRequests#e8d90ea2: field invite_link: %w", err)
		}
		g.InviteLink = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getChatJoinRequests#e8d90ea2: field query: %w", err)
		}
		g.Query = value
	}
	{
		if err := g.OffsetRequest.Decode(b); err != nil {
			return fmt.Errorf("unable to decode getChatJoinRequests#e8d90ea2: field offset_request: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getChatJoinRequests#e8d90ea2: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatJoinRequestsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatJoinRequests#e8d90ea2 as nil")
	}
	b.ObjStart()
	b.PutID("getChatJoinRequests")
	b.FieldStart("chat_id")
	b.PutLong(g.ChatID)
	b.FieldStart("invite_link")
	b.PutString(g.InviteLink)
	b.FieldStart("query")
	b.PutString(g.Query)
	b.FieldStart("offset_request")
	if err := g.OffsetRequest.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getChatJoinRequests#e8d90ea2: field offset_request: %w", err)
	}
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatJoinRequestsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatJoinRequests#e8d90ea2 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatJoinRequests"); err != nil {
				return fmt.Errorf("unable to decode getChatJoinRequests#e8d90ea2: %w", err)
			}
		case "chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode getChatJoinRequests#e8d90ea2: field chat_id: %w", err)
			}
			g.ChatID = value
		case "invite_link":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getChatJoinRequests#e8d90ea2: field invite_link: %w", err)
			}
			g.InviteLink = value
		case "query":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getChatJoinRequests#e8d90ea2: field query: %w", err)
			}
			g.Query = value
		case "offset_request":
			if err := g.OffsetRequest.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode getChatJoinRequests#e8d90ea2: field offset_request: %w", err)
			}
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getChatJoinRequests#e8d90ea2: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatJoinRequestsRequest) GetChatID() (value int64) {
	return g.ChatID
}

// GetInviteLink returns value of InviteLink field.
func (g *GetChatJoinRequestsRequest) GetInviteLink() (value string) {
	return g.InviteLink
}

// GetQuery returns value of Query field.
func (g *GetChatJoinRequestsRequest) GetQuery() (value string) {
	return g.Query
}

// GetOffsetRequest returns value of OffsetRequest field.
func (g *GetChatJoinRequestsRequest) GetOffsetRequest() (value ChatJoinRequest) {
	return g.OffsetRequest
}

// GetLimit returns value of Limit field.
func (g *GetChatJoinRequestsRequest) GetLimit() (value int32) {
	return g.Limit
}

// GetChatJoinRequests invokes method getChatJoinRequests#e8d90ea2 returning error if any.
func (c *Client) GetChatJoinRequests(ctx context.Context, request *GetChatJoinRequestsRequest) (*ChatJoinRequests, error) {
	var result ChatJoinRequests

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
