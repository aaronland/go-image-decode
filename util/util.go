package util

import (
	"bytes"
	"context"
	"github.com/aaronland/go-image-decode"
	"image"
	"io"
	"io/ioutil"
)

func DecodeFromReadCloser(ctx context.Context, dec decode.Decoder, r io.ReadCloser) (image.Image, string, error) {

	body, err := ioutil.ReadAll(r)

	if err != nil {
		return nil, "", err
	}

	br := bytes.NewReader(body)
	return dec.Decode(ctx, br)
}
