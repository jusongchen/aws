package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"gocloud.dev/blob/s3blob"
)

func main() {
	err := s3()
	if err != nil {
		log.Fatalf("Error:%v", err)
		return
	}
	log.Printf("OK")
}

func s3() error {

	// Establish an AWS session.
	// See https://docs.aws.amazon.com/sdk-for-go/api/aws/session/ for more info.
	// The region must match the region for "my-bucket".
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-1"),
	})
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	// Create a *blob.Bucket.
	bucket, err := s3blob.OpenBucket(ctx, sess, "s3blob://my-bucket", nil)
	if err != nil {
		return err
	}
	defer bucket.Close()

	w, err := bucket.NewWriter(ctx, "foo.txt", nil)
	if err != nil {
		return err
	}
	_, writeErr := fmt.Fprintln(w, "Hello, World!")
	// Always check the return value of Close when writing.
	closeErr := w.Close()
	if writeErr != nil {
		return writeErr
	}
	if closeErr != nil {
		return closeErr
	}
	return nil
}
