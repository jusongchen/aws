package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"gocloud.dev/blob"
	"gocloud.dev/blob/s3blob"
	_ "gocloud.dev/blob/s3blob"
)

func main() {
	var err error
	err = blobOpenS3()

	if err != nil {
		log.Fatalf("blobOpenS3 - Error:%v", err)
		return
	}
	log.Printf("blobOpenS3 without session OK")

	err = s3OpenWithSession()
	if err != nil {
		log.Fatalf("s3OpenWithSession - Error:%v", err)
		return
	}
	log.Printf("s3OpenWithSession OK")
}

func s3OpenWithSession() error {

	// Establish an AWS session.
	// See https://docs.aws.amazon.com/sdk-for-go/api/aws/session/ for more info.
	// The region must match the region for "my-bucket".
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	// Create a *blob.Bucket.
	// bucket, err := s3blob.OpenBucket(ctx, sess, "s3://cpf-merlin-uswest2-1/test", nil)
	bucket, err := s3blob.OpenBucket(ctx, sess, "cpf-merlin-uswest2-1", nil)
	if err != nil {
		return err
	}
	defer bucket.Close()

	w, err := bucket.NewWriter(ctx, "foo.txt", nil)
	if err != nil {
		return err
	}
	_, writeErr := fmt.Fprintln(w, "s3blob openBucket():Greeting from Merlin!")
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

func blobOpenS3() error {

	// Establish an AWS session.
	// See https://docs.aws.amazon.com/sdk-for-go/api/aws/session/ for more info.
	// The region must match the region for "my-bucket".
	// sess, err := session.NewSession(&aws.Config{
	// 	Region: aws.String("us-west-1"),
	// })
	// if err != nil {
	// 	return err
	// }

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	// Create a *blob.Bucket.
	bucket, err := blob.OpenBucket(ctx, "s3://cpf-merlin-uswest2-1?region=us-west-2")
	// bucket, err := s3blob.OpenBucket(ctx, sess, "s3://cpf-merlin-uswest2-1/test", nil)
	if err != nil {
		return err
	}
	defer bucket.Close()

	w, err := bucket.NewWriter(ctx, "s3Spike/CentralPerfFoundations/Jusong.txt", nil)
	if err != nil {
		return err
	}
	_, writeErr := fmt.Fprintln(w, "blob.OpenBucket():Greeting from Merlin!")
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
