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
	var v string
	args := flag.Args()
	if len(args) != 1 || help {
		showHelp()
		return
	}
	extid := args[0]
	ext, err := gallery.FindExtension(extid, version)
	check(err)
	r, err := asset.DownloadExtension(ext, version)
	check(err)
	b, err := io.ReadAll(r)
	check(err)
	if version == "" {
		v = ext.LatestVersion()
	} else {
		v = version
	}
	file := fmt.Sprintf("%s-%s.vsix", extid, v)
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
}
