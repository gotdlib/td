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

// User represents TL type `user#598c3933`.
type User struct {
	// User identifier
	ID int64
	// First name of the user
	FirstName string
	// Last name of the user
	LastName string
	// Usernames of the user; may be null
	Usernames Usernames
	// Phone number of the user
	PhoneNumber string
	// Current online status of the user
	Status UserStatusClass
	// Profile photo of the user; may be null
	ProfilePhoto ProfilePhoto
	// Identifier of the accent color for name, and backgrounds of profile photo, reply
	// header, and link preview
	AccentColorID int32
	// Identifier of a custom emoji to be shown on the reply header and link preview
	// background; 0 if none
	BackgroundCustomEmojiID int64
	// Identifier of the accent color for the user's profile; -1 if none
	ProfileAccentColorID int32
	// Identifier of a custom emoji to be shown on the background of the user's profile; 0 if
	// none
	ProfileBackgroundCustomEmojiID int64
	// Emoji status to be shown instead of the default Telegram Premium badge; may be null
	EmojiStatus EmojiStatus
	// The user is a contact of the current user
	IsContact bool
	// The user is a contact of the current user and the current user is a contact of the
	// user
	IsMutualContact bool
	// The user is a close friend of the current user; implies that the user is a contact
	IsCloseFriend bool
	// Information about verification status of the user; may be null if none
	VerificationStatus VerificationStatus
	// True, if the user is a Telegram Premium user
	IsPremium bool
	// True, if the user is Telegram support account
	IsSupport bool
	// If non-empty, it contains a human-readable description of the reason why access to
	// this user must be restricted
	RestrictionReason string
	// True, if the user has non-expired stories available to the current user
	HasActiveStories bool
	// True, if the user has unread non-expired stories available to the current user
	HasUnreadActiveStories bool
	// True, if the user may restrict new chats with non-contacts. Use canSendMessageToUser
	// to check whether the current user can message the user or try to create a chat with
	// them
	RestrictsNewChats bool
	// If false, the user is inaccessible, and the only information known about the user is
	// inside this class. Identifier of the user can't be passed to any method
	HaveAccess bool
	// Type of the user
	Type UserTypeClass
	// IETF language tag of the user's language; only available to bots
	LanguageCode string
	// True, if the user added the current bot to attachment menu; only available to bots
	AddedToAttachmentMenu bool
}

// UserTypeID is TL type id of User.
const UserTypeID = 0x598c3933

// Ensuring interfaces in compile-time for User.
var (
	_ bin.Encoder     = &User{}
	_ bin.Decoder     = &User{}
	_ bin.BareEncoder = &User{}
	_ bin.BareDecoder = &User{}
)

