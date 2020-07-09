// Copyright 2014 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate png2compressedrgba -input text.png -output /tmp/compressedTextRGBA
//go:generate file2byteslice -input /tmp/compressedTextRGBA -output textrgba.go -package assets -var compressedTextRGBA
//go:generate gofmt -s -w .

package assets

import (
	"fmt"
	"image"
	"os"

	"golang.org/x/image/bmp"
)

const (
	imgWidth  = 256
	imgHeight = 128

	CharWidth  = 8
	CharHeight = 16
)

func CreateTextImage() image.Image {
	// s, err := gzip.NewReader(bytes.NewReader(compressedTextRGBA))
	// if err != nil {
	// 	panic(fmt.Sprintf("assets: gzip.NewReader failed: %v", err))
	// }
	// defer s.Close()
	pwd, _ := os.Getwd()
	infile, err := os.Open(pwd + "\\d2common\\d2DebugUtil\\internal\\assets\\noto_sans_mono_8x16.bmp")

	pwd = pwd
	if err != nil {
		// replace this with real error handling
		panic("ahhhh")
	}
	defer infile.Close()
	testbmp, err := bmp.Decode(infile)
	if err != nil {
		panic(fmt.Sprintf("assets: bmp.Decode failed: %v", err))
	}

	return testbmp
}
