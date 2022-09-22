package s3

import (
	"os"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func TestGetObjectsList(t *testing.T) {
	type args struct {
		sess *session.Session
		date string
		src  string
	}
	tests := []struct {
		name    string
		args    args
		want    *s3.ListObjectsV2Output
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetObjectsList(tt.args.sess, tt.args.date, tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetObjectsList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetObjectsList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMkTmpDir(t *testing.T) {
	type args struct {
		prefix string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MkTmpDir(tt.args.prefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("MkTmpDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MkTmpDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRmTmpDir(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RmTmpDir(tt.args.dir); (err != nil) != tt.wantErr {
				t.Errorf("RmTmpDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetObject(t *testing.T) {
	type args struct {
		sess   *session.Session
		src    string
		tmpDir string
		objs   *s3.ListObjectsV2Output
	}
	tests := []struct {
		name string
		args args
		want []*os.File
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetObject(tt.args.sess, tt.args.src, tt.args.tmpDir, tt.args.objs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetJsonFileList(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetJsonFileList(tt.args.dir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetJsonFileList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeHIVEFormat(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeHIVEFormat(tt.args.filename); got != tt.want {
				t.Errorf("makeHIVEFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPutObject(t *testing.T) {
	type args struct {
		sess  *session.Session
		dst   string
		files []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PutObject(tt.args.sess, tt.args.dst, tt.args.files); (err != nil) != tt.wantErr {
				t.Errorf("PutObject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