func (u *User) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.ID == 0) {
		return false
	}
	if !(u.FirstName == "") {
		return false
	}
	if !(u.LastName == "") {
		return false
	}
	if !(u.Usernames.Zero()) {
		return false
	}
	if !(u.PhoneNumber == "") {
		return false
	}
	if !(u.Status == nil) {
		return false
	}
	if !(u.ProfilePhoto.Zero()) {
		return false
	}
	if !(u.AccentColorID == 0) {
		return false
	}
	if !(u.BackgroundCustomEmojiID == 0) {
		return false
	}
	if !(u.ProfileAccentColorID == 0) {
		return false
	}
	if !(u.ProfileBackgroundCustomEmojiID == 0) {
		return false
	}
	if !(u.EmojiStatus.Zero()) {
		return false
	}
	if !(u.IsContact == false) {
		return false
	}
	if !(u.IsMutualContact == false) {
		return false
	}
	if !(u.IsCloseFriend == false) {
		return false
	}
	if !(u.VerificationStatus.Zero()) {
		return false
	}
	if !(u.IsPremium == false) {
		return false
	}
	if !(u.IsSupport == false) {
		return false
	}
	if !(u.RestrictionReason == "") {
		return false
	}
	if !(u.HasActiveStories == false) {
		return false
	}
	if !(u.HasUnreadActiveStories == false) {
		return false
	}
	if !(u.RestrictsNewChats == false) {
		return false
	}
	if !(u.HaveAccess == false) {
		return false
	}
	if !(u.Type == nil) {
		return false
	}
	if !(u.LanguageCode == "") {
		return false
	}
	if !(u.AddedToAttachmentMenu == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *User) String() string {
	if u == nil {
		return "User(nil)"
	}
	type Alias User
	return fmt.Sprintf("User%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*User) TypeID() uint32 {
	return UserTypeID
}

// TypeName returns name of type in TL schema.
func (*User) TypeName() string {
	return "user"
}

// TypeInfo returns info about TL type.
func (u *User) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "user",
		ID:   UserTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "FirstName",
			SchemaName: "first_name",
		},
		{
			Name:       "LastName",
			SchemaName: "last_name",
		},
		{
			Name:       "Usernames",
			SchemaName: "usernames",
		},
		{
			Name:       "PhoneNumber",
			SchemaName: "phone_number",
		},
		{
			Name:       "Status",
			SchemaName: "status",
		},
		{
			Name:       "ProfilePhoto",
			SchemaName: "profile_photo",
		},
		{
			Name:       "AccentColorID",
			SchemaName: "accent_color_id",
		},
		{
			Name:       "BackgroundCustomEmojiID",
			SchemaName: "background_custom_emoji_id",
		},
		{
			Name:       "ProfileAccentColorID",
			SchemaName: "profile_accent_color_id",
		},
		{
			Name:       "ProfileBackgroundCustomEmojiID",
			SchemaName: "profile_background_custom_emoji_id",
		},
		{
			Name:       "EmojiStatus",
			SchemaName: "emoji_status",
		},
		{
			Name:       "IsContact",
			SchemaName: "is_contact",
		},
		{
			Name:       "IsMutualContact",
			SchemaName: "is_mutual_contact",
		},
		{
			Name:       "IsCloseFriend",
			SchemaName: "is_close_friend",
		},
		{
			Name:       "VerificationStatus",
			SchemaName: "verification_status",
		},
		{
			Name:       "IsPremium",
			SchemaName: "is_premium",
		},
		{
			Name:       "IsSupport",
			SchemaName: "is_support",
		},
		{
			Name:       "RestrictionReason",
			SchemaName: "restriction_reason",
		},
		{
			Name:       "HasActiveStories",
			SchemaName: "has_active_stories",
		},
		{
			Name:       "HasUnreadActiveStories",
			SchemaName: "has_unread_active_stories",
		},
		{
			Name:       "RestrictsNewChats",
			SchemaName: "restricts_new_chats",
		},
		{
			Name:       "HaveAccess",
			SchemaName: "have_access",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "LanguageCode",
			SchemaName: "language_code",
		},
		{
			Name:       "AddedToAttachmentMenu",
			SchemaName: "added_to_attachment_menu",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *User) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode user#598c3933 as nil")
	}
	b.PutID(UserTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *User) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode user#598c3933 as nil")
	}
	b.PutInt53(u.ID)
	b.PutString(u.FirstName)
	b.PutString(u.LastName)
	if err := u.Usernames.Encode(b); err != nil {
		return fmt.Errorf("unable to encode user#598c3933: field usernames: %w", err)
	}
	b.PutString(u.PhoneNumber)
	if u.Status == nil {
		return fmt.Errorf("unable to encode user#598c3933: field status is nil")
	}
	if err := u.Status.Encode(b); err != nil {
		return fmt.Errorf("unable to encode user#598c3933: field status: %w", err)
	}
	if err := u.ProfilePhoto.Encode(b); err != nil {
		return fmt.Errorf("unable to encode user#598c3933: field profile_photo: %w", err)
	}
	b.PutInt32(u.AccentColorID)
	b.PutLong(u.BackgroundCustomEmojiID)
	b.PutInt32(u.ProfileAccentColorID)
	b.PutLong(u.ProfileBackgroundCustomEmojiID)
	if err := u.EmojiStatus.Encode(b); err != nil {
		return fmt.Errorf("unable to encode user#598c3933: field emoji_status: %w", err)
	}
	b.PutBool(u.IsContact)
	b.PutBool(u.IsMutualContact)
	b.PutBool(u.IsCloseFriend)
	if err := u.VerificationStatus.Encode(b); err != nil {
		return fmt.Errorf("unable to encode user#598c3933: field verification_status: %w", err)
	}
	b.PutBool(u.IsPremium)
	b.PutBool(u.IsSupport)
	b.PutString(u.RestrictionReason)
	b.PutBool(u.HasActiveStories)
	b.PutBool(u.HasUnreadActiveStories)
	b.PutBool(u.RestrictsNewChats)
	b.PutBool(u.HaveAccess)
	if u.Type == nil {
		return fmt.Errorf("unable to encode user#598c3933: field type is nil")
	}
	if err := u.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode user#598c3933: field type: %w", err)
	}
	b.PutString(u.LanguageCode)
	b.PutBool(u.AddedToAttachmentMenu)
	return nil
}

