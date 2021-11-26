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

// File represents TL type `file#2dad6278`.
type File struct {
	// Unique file identifier
	ID int32
	// File size, in bytes; 0 if unknown
	Size int32
	// Approximate file size in bytes in case the exact file size is unknown. Can be used to
	// show download/upload progress
	ExpectedSize int32
	// Information about the local copy of the file
	Local LocalFile
	// Information about the remote copy of the file
	Remote RemoteFile
}

// FileTypeID is TL type id of File.
const FileTypeID = 0x2dad6278

// Ensuring interfaces in compile-time for File.
var (
	_ bin.Encoder     = &File{}
	_ bin.Decoder     = &File{}
	_ bin.BareEncoder = &File{}
	_ bin.BareDecoder = &File{}
)

func (f *File) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.ID == 0) {
		return false
	}
	if !(f.Size == 0) {
		return false
	}
	if !(f.ExpectedSize == 0) {
		return false
	}
	if !(f.Local.Zero()) {
		return false
	}
	if !(f.Remote.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *File) String() string {
	if f == nil {
		return "File(nil)"
	}
	type Alias File
	return fmt.Sprintf("File%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*File) TypeID() uint32 {
	return FileTypeID
}

// TypeName returns name of type in TL schema.
func (*File) TypeName() string {
	return "file"
}

// TypeInfo returns info about TL type.
func (f *File) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "file",
		ID:   FileTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Size",
			SchemaName: "size",
		},
		{
			Name:       "ExpectedSize",
			SchemaName: "expected_size",
		},
		{
			Name:       "Local",
			SchemaName: "local",
		},
		{
			Name:       "Remote",
			SchemaName: "remote",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *File) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode file#2dad6278 as nil")
	}
	b.PutID(FileTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *File) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode file#2dad6278 as nil")
	}
	b.PutInt32(f.ID)
	b.PutInt32(f.Size)
	b.PutInt32(f.ExpectedSize)
	if err := f.Local.Encode(b); err != nil {
		return fmt.Errorf("unable to encode file#2dad6278: field local: %w", err)
	}
	if err := f.Remote.Encode(b); err != nil {
		return fmt.Errorf("unable to encode file#2dad6278: field remote: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *File) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode file#2dad6278 to nil")
	}
	if err := b.ConsumeID(FileTypeID); err != nil {
		return fmt.Errorf("unable to decode file#2dad6278: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *File) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode file#2dad6278 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode file#2dad6278: field id: %w", err)
		}
		f.ID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode file#2dad6278: field size: %w", err)
		}
		f.Size = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode file#2dad6278: field expected_size: %w", err)
		}
		f.ExpectedSize = value
	}
	{
		if err := f.Local.Decode(b); err != nil {
			return fmt.Errorf("unable to decode file#2dad6278: field local: %w", err)
		}
	}
	{
		if err := f.Remote.Decode(b); err != nil {
			return fmt.Errorf("unable to decode file#2dad6278: field remote: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (f *File) EncodeTDLibJSON(b tdjson.Encoder) error {
	if f == nil {
		return fmt.Errorf("can't encode file#2dad6278 as nil")
	}
	b.ObjStart()
	b.PutID("file")
	b.FieldStart("id")
	b.PutInt32(f.ID)
	b.FieldStart("size")
	b.PutInt32(f.Size)
	b.FieldStart("expected_size")
	b.PutInt32(f.ExpectedSize)
	b.FieldStart("local")
	if err := f.Local.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode file#2dad6278: field local: %w", err)
	}
	b.FieldStart("remote")
	if err := f.Remote.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode file#2dad6278: field remote: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (f *File) DecodeTDLibJSON(b tdjson.Decoder) error {
	if f == nil {
		return fmt.Errorf("can't decode file#2dad6278 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("file"); err != nil {
				return fmt.Errorf("unable to decode file#2dad6278: %w", err)
			}
		case "id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode file#2dad6278: field id: %w", err)
			}
			f.ID = value
		case "size":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode file#2dad6278: field size: %w", err)
			}
			f.Size = value
		case "expected_size":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode file#2dad6278: field expected_size: %w", err)
			}
			f.ExpectedSize = value
		case "local":
			if err := f.Local.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode file#2dad6278: field local: %w", err)
			}
		case "remote":
			if err := f.Remote.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode file#2dad6278: field remote: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (f *File) GetID() (value int32) {
	return f.ID
}

// GetSize returns value of Size field.
func (f *File) GetSize() (value int32) {
	return f.Size
}

// GetExpectedSize returns value of ExpectedSize field.
func (f *File) GetExpectedSize() (value int32) {
	return f.ExpectedSize
}

// GetLocal returns value of Local field.
func (f *File) GetLocal() (value LocalFile) {
	return f.Local
}

// GetRemote returns value of Remote field.
func (f *File) GetRemote() (value RemoteFile) {
	return f.Remote
}
