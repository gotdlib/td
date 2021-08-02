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

// MessagesSearchGlobalRequest represents TL type `messages.searchGlobal#4bc6589a`.
// Search for messages and peers globally
//
// See https://core.telegram.org/method/messages.searchGlobal for reference.
type MessagesSearchGlobalRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Peer folder ID, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders#peer-folders
	//
	// Use SetFolderID and GetFolderID helpers.
	FolderID int
	// Query
	Q string
	// Global search filter
	Filter MessagesFilterClass
	// If a positive value was specified, the method will return only messages with date
	// bigger than min_date
	MinDate int
	// If a positive value was transferred, the method will return only messages with date
	// smaller than max_date
	MaxDate int
	// Initially 0, then set to the next_rate parameter of messages.messagesSlice¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/messages.messagesSlice
	OffsetRate int
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetPeer InputPeerClass
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetID int
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
}

// MessagesSearchGlobalRequestTypeID is TL type id of MessagesSearchGlobalRequest.
const MessagesSearchGlobalRequestTypeID = 0x4bc6589a

func (s *MessagesSearchGlobalRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.FolderID == 0) {
		return false
	}
	if !(s.Q == "") {
		return false
	}
	if !(s.Filter == nil) {
		return false
	}
	if !(s.MinDate == 0) {
		return false
	}
	if !(s.MaxDate == 0) {
		return false
	}
	if !(s.OffsetRate == 0) {
		return false
	}
	if !(s.OffsetPeer == nil) {
		return false
	}
	if !(s.OffsetID == 0) {
		return false
	}
	if !(s.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSearchGlobalRequest) String() string {
	if s == nil {
		return "MessagesSearchGlobalRequest(nil)"
	}
	type Alias MessagesSearchGlobalRequest
	return fmt.Sprintf("MessagesSearchGlobalRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSearchGlobalRequest from given interface.
func (s *MessagesSearchGlobalRequest) FillFrom(from interface {
	GetFolderID() (value int, ok bool)
	GetQ() (value string)
	GetFilter() (value MessagesFilterClass)
	GetMinDate() (value int)
	GetMaxDate() (value int)
	GetOffsetRate() (value int)
	GetOffsetPeer() (value InputPeerClass)
	GetOffsetID() (value int)
	GetLimit() (value int)
}) {
	if val, ok := from.GetFolderID(); ok {
		s.FolderID = val
	}

	s.Q = from.GetQ()
	s.Filter = from.GetFilter()
	s.MinDate = from.GetMinDate()
	s.MaxDate = from.GetMaxDate()
	s.OffsetRate = from.GetOffsetRate()
	s.OffsetPeer = from.GetOffsetPeer()
	s.OffsetID = from.GetOffsetID()
	s.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSearchGlobalRequest) TypeID() uint32 {
	return MessagesSearchGlobalRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSearchGlobalRequest) TypeName() string {
	return "messages.searchGlobal"
}

// TypeInfo returns info about TL type.
func (s *MessagesSearchGlobalRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.searchGlobal",
		ID:   MessagesSearchGlobalRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FolderID",
			SchemaName: "folder_id",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Q",
			SchemaName: "q",
		},
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
		{
			Name:       "MinDate",
			SchemaName: "min_date",
		},
		{
			Name:       "MaxDate",
			SchemaName: "max_date",
		},
		{
			Name:       "OffsetRate",
			SchemaName: "offset_rate",
		},
		{
			Name:       "OffsetPeer",
			SchemaName: "offset_peer",
		},
		{
			Name:       "OffsetID",
			SchemaName: "offset_id",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSearchGlobalRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode %s as nil", "messages.searchGlobal#4bc6589a")
	}
	b.PutID(MessagesSearchGlobalRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSearchGlobalRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode %s as nil", "messages.searchGlobal#4bc6589a")
	}
	if !(s.FolderID == 0) {
		s.Flags.Set(0)
	}
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "messages.searchGlobal#4bc6589a", "flags", err)
	}
	if s.Flags.Has(0) {
		b.PutInt(s.FolderID)
	}
	b.PutString(s.Q)
	if s.Filter == nil {
		return fmt.Errorf("unable to encode %s: field %s is nil", "messages.searchGlobal#4bc6589a", "filter")
	}
	if err := s.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "messages.searchGlobal#4bc6589a", "filter", err)
	}
	b.PutInt(s.MinDate)
	b.PutInt(s.MaxDate)
	b.PutInt(s.OffsetRate)
	if s.OffsetPeer == nil {
		return fmt.Errorf("unable to encode %s: field %s is nil", "messages.searchGlobal#4bc6589a", "offset_peer")
	}
	if err := s.OffsetPeer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "messages.searchGlobal#4bc6589a", "offset_peer", err)
	}
	b.PutInt(s.OffsetID)
	b.PutInt(s.Limit)
	return nil
}

