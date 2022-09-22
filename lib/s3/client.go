package s3

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetObjectsList(sess *session.Session, date, src string) (*s3.ListObjectsV2Output, error) {
	svc := s3.New(sess)
	key := "2022/09/06/"

	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(src),
		Prefix: aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetObject(sess *session.Session, src string, objs *s3.ListObjectsV2Output) {
	for _, item := range objs.Contents {
		fmt.Println("Name:", *item.Key)
	}
}
