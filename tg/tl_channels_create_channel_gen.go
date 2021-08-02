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

// ChannelsCreateChannelRequest represents TL type `channels.createChannel#3d5fb10f`.
// Create a supergroup/channel¹.
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.createChannel for reference.
type ChannelsCreateChannelRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to create a channel¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Broadcast bool
	// Whether to create a supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Megagroup bool
	// ForImport field of ChannelsCreateChannelRequest.
	ForImport bool
	// Channel title
	Title string
	// Channel description
	About string
	// Geogroup location
	//
	// Use SetGeoPoint and GetGeoPoint helpers.
	GeoPoint InputGeoPointClass
	// Geogroup address
	//
	// Use SetAddress and GetAddress helpers.
	Address string
}

// ChannelsCreateChannelRequestTypeID is TL type id of ChannelsCreateChannelRequest.
const ChannelsCreateChannelRequestTypeID = 0x3d5fb10f

func (c *ChannelsCreateChannelRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Broadcast == false) {
		return false
	}
	if !(c.Megagroup == false) {
		return false
	}
	if !(c.ForImport == false) {
		return false
	}
	if !(c.Title == "") {
		return false
	}
	if !(c.About == "") {
		return false
	}
	if !(c.GeoPoint == nil) {
		return false
	}
	if !(c.Address == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelsCreateChannelRequest) String() string {
	if c == nil {
		return "ChannelsCreateChannelRequest(nil)"
	}
	type Alias ChannelsCreateChannelRequest
	return fmt.Sprintf("ChannelsCreateChannelRequest%+v", Alias(*c))
}

// FillFrom fills ChannelsCreateChannelRequest from given interface.
func (c *ChannelsCreateChannelRequest) FillFrom(from interface {
	GetBroadcast() (value bool)
	GetMegagroup() (value bool)
	GetForImport() (value bool)
	GetTitle() (value string)
	GetAbout() (value string)
	GetGeoPoint() (value InputGeoPointClass, ok bool)
	GetAddress() (value string, ok bool)
}) {
	c.Broadcast = from.GetBroadcast()
	c.Megagroup = from.GetMegagroup()
	c.ForImport = from.GetForImport()
	c.Title = from.GetTitle()
	c.About = from.GetAbout()
	if val, ok := from.GetGeoPoint(); ok {
		c.GeoPoint = val
	}

	if val, ok := from.GetAddress(); ok {
		c.Address = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsCreateChannelRequest) TypeID() uint32 {
	return ChannelsCreateChannelRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsCreateChannelRequest) TypeName() string {
	return "channels.createChannel"
}

// TypeInfo returns info about TL type.
func (c *ChannelsCreateChannelRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.createChannel",
		ID:   ChannelsCreateChannelRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Broadcast",
			SchemaName: "broadcast",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "Megagroup",
			SchemaName: "megagroup",
			Null:       !c.Flags.Has(1),
		},
		{
			Name:       "ForImport",
			SchemaName: "for_import",
			Null:       !c.Flags.Has(3),
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "About",
			SchemaName: "about",
		},
		{
			Name:       "GeoPoint",
			SchemaName: "geo_point",
			Null:       !c.Flags.Has(2),
		},
		{
			Name:       "Address",
			SchemaName: "address",
			Null:       !c.Flags.Has(2),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChannelsCreateChannelRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode %s as nil", "channels.createChannel#3d5fb10f")
	}
	b.PutID(ChannelsCreateChannelRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChannelsCreateChannelRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode %s as nil", "channels.createChannel#3d5fb10f")
	}
	if !(c.Broadcast == false) {
		c.Flags.Set(0)
	}
	if !(c.Megagroup == false) {
		c.Flags.Set(1)
	}
	if !(c.ForImport == false) {
		c.Flags.Set(3)
	}
	if !(c.GeoPoint == nil) {
		c.Flags.Set(2)
	}
	if !(c.Address == "") {
		c.Flags.Set(2)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "channels.createChannel#3d5fb10f", "flags", err)
	}
	b.PutString(c.Title)
	b.PutString(c.About)
	if c.Flags.Has(2) {
		if c.GeoPoint == nil {
			return fmt.Errorf("unable to encode %s: field %s is nil", "channels.createChannel#3d5fb10f", "geo_point")
		}
		if err := c.GeoPoint.Encode(b); err != nil {
			return fmt.Errorf("unable to encode %s: field %s: %w", "channels.createChannel#3d5fb10f", "geo_point", err)
		}
	}
	if c.Flags.Has(2) {
		b.PutString(c.Address)
	}
	return nil
}

// SetBroadcast sets value of Broadcast conditional field.
func (c *ChannelsCreateChannelRequest) SetBroadcast(value bool) {
	if value {
		c.Flags.Set(0)
		c.Broadcast = true
	} else {
		c.Flags.Unset(0)
		c.Broadcast = false
	}
}

