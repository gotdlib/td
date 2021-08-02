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

// MessagesBotResults represents TL type `messages.botResults#947ca848`.
// Result of a query to an inline bot
//
// See https://core.telegram.org/constructor/messages.botResults for reference.
type MessagesBotResults struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the result is a picture gallery
	Gallery bool
	// Query ID
	QueryID int64
	// The next offset to use when navigating through results
	//
	// Use SetNextOffset and GetNextOffset helpers.
	NextOffset string
	// Whether the bot requested the user to message him in private
	//
	// Use SetSwitchPm and GetSwitchPm helpers.
	SwitchPm InlineBotSwitchPM
	// The results
	Results []BotInlineResultClass
	// Caching validity of the results
	CacheTime int
	// Users mentioned in the results
	Users []UserClass
}

// MessagesBotResultsTypeID is TL type id of MessagesBotResults.
const MessagesBotResultsTypeID = 0x947ca848

func (b *MessagesBotResults) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Flags.Zero()) {
		return false
	}
	if !(b.Gallery == false) {
		return false
	}
	if !(b.QueryID == 0) {
		return false
	}
	if !(b.NextOffset == "") {
		return false
	}
	if !(b.SwitchPm.Zero()) {
		return false
	}
	if !(b.Results == nil) {
		return false
	}
	if !(b.CacheTime == 0) {
		return false
	}
	if !(b.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *MessagesBotResults) String() string {
	if b == nil {
		return "MessagesBotResults(nil)"
	}
	type Alias MessagesBotResults
	return fmt.Sprintf("MessagesBotResults%+v", Alias(*b))
}

// FillFrom fills MessagesBotResults from given interface.
func (b *MessagesBotResults) FillFrom(from interface {
	GetGallery() (value bool)
	GetQueryID() (value int64)
	GetNextOffset() (value string, ok bool)
	GetSwitchPm() (value InlineBotSwitchPM, ok bool)
	GetResults() (value []BotInlineResultClass)
	GetCacheTime() (value int)
	GetUsers() (value []UserClass)
}) {
	b.Gallery = from.GetGallery()
	b.QueryID = from.GetQueryID()
	if val, ok := from.GetNextOffset(); ok {
		b.NextOffset = val
	}

	if val, ok := from.GetSwitchPm(); ok {
		b.SwitchPm = val
	}

	b.Results = from.GetResults()
	b.CacheTime = from.GetCacheTime()
	b.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesBotResults) TypeID() uint32 {
	return MessagesBotResultsTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesBotResults) TypeName() string {
	return "messages.botResults"
}

// TypeInfo returns info about TL type.
func (b *MessagesBotResults) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.botResults",
		ID:   MessagesBotResultsTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Gallery",
			SchemaName: "gallery",
			Null:       !b.Flags.Has(0),
		},
		{
			Name:       "QueryID",
			SchemaName: "query_id",
		},
		{
			Name:       "NextOffset",
			SchemaName: "next_offset",
			Null:       !b.Flags.Has(1),
		},
		{
			Name:       "SwitchPm",
			SchemaName: "switch_pm",
			Null:       !b.Flags.Has(2),
		},
		{
			Name:       "Results",
			SchemaName: "results",
		},
		{
			Name:       "CacheTime",
			SchemaName: "cache_time",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *MessagesBotResults) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode %s as nil", "messages.botResults#947ca848")
	}
	buf.PutID(MessagesBotResultsTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *MessagesBotResults) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode %s as nil", "messages.botResults#947ca848")
	}
	if !(b.Gallery == false) {
		b.Flags.Set(0)
	}
	if !(b.NextOffset == "") {
		b.Flags.Set(1)
	}
	if !(b.SwitchPm.Zero()) {
		b.Flags.Set(2)
	}
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "messages.botResults#947ca848", "flags", err)
	}
	buf.PutLong(b.QueryID)
	if b.Flags.Has(1) {
		buf.PutString(b.NextOffset)
	}
	if b.Flags.Has(2) {
		if err := b.SwitchPm.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode %s: field %s: %w", "messages.botResults#947ca848", "switch_pm", err)
		}
	}
	buf.PutVectorHeader(len(b.Results))
	for idx, v := range b.Results {
		if v == nil {
			return fmt.Errorf("unable to encode %s: field %s element with index %d is nil", "messages.botResults#947ca848", "results", idx)
		}
		if err := v.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode %s: field %s element with index %d: %w", "messages.botResults#947ca848", "results", idx, err)
		}
	}
	buf.PutInt(b.CacheTime)
	buf.PutVectorHeader(len(b.Users))
	for idx, v := range b.Users {
		if v == nil {
			return fmt.Errorf("unable to encode %s: field %s element with index %d is nil", "messages.botResults#947ca848", "users", idx)
		}
		if err := v.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode %s: field %s element with index %d: %w", "messages.botResults#947ca848", "users", idx, err)
		}
	}
	return nil
}

