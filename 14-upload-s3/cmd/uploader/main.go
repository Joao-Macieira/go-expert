package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(
				"user-id",
				"user-secret",
				"",
			),
		},
	)

	if err != nil {
		panic(err)
	}

	s3Client = s3.New(sess)
	s3Bucket = "goexpert-bucket-exemple"
}

func main() {
	dir, err := os.Open("../../tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	uploadControl := make(chan struct{}, 100)
	errFileUpload := make(chan string, 10)

	go func() {
		for {
			select {
			case filename := <-errFileUpload:
				uploadControl <- struct{}{}
				wg.Add(1)
				go uploadFile(filename, uploadControl, errFileUpload)
			}
		}
	}()

	for {
		files, err := dir.ReadDir(1)

		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Printf("Error reading directory: %s\n", err)
			continue
		}

		wg.Add(1)
		uploadControl <- struct{}{}
		go uploadFile(files[0].Name(), uploadControl, errFileUpload)
	}

	wg.Wait()
}

func uploadFile(filename string, uploadControl <-chan struct{}, errFileUpload chan<- string) {
	defer wg.Done()
	completeFilename := fmt.Sprintf("../../tmp/%s", filename)
	fmt.Printf("Uploading file %s to bucket %s\n", completeFilename, s3Bucket)
	f, err := os.Open(completeFilename)

	if err != nil {
		fmt.Printf("Error reading file %s: %s\n", completeFilename, err)
		<-uploadControl // empty channel
		errFileUpload <- completeFilename
		return
	}
	defer f.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})

	if err != nil {
		fmt.Printf("Error uploading file %s: %s\n", completeFilename, err)
		<-uploadControl // empty channel
		errFileUpload <- completeFilename
		return
	}

	fmt.Printf("File %s uploaded to %s\n", completeFilename, s3Bucket)
	<-uploadControl // empty channel
}
