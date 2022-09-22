package s3

import (
	"bytes"
	"os"
	"testing"
)

func Test_message2CSV(t *testing.T) {
	type args struct {
		csvString string
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := message2CSV(tt.args.csvString, w); (err != nil) != tt.wantErr {
				t.Errorf("message2CSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("message2CSV() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_auditLog2Json(t *testing.T) {
	type args struct {
		jsonString string
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := auditLog2Json(tt.args.jsonString, w); (err != nil) != tt.wantErr {
				t.Errorf("auditLog2Json() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("auditLog2Json() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestReadGzip(t *testing.T) {
	type args struct {
		fp *os.File
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
			if err := ReadGzip(tt.args.fp); (err != nil) != tt.wantErr {
				t.Errorf("ReadGzip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
