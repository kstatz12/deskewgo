package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io/ioutil"
	"log"
	"os"
)

func download(bucket string, key string) (f * os.File) {
	sess := session.Must(session.NewSession())
	downloader := s3manager.NewDownloader(sess)
	file, err := ioutil.TempFile("tmp", "deskew")
	if err != nil {
		log.Fatal(err)
	}

	b, err := downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key: aws.String(key),
	})

	fmt.Printf("Downloaded File from %s/%s with bytes %d", bucket, key, b)
	f = file
	return f
}