// GetBroadcast returns value of Broadcast conditional field.
func (c *ChannelsCreateChannelRequest) GetBroadcast() (value bool) {
	return c.Flags.Has(0)
}

// SetMegagroup sets value of Megagroup conditional field.
func (c *ChannelsCreateChannelRequest) SetMegagroup(value bool) {
	if value {
		c.Flags.Set(1)
		c.Megagroup = true
	} else {
		c.Flags.Unset(1)
		c.Megagroup = false
	}
}

// GetMegagroup returns value of Megagroup conditional field.
func (c *ChannelsCreateChannelRequest) GetMegagroup() (value bool) {
	return c.Flags.Has(1)
}

// SetForImport sets value of ForImport conditional field.
func (c *ChannelsCreateChannelRequest) SetForImport(value bool) {
	if value {
		c.Flags.Set(3)
		c.ForImport = true
	} else {
		c.Flags.Unset(3)
		c.ForImport = false
	}
}

// GetForImport returns value of ForImport conditional field.
func (c *ChannelsCreateChannelRequest) GetForImport() (value bool) {
	return c.Flags.Has(3)
}

// GetTitle returns value of Title field.
func (c *ChannelsCreateChannelRequest) GetTitle() (value string) {
	return c.Title
}

// GetAbout returns value of About field.
func (c *ChannelsCreateChannelRequest) GetAbout() (value string) {
	return c.About
}

// SetGeoPoint sets value of GeoPoint conditional field.
func (c *ChannelsCreateChannelRequest) SetGeoPoint(value InputGeoPointClass) {
	c.Flags.Set(2)
	c.GeoPoint = value
}

// GetGeoPoint returns value of GeoPoint conditional field and
// boolean which is true if field was set.
func (c *ChannelsCreateChannelRequest) GetGeoPoint() (value InputGeoPointClass, ok bool) {
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.GeoPoint, true
}

// GetGeoPointAsNotEmpty returns mapped value of GeoPoint conditional field and
// boolean which is true if field was set.
func (c *ChannelsCreateChannelRequest) GetGeoPointAsNotEmpty() (*InputGeoPoint, bool) {
	if value, ok := c.GetGeoPoint(); ok {
		return value.AsNotEmpty()
	}
	return nil, false
}

// SetAddress sets value of Address conditional field.
func (c *ChannelsCreateChannelRequest) SetAddress(value string) {
	c.Flags.Set(2)
	c.Address = value
}

// GetAddress returns value of Address conditional field and
// boolean which is true if field was set.
func (c *ChannelsCreateChannelRequest) GetAddress() (value string, ok bool) {
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.Address, true
}

// Decode implements bin.Decoder.
func (c *ChannelsCreateChannelRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode %s to nil", "channels.createChannel#3d5fb10f")
	}
	if err := b.ConsumeID(ChannelsCreateChannelRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "channels.createChannel#3d5fb10f", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChannelsCreateChannelRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode %s to nil", "channels.createChannel#3d5fb10f")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "channels.createChannel#3d5fb10f", "flags", err)
		}
	}
	c.Broadcast = c.Flags.Has(0)
	c.Megagroup = c.Flags.Has(1)
	c.ForImport = c.Flags.Has(3)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "channels.createChannel#3d5fb10f", "title", err)
		}
		c.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "channels.createChannel#3d5fb10f", "about", err)
		}
		c.About = value
	}
	if c.Flags.Has(2) {
		value, err := DecodeInputGeoPoint(b)
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "channels.createChannel#3d5fb10f", "geo_point", err)
		}
		c.GeoPoint = value
	}
	if c.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "channels.createChannel#3d5fb10f", "address", err)
		}
		c.Address = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsCreateChannelRequest.
var (
	_ bin.Encoder     = &ChannelsCreateChannelRequest{}
	_ bin.Decoder     = &ChannelsCreateChannelRequest{}
	_ bin.BareEncoder = &ChannelsCreateChannelRequest{}
	_ bin.BareDecoder = &ChannelsCreateChannelRequest{}
)

// ChannelsCreateChannel invokes method channels.createChannel#3d5fb10f returning error if any.
// Create a supergroup/channel¹.
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNELS_TOO_MUCH: You have joined too many channels/supergroups
//  400 CHAT_ABOUT_TOO_LONG: Chat about too long
//  400 CHAT_TITLE_EMPTY: No chat title provided
//  403 USER_RESTRICTED: You're spamreported, you can't create channels or chats.
//
// See https://core.telegram.org/method/channels.createChannel for reference.
func (c *Client) ChannelsCreateChannel(ctx context.Context, request *ChannelsCreateChannelRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
