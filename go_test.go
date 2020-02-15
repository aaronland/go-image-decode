package decode

import (
	"context"
	"os"
	"testing"
)

func TestDecodeGIF(t *testing.T) {

	ctx := context.Background()

	fh, err := os.Open("fixtures/tokyo.gif")

	if err != nil {
		t.Fatal(err)
	}

	defer fh.Close()

	dec, err := NewDecoder(ctx, "gif://")

	if err != nil {
		t.Fatal(err)
	}

	_, format, err := dec.Decode(ctx, fh)

	if err != nil {
		t.Fatal(err)
	}

	if format != "gif" {
		t.Fatalf("Invalid format. Expected 'gif' but got '%s'", format)
	}

}
