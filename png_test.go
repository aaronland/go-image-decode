package decode

import (
	"context"
	"os"
	"testing"
)

func TestDecodePNG(t *testing.T) {

	ctx := context.Background()

	fh, err := os.Open("fixtures/tokyo.png")

	if err != nil {
		t.Fatal(err)
	}

	defer fh.Close()

	dec, err := NewDecoder(ctx, "png://")

	if err != nil {
		t.Fatal(err)
	}

	_, format, err := dec.Decode(ctx, fh)

	if err != nil {
		t.Fatal(err)
	}

	if format != "png" {
		t.Fatalf("Invalid format. Expected 'png' but got '%s'", format)
	}

}
