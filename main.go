package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/jessevdk/go-flags"
	"github.com/ryuichi1208/ral2jqlog-f/lib/s3"
)

type options struct {
	DST_BUCKET string `short:"d" long:"dst-bucket" description:"audit log file" required:"false"`
	SRC_BUCKET string `short:"s" long:"src-bucket" description:"File Content Type" required:"false"`
	REGION     string `short:"r" long:"region" description:"AWS Region"`
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
		log.Println("[ERROR] SRC_BUCKET and DST_BUCKET is null")
		os.Exit(1)
	}

	if opts.DATE == "" {
		const layout2 = "20060102"
		diff := 24 * time.Hour
		t := time.Now().Add(-diff)
		DATE = t.Format(layout2)
	} else {
		DATE = opts.DATE

		if len(DATE) != 8 && len(DATE) != 10 {
			log.Println("[ERROR] invalid date format want(20000101 or 2020010101)")
			os.Exit(1)
		}
	}

	if opts.REGION != "" {
		os.Setenv("AWS_REGION", opts.REGION)
	}

	if os.Getenv("AWS_REGION") == "" {
		fmt.Println("[ERROR] SET the environment variable AWS_REGION")
		os.Exit(1)
	}
}

func Do() {
	log.Printf("START")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	ctx := context.Background()
	var cancelFn func()
	ctx, cancelFn = context.WithTimeout(ctx, 60*time.Second)

	if cancelFn != nil {
		defer cancelFn()
	}

	resp, err := s3.GetObjectsList(sess, DATE, SRC_BUCKET)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	tmpDir, err := s3.MkTmpDir("audit_")
	defer func() {
		s3.RmTmpDir(tmpDir)
		log.Printf("END")
	}()

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fps, err := s3.GetObject(sess, SRC_BUCKET, tmpDir, resp, ctx)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	} else if len(fps) == 0 {
		log.Println("file is nil")
		os.Exit(1)
	}

	for _, fp := range fps {
		s3.ReadGzip(fp)
	}

	s3.PutObject(sess, DST_BUCKET, s3.GetJsonFileList(tmpDir))
}

func main() {
	if os.Getenv("AWS_LAMBDA_RUNTIME_API") != "" {
		lambda.Start(Do)
	} else {
		Do()
	}
}