// Decode implements bin.Decoder.
func (u *User) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode user#598c3933 to nil")
	}
	if err := b.ConsumeID(UserTypeID); err != nil {
		return fmt.Errorf("unable to decode user#598c3933: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *User) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode user#598c3933 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field id: %w", err)
		}
		u.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field first_name: %w", err)
		}
		u.FirstName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field last_name: %w", err)
		}
		u.LastName = value
	}
	{
		if err := u.Usernames.Decode(b); err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field usernames: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field phone_number: %w", err)
		}
		u.PhoneNumber = value
	}
	{
		value, err := DecodeUserStatus(b)
		if err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field status: %w", err)
		}
		u.Status = value
	}
	{
		if err := u.ProfilePhoto.Decode(b); err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field profile_photo: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field accent_color_id: %w", err)
		}
		u.AccentColorID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field background_custom_emoji_id: %w", err)
		}
		u.BackgroundCustomEmojiID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field profile_accent_color_id: %w", err)
		}
		u.ProfileAccentColorID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field profile_background_custom_emoji_id: %w", err)
		}
		u.ProfileBackgroundCustomEmojiID = value
	}
	{
		if err := u.EmojiStatus.Decode(b); err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field emoji_status: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field is_contact: %w", err)
		}
		u.IsContact = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field is_mutual_contact: %w", err)
		}
		u.IsMutualContact = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field is_close_friend: %w", err)
		}
		u.IsCloseFriend = value
	}
	{
		if err := u.VerificationStatus.Decode(b); err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field verification_status: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field is_premium: %w", err)
		}
		u.IsPremium = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field is_support: %w", err)
		}
		u.IsSupport = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field restriction_reason: %w", err)
		}
		u.RestrictionReason = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field has_active_stories: %w", err)
		}
		u.HasActiveStories = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field has_unread_active_stories: %w", err)
		}
		u.HasUnreadActiveStories = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field restricts_new_chats: %w", err)
		}
		u.RestrictsNewChats = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field have_access: %w", err)
		}
		u.HaveAccess = value
	}
	{
		value, err := DecodeUserType(b)
		if err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field type: %w", err)
		}
		u.Type = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field language_code: %w", err)
		}
		u.LanguageCode = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode user#598c3933: field added_to_attachment_menu: %w", err)
		}
		u.AddedToAttachmentMenu = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (u *User) EncodeTDLibJSON(b tdjson.Encoder) error {
	if u == nil {
		return fmt.Errorf("can't encode user#598c3933 as nil")
	}
	b.ObjStart()
	b.PutID("user")
	b.Comma()
	b.FieldStart("id")
	b.PutInt53(u.ID)
	b.Comma()
	b.FieldStart("first_name")
	b.PutString(u.FirstName)
	b.Comma()
	b.FieldStart("last_name")
	b.PutString(u.LastName)
	b.Comma()
	b.FieldStart("usernames")
	if err := u.Usernames.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode user#598c3933: field usernames: %w", err)
	}
	b.Comma()
	b.FieldStart("phone_number")
	b.PutString(u.PhoneNumber)
	b.Comma()
	b.FieldStart("status")
	if u.Status == nil {
		return fmt.Errorf("unable to encode user#598c3933: field status is nil")
	}
	if err := u.Status.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode user#598c3933: field status: %w", err)
	}
	b.Comma()
	b.FieldStart("profile_photo")
	if err := u.ProfilePhoto.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode user#598c3933: field profile_photo: %w", err)
	}
	b.Comma()
	b.FieldStart("accent_color_id")
	b.PutInt32(u.AccentColorID)
	b.Comma()
	b.FieldStart("background_custom_emoji_id")
	b.PutLong(u.BackgroundCustomEmojiID)
	b.Comma()
	b.FieldStart("profile_accent_color_id")
	b.PutInt32(u.ProfileAccentColorID)
	b.Comma()
	b.FieldStart("profile_background_custom_emoji_id")
	b.PutLong(u.ProfileBackgroundCustomEmojiID)
	b.Comma()
	b.FieldStart("emoji_status")
	if err := u.EmojiStatus.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode user#598c3933: field emoji_status: %w", err)
	}
	b.Comma()
	b.FieldStart("is_contact")
	b.PutBool(u.IsContact)
	b.Comma()
	b.FieldStart("is_mutual_contact")
	b.PutBool(u.IsMutualContact)
	b.Comma()
	b.FieldStart("is_close_friend")
	b.PutBool(u.IsCloseFriend)
	b.Comma()
	b.FieldStart("verification_status")
	if err := u.VerificationStatus.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode user#598c3933: field verification_status: %w", err)
	}
	b.Comma()
	b.FieldStart("is_premium")
	b.PutBool(u.IsPremium)
	b.Comma()
	b.FieldStart("is_support")
	b.PutBool(u.IsSupport)
	b.Comma()
	b.FieldStart("restriction_reason")
	b.PutString(u.RestrictionReason)
	b.Comma()
	b.FieldStart("has_active_stories")
	b.PutBool(u.HasActiveStories)
	b.Comma()
	b.FieldStart("has_unread_active_stories")
	b.PutBool(u.HasUnreadActiveStories)
	b.Comma()
	b.FieldStart("restricts_new_chats")
	b.PutBool(u.RestrictsNewChats)
	b.Comma()
	b.FieldStart("have_access")
	b.PutBool(u.HaveAccess)
	b.Comma()
	b.FieldStart("type")
	if u.Type == nil {
		return fmt.Errorf("unable to encode user#598c3933: field type is nil")
	}
	if err := u.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode user#598c3933: field type: %w", err)
	}
	b.Comma()
	b.FieldStart("language_code")
	b.PutString(u.LanguageCode)
	b.Comma()
	b.FieldStart("added_to_attachment_menu")
	b.PutBool(u.AddedToAttachmentMenu)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (u *User) DecodeTDLibJSON(b tdjson.Decoder) error {
	if u == nil {
		return fmt.Errorf("can't decode user#598c3933 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("user"); err != nil {
				return fmt.Errorf("unable to decode user#598c3933: %w", err)
			}
		case "id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field id: %w", err)
			}
			u.ID = value
		case "first_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field first_name: %w", err)
			}
			u.FirstName = value
		case "last_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field last_name: %w", err)
			}
			u.LastName = value
		case "usernames":
			if err := u.Usernames.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field usernames: %w", err)
			}
		case "phone_number":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field phone_number: %w", err)
			}
			u.PhoneNumber = value
		case "status":
			value, err := DecodeTDLibJSONUserStatus(b)
			if err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field status: %w", err)
			}
			u.Status = value
		case "profile_photo":
			if err := u.ProfilePhoto.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field profile_photo: %w", err)
			}
		case "accent_color_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field accent_color_id: %w", err)
			}
			u.AccentColorID = value
		case "background_custom_emoji_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field background_custom_emoji_id: %w", err)
			}
			u.BackgroundCustomEmojiID = value
		case "profile_accent_color_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field profile_accent_color_id: %w", err)
			}
			u.ProfileAccentColorID = value
		case "profile_background_custom_emoji_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field profile_background_custom_emoji_id: %w", err)
			}
			u.ProfileBackgroundCustomEmojiID = value
		case "emoji_status":
			if err := u.EmojiStatus.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field emoji_status: %w", err)
			}
		case "is_contact":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field is_contact: %w", err)
			}
			u.IsContact = value
		case "is_mutual_contact":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field is_mutual_contact: %w", err)
			}
			u.IsMutualContact = value
		case "is_close_friend":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field is_close_friend: %w", err)
			}
			u.IsCloseFriend = value
		case "verification_status":
			if err := u.VerificationStatus.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field verification_status: %w", err)
			}
		case "is_premium":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field is_premium: %w", err)
			}
			u.IsPremium = value
		case "is_support":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field is_support: %w", err)
			}
			u.IsSupport = value
		case "restriction_reason":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field restriction_reason: %w", err)
			}
			u.RestrictionReason = value
		case "has_active_stories":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field has_active_stories: %w", err)
			}
			u.HasActiveStories = value
		case "has_unread_active_stories":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field has_unread_active_stories: %w", err)
			}
			u.HasUnreadActiveStories = value
		case "restricts_new_chats":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field restricts_new_chats: %w", err)
			}
			u.RestrictsNewChats = value
		case "have_access":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field have_access: %w", err)
			}
			u.HaveAccess = value
		case "type":
			value, err := DecodeTDLibJSONUserType(b)
			if err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field type: %w", err)
			}
			u.Type = value
		case "language_code":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field language_code: %w", err)
			}
			u.LanguageCode = value
		case "added_to_attachment_menu":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode user#598c3933: field added_to_attachment_menu: %w", err)
			}
			u.AddedToAttachmentMenu = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (u *User) GetID() (value int64) {
	if u == nil {
		return
	}
	return u.ID
}

