package message

import (
	"context"
	"testing"
	"time"

	"github.com/gotd/td/tg"
)

func TestPhoto(t *testing.T) {
	ctx := context.Background()
	sender, mock := testSender(t)
	loc := &tg.InputPhoto{
		ID: 10,
	}

	expectSendMedia(&tg.InputMediaPhoto{ID: loc}, mock)
	expectSendMedia(&tg.InputMediaPhoto{
		ID:         loc,
		TTLSeconds: 10,
	}, mock)

	_, err := sender.Self().Photo(ctx, loc)
	mock.NoError(err)
	_, err = sender.Self().Media(ctx, Photo(loc).TTL(10*time.Second))
	mock.NoError(err)
}

func TestPhotoExternal(t *testing.T) {
	ctx := context.Background()
	sender, mock := testSender(t)

	expectSendMedia(&tg.InputMediaPhotoExternal{URL: "https://google.com"}, mock)
	expectSendMedia(&tg.InputMediaPhotoExternal{
		URL:        "https://github.com",
		TTLSeconds: 10,
	}, mock)

	_, err := sender.Self().PhotoExternal(ctx, "https://google.com")
	mock.NoError(err)
	_, err = sender.Self().Media(ctx, PhotoExternal("https://github.com").TTL(10*time.Second))
	mock.NoError(err)
}

func TestUploadedPhoto(t *testing.T) {
	ctx := context.Background()
	sender, mock := testSender(t)
	file := &tg.InputFile{
		ID: 10,
	}
	loc := &tg.InputDocumentFileLocation{
		ID: 10,
	}

	expectSendMedia(&tg.InputMediaUploadedPhoto{
		File: file,
	}, mock)
	expectSendMedia(&tg.InputMediaUploadedPhoto{
		File:       file,
		TTLSeconds: 10,
	}, mock)
	expectSendMedia(&tg.InputMediaUploadedPhoto{
		File: file,
		Stickers: []tg.InputDocumentClass{&tg.InputDocument{
			ID: loc.GetID(),
		}},
	}, mock)

	_, err := sender.Self().UploadedPhoto(ctx, file)
	mock.NoError(err)
	_, err = sender.Self().Media(ctx, UploadedPhoto(file).TTL(10*time.Second))
	mock.NoError(err)
	_, err = sender.Self().Media(ctx, UploadedPhoto(file).Stickers(loc))
	mock.NoError(err)
}
