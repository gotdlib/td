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

// PollAnswer represents TL type `pollAnswer#6ca9c2e9`.
// A possible answer of a poll
//
// See https://core.telegram.org/constructor/pollAnswer for reference.
type PollAnswer struct {
	// Textual representation of the answer
	Text string
	// The param that has to be passed to messages.sendVote¹.
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.sendVote
	Option []byte
}

// PollAnswerTypeID is TL type id of PollAnswer.
const PollAnswerTypeID = 0x6ca9c2e9

func (p *PollAnswer) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Text == "") {
		return false
	}
	if !(p.Option == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PollAnswer) String() string {
	if p == nil {
		return "PollAnswer(nil)"
	}
	type Alias PollAnswer
	return fmt.Sprintf("PollAnswer%+v", Alias(*p))
}

// FillFrom fills PollAnswer from given interface.
func (p *PollAnswer) FillFrom(from interface {
	GetText() (value string)
	GetOption() (value []byte)
}) {
	p.Text = from.GetText()
	p.Option = from.GetOption()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PollAnswer) TypeID() uint32 {
	return PollAnswerTypeID
}

// TypeName returns name of type in TL schema.
func (*PollAnswer) TypeName() string {
	return "pollAnswer"
}

// TypeInfo returns info about TL type.
func (p *PollAnswer) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "pollAnswer",
		ID:   PollAnswerTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "Option",
			SchemaName: "option",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PollAnswer) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode %s as nil", "pollAnswer#6ca9c2e9")
	}
	b.PutID(PollAnswerTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PollAnswer) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode %s as nil", "pollAnswer#6ca9c2e9")
	}
	b.PutString(p.Text)
	b.PutBytes(p.Option)
	return nil
}

// GetText returns value of Text field.
func (p *PollAnswer) GetText() (value string) {
	return p.Text
}

// GetOption returns value of Option field.
func (p *PollAnswer) GetOption() (value []byte) {
	return p.Option
}

// Decode implements bin.Decoder.
func (p *PollAnswer) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode %s to nil", "pollAnswer#6ca9c2e9")
	}
	if err := b.ConsumeID(PollAnswerTypeID); err != nil {
		return fmt.Errorf("unable to decode %s: %w", "pollAnswer#6ca9c2e9", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PollAnswer) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode %s to nil", "pollAnswer#6ca9c2e9")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "pollAnswer#6ca9c2e9", "text", err)
		}
		p.Text = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode %s: field %s: %w", "pollAnswer#6ca9c2e9", "option", err)
		}
		p.Option = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PollAnswer.
var (
	_ bin.Encoder     = &PollAnswer{}
	_ bin.Decoder     = &PollAnswer{}
	_ bin.BareEncoder = &PollAnswer{}
	_ bin.BareDecoder = &PollAnswer{}
)
