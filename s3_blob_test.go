package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"gocloud.dev/blob"
)

func TestS3OpenBucket(t *testing.T) {

	// PRAGMA: This example is used on gocloud.dev; PRAGMA comments adjust how it is shown and can be ignored.
	// PRAGMA: On gocloud.dev, add a blank import: _ "gocloud.dev/blob/s3blob"
	// PRAGMA: On gocloud.dev, hide lines until the next blank line.
	ctx := context.Background()

	// blob.OpenBucket creates a *blob.Bucket from a URL.
	bucket, err := blob.OpenBucket(ctx, "s3://central-performance-foundation?region=us-east-2")
	require.NoError(t, err)
	if err != nil {
		log.Fatal(err)
	}
	defer bucket.Close()

	w, err := bucket.NewWriter(ctx, "foo.txt", nil)
	require.NoError(t, err)

	_, writeErr := fmt.Fprintln(w, "Hello, World!")
	require.NoError(t, writeErr)

	// Always check the return value of Close when writing.
	closeErr := w.Close()
	require.NoError(t, closeErr)

}