// SetFolderID sets value of FolderID conditional field.
func (s *MessagesSearchGlobalRequest) SetFolderID(value int) {
	s.Flags.Set(0)
	s.FolderID = value
}

// GetFolderID returns value of FolderID conditional field and
// boolean which is true if field was set.
func (s *MessagesSearchGlobalRequest) GetFolderID() (value int, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.FolderID, true
}

// GetQ returns value of Q field.
func (s *MessagesSearchGlobalRequest) GetQ() (value string) {
	return s.Q
}

// GetFilter returns value of Filter field.
func (s *MessagesSearchGlobalRequest) GetFilter() (value MessagesFilterClass) {
	return s.Filter
}

// GetMinDate returns value of MinDate field.
func (s *MessagesSearchGlobalRequest) GetMinDate() (value int) {
	return s.MinDate
}

// GetMaxDate returns value of MaxDate field.
func (s *MessagesSearchGlobalRequest) GetMaxDate() (value int) {
	return s.MaxDate
}

// GetOffsetRate returns value of OffsetRate field.
func (s *MessagesSearchGlobalRequest) GetOffsetRate() (value int) {
	return s.OffsetRate
}

// GetOffsetPeer returns value of OffsetPeer field.
func (s *MessagesSearchGlobalRequest) GetOffsetPeer() (value InputPeerClass) {
	return s.OffsetPeer
}

// GetOffsetID returns value of OffsetID field.
func (s *MessagesSearchGlobalRequest) GetOffsetID() (value int) {
	return s.OffsetID
}

// GetLimit returns value of Limit field.
func (s *MessagesSearchGlobalRequest) GetLimit() (value int) {
	return s.Limit
}

// Decode implements bin.Decoder.
func (s *MessagesSearchGlobalRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode %s to nil", "messages.searchGlobal#4bc6589a")
	}
	if err := b.ConsumeID(MessagesSearchGlobalRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "messages.searchGlobal#4bc6589a", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSearchGlobalRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode %s to nil", "messages.searchGlobal#4bc6589a")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.searchGlobal#4bc6589a", "flags", err)
		}
	}
	if s.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.searchGlobal#4bc6589a", "folder_id", err)
		}
		s.FolderID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.searchGlobal#4bc6589a", "q", err)
		}
		s.Q = value
	}
	{
		value, err := DecodeMessagesFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.searchGlobal#4bc6589a", "filter", err)
		}
		s.Filter = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.searchGlobal#4bc6589a", "min_date", err)
		}
		s.MinDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.searchGlobal#4bc6589a", "max_date", err)
		}
		s.MaxDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.searchGlobal#4bc6589a", "offset_rate", err)
		}
		s.OffsetRate = value
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.searchGlobal#4bc6589a", "offset_peer", err)
		}
		s.OffsetPeer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.searchGlobal#4bc6589a", "offset_id", err)
		}
		s.OffsetID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.searchGlobal#4bc6589a", "limit", err)
		}
		s.Limit = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSearchGlobalRequest.
var (
	_ bin.Encoder     = &MessagesSearchGlobalRequest{}
	_ bin.Decoder     = &MessagesSearchGlobalRequest{}
	_ bin.BareEncoder = &MessagesSearchGlobalRequest{}
	_ bin.BareDecoder = &MessagesSearchGlobalRequest{}
)

// MessagesSearchGlobal invokes method messages.searchGlobal#4bc6589a returning error if any.
// Search for messages and peers globally
//
// Possible errors:
//  400 FOLDER_ID_INVALID: Invalid folder ID
//  400 SEARCH_QUERY_EMPTY: The search query is empty
//
// See https://core.telegram.org/method/messages.searchGlobal for reference.
func (c *Client) MessagesSearchGlobal(ctx context.Context, request *MessagesSearchGlobalRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
