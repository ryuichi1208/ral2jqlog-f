package s3

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func GetObjectsList(sess *session.Session, date, src string) (*s3.ListObjectsV2Output, error) {
	svc := s3.New(sess)
	key := "2022/09/06/23/"

	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(src),
		Prefix: aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func mkTmpDir(prefix string) (string, error) {
	dir, err := ioutil.TempDir("", prefix)
	if err != nil {
		return "", err
	}
	return dir, nil
}

func GetObject(sess *session.Session, src string, objs *s3.ListObjectsV2Output) []*os.File {
	var fps []*os.File
	downloader := s3manager.NewDownloader(sess)
	tmpDir, err := mkTmpDir("audit_")
	if err != nil {
		fmt.Println(tmpDir)
		return fps
	}
	for _, item := range objs.Contents {
		filename := fmt.Sprintf("%s/%s", tmpDir, filepath.Base(*item.Key))
		fmt.Println(filename)
		fp, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			fmt.Println(err)
			return fps
		}
		n, err := downloader.Download(fp, &s3.GetObjectInput{
			Bucket: aws.String(src),
			Key:    aws.String(*item.Key),
		})
		fps = append(fps, fp)
		if err != nil {
			return fps
		}
		fmt.Println(n)
	}

	return fps
}
