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

// PhoneGroupParticipants represents TL type `phone.groupParticipants#f47751b6`.
//
// See https://core.telegram.org/constructor/phone.groupParticipants for reference.
type PhoneGroupParticipants struct {
	// Count field of PhoneGroupParticipants.
	Count int
	// Participants field of PhoneGroupParticipants.
	Participants []GroupCallParticipant
	// NextOffset field of PhoneGroupParticipants.
	NextOffset string
	// Chats field of PhoneGroupParticipants.
	Chats []ChatClass
	// Users field of PhoneGroupParticipants.
	Users []UserClass
	// Version field of PhoneGroupParticipants.
	Version int
}

// PhoneGroupParticipantsTypeID is TL type id of PhoneGroupParticipants.
const PhoneGroupParticipantsTypeID = 0xf47751b6

func (g *PhoneGroupParticipants) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Count == 0) {
		return false
	}
	if !(g.Participants == nil) {
		return false
	}
	if !(g.NextOffset == "") {
		return false
	}
	if !(g.Chats == nil) {
		return false
	}
	if !(g.Users == nil) {
		return false
	}
	if !(g.Version == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PhoneGroupParticipants) String() string {
	if g == nil {
		return "PhoneGroupParticipants(nil)"
	}
	type Alias PhoneGroupParticipants
	return fmt.Sprintf("PhoneGroupParticipants%+v", Alias(*g))
}

// FillFrom fills PhoneGroupParticipants from given interface.
func (g *PhoneGroupParticipants) FillFrom(from interface {
	GetCount() (value int)
	GetParticipants() (value []GroupCallParticipant)
	GetNextOffset() (value string)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
	GetVersion() (value int)
}) {
	g.Count = from.GetCount()
	g.Participants = from.GetParticipants()
	g.NextOffset = from.GetNextOffset()
	g.Chats = from.GetChats()
	g.Users = from.GetUsers()
	g.Version = from.GetVersion()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneGroupParticipants) TypeID() uint32 {
	return PhoneGroupParticipantsTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneGroupParticipants) TypeName() string {
	return "phone.groupParticipants"
}

// TypeInfo returns info about TL type.
func (g *PhoneGroupParticipants) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.groupParticipants",
		ID:   PhoneGroupParticipantsTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "Participants",
			SchemaName: "participants",
		},
		{
			Name:       "NextOffset",
			SchemaName: "next_offset",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
		{
			Name:       "Version",
			SchemaName: "version",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *PhoneGroupParticipants) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode %s as nil", "phone.groupParticipants#f47751b6")
	}
	b.PutID(PhoneGroupParticipantsTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *PhoneGroupParticipants) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode %s as nil", "phone.groupParticipants#f47751b6")
	}
	b.PutInt(g.Count)
	b.PutVectorHeader(len(g.Participants))
	for idx, v := range g.Participants {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode %s: field %s element with index %d: %w", "phone.groupParticipants#f47751b6", "participants", idx, err)
		}
	}
	b.PutString(g.NextOffset)
	b.PutVectorHeader(len(g.Chats))
	for idx, v := range g.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode %s: field %s element with index %d is nil", "phone.groupParticipants#f47751b6", "chats", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode %s: field %s element with index %d: %w", "phone.groupParticipants#f47751b6", "chats", idx, err)
		}
	}
	b.PutVectorHeader(len(g.Users))
	for idx, v := range g.Users {
		if v == nil {
			return fmt.Errorf("unable to encode %s: field %s element with index %d is nil", "phone.groupParticipants#f47751b6", "users", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode %s: field %s element with index %d: %w", "phone.groupParticipants#f47751b6", "users", idx, err)
		}
	}
	b.PutInt(g.Version)
	return nil
}

// GetCount returns value of Count field.
func (g *PhoneGroupParticipants) GetCount() (value int) {
	return g.Count
}

// GetParticipants returns value of Participants field.
func (g *PhoneGroupParticipants) GetParticipants() (value []GroupCallParticipant) {
	return g.Participants
}

// GetNextOffset returns value of NextOffset field.
func (g *PhoneGroupParticipants) GetNextOffset() (value string) {
	return g.NextOffset
}

// GetChats returns value of Chats field.
func (g *PhoneGroupParticipants) GetChats() (value []ChatClass) {
	return g.Chats
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (g *PhoneGroupParticipants) MapChats() (value ChatClassArray) {
	return ChatClassArray(g.Chats)
}

// GetUsers returns value of Users field.
func (g *PhoneGroupParticipants) GetUsers() (value []UserClass) {
	return g.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (g *PhoneGroupParticipants) MapUsers() (value UserClassArray) {
	return UserClassArray(g.Users)
}

// GetVersion returns value of Version field.
func (g *PhoneGroupParticipants) GetVersion() (value int) {
	return g.Version
}

// Decode implements bin.Decoder.
func (g *PhoneGroupParticipants) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode %s to nil", "phone.groupParticipants#f47751b6")
	}
	if err := b.ConsumeID(PhoneGroupParticipantsTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "phone.groupParticipants#f47751b6", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *PhoneGroupParticipants) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode %s to nil", "phone.groupParticipants#f47751b6")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "phone.groupParticipants#f47751b6", "count", err)
		}
		g.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "phone.groupParticipants#f47751b6", "participants", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value GroupCallParticipant
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode %s: field %s: %w", "phone.groupParticipants#f47751b6", "participants", err)
			}
			g.Participants = append(g.Participants, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "phone.groupParticipants#f47751b6", "next_offset", err)
		}
		g.NextOffset = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "phone.groupParticipants#f47751b6", "chats", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode %s: field %s: %w", "phone.groupParticipants#f47751b6", "chats", err)
			}
			g.Chats = append(g.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "phone.groupParticipants#f47751b6", "users", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode %s: field %s: %w", "phone.groupParticipants#f47751b6", "users", err)
			}
			g.Users = append(g.Users, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "phone.groupParticipants#f47751b6", "version", err)
		}
		g.Version = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneGroupParticipants.
var (
	_ bin.Encoder     = &PhoneGroupParticipants{}
	_ bin.Decoder     = &PhoneGroupParticipants{}
	_ bin.BareEncoder = &PhoneGroupParticipants{}
	_ bin.BareDecoder = &PhoneGroupParticipants{}
)
