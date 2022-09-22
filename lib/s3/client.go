package s3

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func GetObjectsList(sess *session.Session, date, src string) (*s3.ListObjectsV2Output, error) {
	svc := s3.New(sess)
	key := fmt.Sprintf("%s/%s/%s", date[0:4], date[4:6], date[6:8])

	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(src),
		Prefix: aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func MkTmpDir(prefix string) (string, error) {
	dir, err := ioutil.TempDir("", prefix)
	if err != nil {
		return "", err
	}
	return dir, nil
}

func RmTmpDir(dir string) error {
	err := os.RemoveAll(dir)
	if err != nil {
		return err
	}
	return nil
}

func GetObject(sess *session.Session, src string, tmpDir string, objs *s3.ListObjectsV2Output, ctx context.Context) []*os.File {
	var fps []*os.File
	downloader := s3manager.NewDownloader(sess)
	for _, item := range objs.Contents {
		filename := fmt.Sprintf("%s/%s", tmpDir, filepath.Base(*item.Key))
		fp, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			fmt.Println(err)
			return fps
		}
		_, err = downloader.DownloadWithContext(ctx, fp, &s3.GetObjectInput{
			Bucket: aws.String(src),
			Key:    aws.String(*item.Key),
		})
		fps = append(fps, fp)
		if err != nil {
			return fps
		}
	}

	return fps
}

func GetJsonFileList(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths, jsons []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, GetJsonFileList(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
		if filepath.Ext(filepath.Join(dir, file.Name())) == ".json" {
			jsons = append(jsons, filepath.Join(dir, file.Name()))
		}
	}
	return jsons
}

func makeHIVEFormat(filename string) string {
	begin, end := 0, 0
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '/' {
			begin = i + 1
			break
		}
		if filename[i] == '.' {
			end = i
		}
	}
	arr := strings.Split(filename[begin:end], "-")
	hive := fmt.Sprintf("dt=%s-%s-%s-%s/%s", arr[7], arr[8], arr[9], arr[10], filename[begin:])
	return hive
}

func PutObject(sess *session.Session, dst string, files []string) error {
	uploader := s3manager.NewUploader(sess)

	for _, f := range files {
		key := makeHIVEFormat(f)
		file, err := os.Open(f)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(dst),
			Key:    aws.String(key),
			Body:   file,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
