package decode

import (
	"context"
	"testing"
	"os"
)

func TestDecodeJPEG(t *testing.T) {

	ctx := context.Background()

	fh, err := os.Open("fixtures/tokyo.jpg")

	if err != nil {
		t.Fatal(err)
	}

	defer fh.Close()

	dec, err := NewDecoder(ctx, "jpeg://")

	if err != nil {
		t.Fatal(err)
	}

	_, format, err := dec.Decode(ctx, fh)

	if err != nil {
		t.Fatal(err)
	}

	if format != "jpeg" {
		t.Fatalf("Invalid format. Expected 'jpeg' but got '%s'", format)
	}

}
