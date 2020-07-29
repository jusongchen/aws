package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	// "gocloud.dev/blob"
	// "gocloud.dev/blob/s3blob"
	"gocloud.dev/blob/s3blob"
	_ "gocloud.dev/blob/s3blob"
	// _ "gocloud.dev/blob/s3blob"
)

func main() {
	var bucketName, region, object string
	flag.StringVar(&bucketName, "backet", "cpf-merlin-uswest2-1", "the bucket to operation")
	flag.StringVar(&region, "region", "us-west-2", "the bucket region")
	flag.StringVar(&object, "object", "cpf-foo.txt", "the target object in bucket to write to")

	flag.Parse()
	if len(flag.Args()) != 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	var err error

	log.Printf("going to write to %s in bucket %s located in region %s", object, bucketName, region)
	err = write2Bucket(bucketName, region, object)
	if err != nil {
		log.Printf("write2Bucket - Error:%v", err)
		return
	}
	log.Printf("write2Bucket OK")
}

func write2Bucket(bucketName, region, object string) error {

	// Establish an AWS session.
	// See https://docs.aws.amazon.com/sdk-for-go/api/aws/session/ for more info.
	// The region must match the region for "my-bucket".
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	// Create a *blob.Bucket.
	// bucket, err := s3blob.OpenBucket(ctx, sess, "s3://cpf-merlin-uswest2-1/test", nil)
	bucket, err := s3blob.OpenBucket(ctx, sess, bucketName, nil)
	if err != nil {
		return err
	}
	defer bucket.Close()

	w, err := bucket.NewWriter(ctx, object, nil)
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
	log.Printf("successfully wrote to s3://%s/%s in region %s", bucketName, object, region)
	return nil
}