// GetFirstName returns value of FirstName field.
func (u *User) GetFirstName() (value string) {
	if u == nil {
		return
	}
	return u.FirstName
}

// GetLastName returns value of LastName field.
func (u *User) GetLastName() (value string) {
	if u == nil {
		return
	}
	return u.LastName
}

// GetUsernames returns value of Usernames field.
func (u *User) GetUsernames() (value Usernames) {
	if u == nil {
		return
	}
	return u.Usernames
}

// GetPhoneNumber returns value of PhoneNumber field.
func (u *User) GetPhoneNumber() (value string) {
	if u == nil {
		return
	}
	return u.PhoneNumber
}

// GetStatus returns value of Status field.
func (u *User) GetStatus() (value UserStatusClass) {
	if u == nil {
		return
	}
	return u.Status
}

// GetProfilePhoto returns value of ProfilePhoto field.
func (u *User) GetProfilePhoto() (value ProfilePhoto) {
	if u == nil {
		return
	}
	return u.ProfilePhoto
}

// GetAccentColorID returns value of AccentColorID field.
func (u *User) GetAccentColorID() (value int32) {
	if u == nil {
		return
	}
	return u.AccentColorID
}

// GetBackgroundCustomEmojiID returns value of BackgroundCustomEmojiID field.
func (u *User) GetBackgroundCustomEmojiID() (value int64) {
	if u == nil {
		return
	}
	return u.BackgroundCustomEmojiID
}

