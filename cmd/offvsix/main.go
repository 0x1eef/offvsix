package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/0x1eef/offvsix/pkg/asset"
	"github.com/0x1eef/offvsix/pkg/gallery"
)

var help bool
var version string

func main() {
	args := flag.Args()
	extid := args[0]
	ext, err := gallery.FindExtension(extid)
	check(err)
	if version == "" {
		version = ext.LatestVersion()
	}
	r, err := asset.DownloadExtension(ext, version)
	check(err)
	b, err := io.ReadAll(r)
	check(err)
	file := fmt.Sprintf("%s-%s.vsix", extid, version)
	err = os.WriteFile(file, b, 0644)
	check(err)
}

func showHelp() {
	fmt.Println("Usage: offvsix [options] extension")
	flag.PrintDefaults()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	flag.BoolVar(&help, "h", false, "Show help")
	flag.StringVar(&version, "v", "", "Set extension version")
	flag.Parse()

	if len(flag.Args()) != 1 {
		showHelp()
		os.Exit(1)
	} else if help {
		showHelp()
		os.Exit(0)
	}
}
