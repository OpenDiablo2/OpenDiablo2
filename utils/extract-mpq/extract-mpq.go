package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2mpq"
)

const (
	directoryPermissions = 0750
)

func main() {
	var (
		outPath string
		verbose bool
	)

	flag.StringVar(&outPath, "o", "./output/", "output directory")
	flag.BoolVar(&verbose, "v", false, "verbose output")
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Printf("Usage: %s filename.mpq\n", os.Args[0])
		os.Exit(1)
	}

	filename := flag.Arg(0)
	mpq, err := d2mpq.Load(filename)

	if err != nil {
		log.Fatal(err)
	}

	list, err := mpq.GetFileList()
	if err != nil {
		log.Fatal(err)
	}

	_, mpqFile := filepath.Split(strings.ReplaceAll(filename, "\\", "/"))

	for _, filename := range list {
		extractFile(mpq, mpqFile, filename, outPath)

		if verbose {
			fmt.Printf("Writing: %s\n", filename)
		}
	}
}

func extractFile(mpq d2interface.Archive, mpqFile, filename, outPath string) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("recovered from panic in file: %s, %v", filename, r)
		}
	}()

	dir, file := filepath.Split(strings.ReplaceAll(filename, "\\", "/"))
	dir = mpqFile + "/" + dir

	err := os.MkdirAll(outPath+dir, directoryPermissions)
	if err != nil {
		log.Printf("failed to create directory: %s, %v", outPath+dir, err)
		return
	}

	f, err := os.Create(outPath + dir + file)
	if err != nil {
		log.Printf("failed to create file: %s, %v", filename, err)
		return
	}

	defer func() {
		_ = f.Close()
	}()

	buf, err := mpq.ReadFile(filename)
	if err != nil {
		log.Printf("failed to read file: %s, %v", filename, err)
		return
	}

	_, _ = f.Write(buf)
}
