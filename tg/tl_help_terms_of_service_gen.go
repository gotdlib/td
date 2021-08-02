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

// HelpTermsOfService represents TL type `help.termsOfService#780a0310`.
// Info about the latest telegram Terms Of Service
//
// See https://core.telegram.org/constructor/help.termsOfService for reference.
type HelpTermsOfService struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether a prompt must be showed to the user, in order to accept the new terms.
	Popup bool
	// ID of the new terms
	ID DataJSON
	// Text of the new terms
	Text string
	// Message entities for styled text¹
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	Entities []MessageEntityClass
	// Minimum age required to sign up to telegram, the user must confirm that they is older
	// than the minimum age.
	//
	// Use SetMinAgeConfirm and GetMinAgeConfirm helpers.
	MinAgeConfirm int
}

// HelpTermsOfServiceTypeID is TL type id of HelpTermsOfService.
const HelpTermsOfServiceTypeID = 0x780a0310

func (t *HelpTermsOfService) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Flags.Zero()) {
		return false
	}
	if !(t.Popup == false) {
		return false
	}
	if !(t.ID.Zero()) {
		return false
	}
	if !(t.Text == "") {
		return false
	}
	if !(t.Entities == nil) {
		return false
	}
	if !(t.MinAgeConfirm == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *HelpTermsOfService) String() string {
	if t == nil {
		return "HelpTermsOfService(nil)"
	}
	type Alias HelpTermsOfService
	return fmt.Sprintf("HelpTermsOfService%+v", Alias(*t))
}

// FillFrom fills HelpTermsOfService from given interface.
func (t *HelpTermsOfService) FillFrom(from interface {
	GetPopup() (value bool)
	GetID() (value DataJSON)
	GetText() (value string)
	GetEntities() (value []MessageEntityClass)
	GetMinAgeConfirm() (value int, ok bool)
}) {
	t.Popup = from.GetPopup()
	t.ID = from.GetID()
	t.Text = from.GetText()
	t.Entities = from.GetEntities()
	if val, ok := from.GetMinAgeConfirm(); ok {
		t.MinAgeConfirm = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpTermsOfService) TypeID() uint32 {
	return HelpTermsOfServiceTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpTermsOfService) TypeName() string {
	return "help.termsOfService"
}

// TypeInfo returns info about TL type.
func (t *HelpTermsOfService) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.termsOfService",
		ID:   HelpTermsOfServiceTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Popup",
			SchemaName: "popup",
			Null:       !t.Flags.Has(0),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "Entities",
			SchemaName: "entities",
		},
		{
			Name:       "MinAgeConfirm",
			SchemaName: "min_age_confirm",
			Null:       !t.Flags.Has(1),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *HelpTermsOfService) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode %s as nil", "help.termsOfService#780a0310")
	}
	b.PutID(HelpTermsOfServiceTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *HelpTermsOfService) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode %s as nil", "help.termsOfService#780a0310")
	}
	if !(t.Popup == false) {
		t.Flags.Set(0)
	}
	if !(t.MinAgeConfirm == 0) {
		t.Flags.Set(1)
	}
	if err := t.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "help.termsOfService#780a0310", "flags", err)
	}
	if err := t.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode %s: field %s: %w", "help.termsOfService#780a0310", "id", err)
	}
	b.PutString(t.Text)
	b.PutVectorHeader(len(t.Entities))
	for idx, v := range t.Entities {
		if v == nil {
			return fmt.Errorf("unable to encode %s: field %s element with index %d is nil", "help.termsOfService#780a0310", "entities", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode %s: field %s element with index %d: %w", "help.termsOfService#780a0310", "entities", idx, err)
		}
	}
	if t.Flags.Has(1) {
		b.PutInt(t.MinAgeConfirm)
	}
	return nil
}

// SetPopup sets value of Popup conditional field.
func (t *HelpTermsOfService) SetPopup(value bool) {
	if value {
		t.Flags.Set(0)
		t.Popup = true
	} else {
		t.Flags.Unset(0)
		t.Popup = false
	}
}

// GetPopup returns value of Popup conditional field.
func (t *HelpTermsOfService) GetPopup() (value bool) {
	return t.Flags.Has(0)
}

// GetID returns value of ID field.
func (t *HelpTermsOfService) GetID() (value DataJSON) {
	return t.ID
}

// GetText returns value of Text field.
func (t *HelpTermsOfService) GetText() (value string) {
	return t.Text
}

// GetEntities returns value of Entities field.
func (t *HelpTermsOfService) GetEntities() (value []MessageEntityClass) {
	return t.Entities
}

// MapEntities returns field Entities wrapped in MessageEntityClassArray helper.
func (t *HelpTermsOfService) MapEntities() (value MessageEntityClassArray) {
	return MessageEntityClassArray(t.Entities)
}

// SetMinAgeConfirm sets value of MinAgeConfirm conditional field.
func (t *HelpTermsOfService) SetMinAgeConfirm(value int) {
	t.Flags.Set(1)
	t.MinAgeConfirm = value
}

// GetMinAgeConfirm returns value of MinAgeConfirm conditional field and
// boolean which is true if field was set.
func (t *HelpTermsOfService) GetMinAgeConfirm() (value int, ok bool) {
	if !t.Flags.Has(1) {
		return value, false
	}
	return t.MinAgeConfirm, true
}

// Decode implements bin.Decoder.
func (t *HelpTermsOfService) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode %s to nil", "help.termsOfService#780a0310")
	}
	if err := b.ConsumeID(HelpTermsOfServiceTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "help.termsOfService#780a0310", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *HelpTermsOfService) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode %s to nil", "help.termsOfService#780a0310")
	}
	{
		if err := t.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "help.termsOfService#780a0310", "flags", err)
		}
	}
	t.Popup = t.Flags.Has(0)
	{
		if err := t.ID.Decode(b); err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "help.termsOfService#780a0310", "id", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "help.termsOfService#780a0310", "text", err)
		}
		t.Text = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "help.termsOfService#780a0310", "entities", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode %s: field %s: %w", "help.termsOfService#780a0310", "entities", err)
			}
			t.Entities = append(t.Entities, value)
		}
	}
	if t.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "help.termsOfService#780a0310", "min_age_confirm", err)
		}
		t.MinAgeConfirm = value
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpTermsOfService.
var (
	_ bin.Encoder     = &HelpTermsOfService{}
	_ bin.Decoder     = &HelpTermsOfService{}
	_ bin.BareEncoder = &HelpTermsOfService{}
	_ bin.BareDecoder = &HelpTermsOfService{}
)
