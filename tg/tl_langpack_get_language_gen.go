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

// LangpackGetLanguageRequest represents TL type `langpack.getLanguage#6a596502`.
// Get information about a language in a localization pack
//
// See https://core.telegram.org/method/langpack.getLanguage for reference.
type LangpackGetLanguageRequest struct {
	// Language pack name
	LangPack string
	// Language code
	LangCode string
}

// LangpackGetLanguageRequestTypeID is TL type id of LangpackGetLanguageRequest.
const LangpackGetLanguageRequestTypeID = 0x6a596502

func (g *LangpackGetLanguageRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.LangPack == "") {
		return false
	}
	if !(g.LangCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *LangpackGetLanguageRequest) String() string {
	if g == nil {
		return "LangpackGetLanguageRequest(nil)"
	}
	type Alias LangpackGetLanguageRequest
	return fmt.Sprintf("LangpackGetLanguageRequest%+v", Alias(*g))
}

// FillFrom fills LangpackGetLanguageRequest from given interface.
func (g *LangpackGetLanguageRequest) FillFrom(from interface {
	GetLangPack() (value string)
	GetLangCode() (value string)
}) {
	g.LangPack = from.GetLangPack()
	g.LangCode = from.GetLangCode()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LangpackGetLanguageRequest) TypeID() uint32 {
	return LangpackGetLanguageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*LangpackGetLanguageRequest) TypeName() string {
	return "langpack.getLanguage"
}

// TypeInfo returns info about TL type.
func (g *LangpackGetLanguageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "langpack.getLanguage",
		ID:   LangpackGetLanguageRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LangPack",
			SchemaName: "lang_pack",
		},
		{
			Name:       "LangCode",
			SchemaName: "lang_code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *LangpackGetLanguageRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode %s as nil", "langpack.getLanguage#6a596502")
	}
	b.PutID(LangpackGetLanguageRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *LangpackGetLanguageRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode %s as nil", "langpack.getLanguage#6a596502")
	}
	b.PutString(g.LangPack)
	b.PutString(g.LangCode)
	return nil
}

// GetLangPack returns value of LangPack field.
func (g *LangpackGetLanguageRequest) GetLangPack() (value string) {
	return g.LangPack
}

// GetLangCode returns value of LangCode field.
func (g *LangpackGetLanguageRequest) GetLangCode() (value string) {
	return g.LangCode
}

// Decode implements bin.Decoder.
func (g *LangpackGetLanguageRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode %s to nil", "langpack.getLanguage#6a596502")
	}
	if err := b.ConsumeID(LangpackGetLanguageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "langpack.getLanguage#6a596502", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *LangpackGetLanguageRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode %s to nil", "langpack.getLanguage#6a596502")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "langpack.getLanguage#6a596502", "lang_pack", err)
		}
		g.LangPack = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "langpack.getLanguage#6a596502", "lang_code", err)
		}
		g.LangCode = value
	}
	return nil
}

// Ensuring interfaces in compile-time for LangpackGetLanguageRequest.
var (
	_ bin.Encoder     = &LangpackGetLanguageRequest{}
	_ bin.Decoder     = &LangpackGetLanguageRequest{}
	_ bin.BareEncoder = &LangpackGetLanguageRequest{}
	_ bin.BareDecoder = &LangpackGetLanguageRequest{}
)

// LangpackGetLanguage invokes method langpack.getLanguage#6a596502 returning error if any.
// Get information about a language in a localization pack
//
// See https://core.telegram.org/method/langpack.getLanguage for reference.
func (c *Client) LangpackGetLanguage(ctx context.Context, request *LangpackGetLanguageRequest) (*LangPackLanguage, error) {
	var result LangPackLanguage

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
