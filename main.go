package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/jessevdk/go-flags"
	"github.com/ryuichi1208/ral2jqlog-f/lib/s3"
)

type options struct {
	DST_BUCKET string `short:"d" long:"dst-bucket" description:"audit log file" required:"false"`
	SRC_BUCKET string `short:"s" long:"src-bucket" description:"File Content Type" required:"false"`
	DATE       string `long:"date" description:"date"`
}

var DST_BUCKET, SRC_BUCKET, DATE string

func init() {
	var opts options
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if opts.DST_BUCKET == "" {
		DST_BUCKET = os.Getenv("DST_BUCKET")
	} else {
		DST_BUCKET = opts.DST_BUCKET
	}

	if opts.SRC_BUCKET == "" {
		SRC_BUCKET = os.Getenv("SRC_BUCKET")
	} else {
		SRC_BUCKET = opts.SRC_BUCKET
	}

	if DST_BUCKET == "" || SRC_BUCKET == "" {
		fmt.Println("null")
		os.Exit(1)
	}

	if opts.DATE == "" {
		const layout2 = "20060101"
		t := time.Now()
		DATE = t.Format(layout2)
	}
}

func main() {
	log.Printf("START")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	resp, err := s3.GetObjectsList(sess, DATE, SRC_BUCKET)
	if err != nil {
		os.Exit(1)
	}

	tmpDir, err := s3.MkTmpDir("audit_")
	defer func() {
		s3.RmTmpDir(tmpDir)
		log.Printf("END")
	}()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fps := s3.GetObject(sess, SRC_BUCKET, tmpDir, resp)
	for _, fp := range fps {
		s3.ReadGzip(fp)
	}

	s3.PutObject(sess, DST_BUCKET, s3.GetJsonFileList(tmpDir))
}
