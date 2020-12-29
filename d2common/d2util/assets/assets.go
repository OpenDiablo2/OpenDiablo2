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

// Package assets provides files for use by the debug utils
package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"image"
	"log"

	"golang.org/x/image/bmp"
)

const (
	// CharWidth of all glyphs inside of the glyph table image
	CharWidth = 8

	// CharHeight of all glyphs inside of the glyph table image
	CharHeight = 16
)

// CreateTextImage creates
func CreateTextImage() image.Image {
	s, err := gzip.NewReader(bytes.NewReader(CompressedDebugText))
	if err != nil {
		panic(fmt.Sprintf("assets: gzip.NewReader failed: %v", err))
	}

	debugBmp, err := bmp.Decode(s)
	if err != nil {
		panic(fmt.Sprintf("assets: bmp.Decode failed: %v", err))
	}

	err = s.Close()
	if err != nil {
		log.Print(err)
	}

	return debugBmp
}