// GetProfileAccentColorID returns value of ProfileAccentColorID field.
func (u *User) GetProfileAccentColorID() (value int32) {
	if u == nil {
		return
	}
	return u.ProfileAccentColorID
}

// GetProfileBackgroundCustomEmojiID returns value of ProfileBackgroundCustomEmojiID field.
func (u *User) GetProfileBackgroundCustomEmojiID() (value int64) {
	if u == nil {
		return
	}
	return u.ProfileBackgroundCustomEmojiID
}

// GetEmojiStatus returns value of EmojiStatus field.
func (u *User) GetEmojiStatus() (value EmojiStatus) {
	if u == nil {
		return
	}
	return u.EmojiStatus
}

// GetIsContact returns value of IsContact field.
func (u *User) GetIsContact() (value bool) {
	if u == nil {
		return
	}
	return u.IsContact
}

// GetIsMutualContact returns value of IsMutualContact field.
func (u *User) GetIsMutualContact() (value bool) {
	if u == nil {
		return
	}
	return u.IsMutualContact
}

// GetIsCloseFriend returns value of IsCloseFriend field.
func (u *User) GetIsCloseFriend() (value bool) {
	if u == nil {
		return
	}
	return u.IsCloseFriend
}

// GetVerificationStatus returns value of VerificationStatus field.
func (u *User) GetVerificationStatus() (value VerificationStatus) {
	if u == nil {
		return
	}
	return u.VerificationStatus
}

// GetIsPremium returns value of IsPremium field.
func (u *User) GetIsPremium() (value bool) {
	if u == nil {
		return
	}
	return u.IsPremium
}

// GetIsSupport returns value of IsSupport field.
func (u *User) GetIsSupport() (value bool) {
	if u == nil {
		return
	}
	return u.IsSupport
}

// GetRestrictionReason returns value of RestrictionReason field.
func (u *User) GetRestrictionReason() (value string) {
	if u == nil {
		return
	}
	return u.RestrictionReason
}

// GetHasActiveStories returns value of HasActiveStories field.
func (u *User) GetHasActiveStories() (value bool) {
	if u == nil {
		return
	}
	return u.HasActiveStories
}

// GetHasUnreadActiveStories returns value of HasUnreadActiveStories field.
func (u *User) GetHasUnreadActiveStories() (value bool) {
	if u == nil {
		return
	}
	return u.HasUnreadActiveStories
}

// GetRestrictsNewChats returns value of RestrictsNewChats field.
func (u *User) GetRestrictsNewChats() (value bool) {
	if u == nil {
		return
	}
	return u.RestrictsNewChats
}

// GetHaveAccess returns value of HaveAccess field.
func (u *User) GetHaveAccess() (value bool) {
	if u == nil {
		return
	}
	return u.HaveAccess
}

// GetType returns value of Type field.
func (u *User) GetType() (value UserTypeClass) {
	if u == nil {
		return
	}
	return u.Type
}

// GetLanguageCode returns value of LanguageCode field.
func (u *User) GetLanguageCode() (value string) {
	if u == nil {
		return
	}
	return u.LanguageCode
}

// GetAddedToAttachmentMenu returns value of AddedToAttachmentMenu field.
func (u *User) GetAddedToAttachmentMenu() (value bool) {
	if u == nil {
		return
	}
	return u.AddedToAttachmentMenu
}
