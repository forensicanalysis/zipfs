// Copyright (c) 2019-2020 Siemens AG
// Copyright (c) 2019-2021 Jonas Plum
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
//
// Author(s): Jonas Plum

package zipfs

import (
	"bytes"
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/forensicanalysis/fslib/fsio"
	fslibtest "github.com/forensicanalysis/fslib/fstest"
)

func TestNewReadZipFs(t *testing.T) {
	f := bytes.NewReader([]byte{0x50, 0x4b, 0x03, 0x04, 0x14, 0x00, 0x00, 0x00, 0x08, 0x00, 0x79, 0x4d, 0x71, 0x4e, 0x0c, 0x7e, 0x7f, 0xd8, 0x10, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x74, 0x78, 0x74, 0x05, 0xc0, 0x31, 0x0d, 0x00, 0x00, 0x00, 0xc2, 0x30, 0xb5, 0x28, 0xa0, 0xf6, 0x8f, 0xc5, 0x2e, 0x50, 0x4b, 0x01, 0x02, 0x14, 0x00, 0x14, 0x00, 0x00, 0x00, 0x08, 0x00, 0x79, 0x4d, 0x71, 0x4e, 0x0c, 0x7e, 0x7f, 0xd8, 0x10, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x74, 0x78, 0x74, 0x50, 0x4b, 0x05, 0x06, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x36, 0x00, 0x00, 0x00, 0x36, 0x00, 0x00, 0x00, 0x00, 0x00})
	g := bytes.NewReader([]byte{0x00})

	type args struct {
		file fsio.ReadSeekerAt
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Create read FS", args{f}, false},
		{"Create invalid zip", args{g}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsys, err := New(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewReadZipFs() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}

			err = fstest.TestFS(fsys, "file.txt")
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func Test_ZIP(t *testing.T) {
	tests := fslibtest.GetDefaultContainerTests()

	tests["rootTest"].InfoMode = fs.ModeDir

	fslibtest.RunTest(t, "ZIP", "zipfs/testdata/zip.zip", func(f fsio.ReadSeekerAt) (fs.FS, error) { return New(f) }, tests)
}