// SetGallery sets value of Gallery conditional field.
func (b *MessagesBotResults) SetGallery(value bool) {
	if value {
		b.Flags.Set(0)
		b.Gallery = true
	} else {
		b.Flags.Unset(0)
		b.Gallery = false
	}
}

// GetGallery returns value of Gallery conditional field.
func (b *MessagesBotResults) GetGallery() (value bool) {
	return b.Flags.Has(0)
}

// GetQueryID returns value of QueryID field.
func (b *MessagesBotResults) GetQueryID() (value int64) {
	return b.QueryID
}

// SetNextOffset sets value of NextOffset conditional field.
func (b *MessagesBotResults) SetNextOffset(value string) {
	b.Flags.Set(1)
	b.NextOffset = value
}

// GetNextOffset returns value of NextOffset conditional field and
// boolean which is true if field was set.
func (b *MessagesBotResults) GetNextOffset() (value string, ok bool) {
	if !b.Flags.Has(1) {
		return value, false
	}
	return b.NextOffset, true
}

// SetSwitchPm sets value of SwitchPm conditional field.
func (b *MessagesBotResults) SetSwitchPm(value InlineBotSwitchPM) {
	b.Flags.Set(2)
	b.SwitchPm = value
}

// GetSwitchPm returns value of SwitchPm conditional field and
// boolean which is true if field was set.
func (b *MessagesBotResults) GetSwitchPm() (value InlineBotSwitchPM, ok bool) {
	if !b.Flags.Has(2) {
		return value, false
	}
	return b.SwitchPm, true
}

// GetResults returns value of Results field.
func (b *MessagesBotResults) GetResults() (value []BotInlineResultClass) {
	return b.Results
}

// MapResults returns field Results wrapped in BotInlineResultClassArray helper.
func (b *MessagesBotResults) MapResults() (value BotInlineResultClassArray) {
	return BotInlineResultClassArray(b.Results)
}

// GetCacheTime returns value of CacheTime field.
func (b *MessagesBotResults) GetCacheTime() (value int) {
	return b.CacheTime
}

// GetUsers returns value of Users field.
func (b *MessagesBotResults) GetUsers() (value []UserClass) {
	return b.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (b *MessagesBotResults) MapUsers() (value UserClassArray) {
	return UserClassArray(b.Users)
}

// Decode implements bin.Decoder.
func (b *MessagesBotResults) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode %s to nil", "messages.botResults#947ca848")
	}
	if err := buf.ConsumeID(MessagesBotResultsTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "messages.botResults#947ca848", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *MessagesBotResults) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode %s to nil", "messages.botResults#947ca848")
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.botResults#947ca848", "flags", err)
		}
	}
	b.Gallery = b.Flags.Has(0)
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.botResults#947ca848", "query_id", err)
		}
		b.QueryID = value
	}
	if b.Flags.Has(1) {
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.botResults#947ca848", "next_offset", err)
		}
		b.NextOffset = value
	}
	if b.Flags.Has(2) {
		if err := b.SwitchPm.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.botResults#947ca848", "switch_pm", err)
		}
	}
	{
		headerLen, err := buf.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.botResults#947ca848", "results", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeBotInlineResult(buf)
			if err != nil {
				return fmt.Errorf("unable to decode %s: field %s: %w", "messages.botResults#947ca848", "results", err)
			}
			b.Results = append(b.Results, value)
		}
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.botResults#947ca848", "cache_time", err)
		}
		b.CacheTime = value
	}
	{
		headerLen, err := buf.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "messages.botResults#947ca848", "users", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(buf)
			if err != nil {
				return fmt.Errorf("unable to decode %s: field %s: %w", "messages.botResults#947ca848", "users", err)
			}
			b.Users = append(b.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesBotResults.
var (
	_ bin.Encoder     = &MessagesBotResults{}
	_ bin.Decoder     = &MessagesBotResults{}
	_ bin.BareEncoder = &MessagesBotResults{}
	_ bin.BareDecoder = &MessagesBotResults{}
)
