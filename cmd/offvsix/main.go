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

func main() {
	args := flag.Args()
	if len(args) != 1 || help {
		fmt.Println("Usage: offvsix [options] extension")
		flag.PrintDefaults()
		return
	}
	extid := args[0]
	ext, err := gallery.FindExtension(extid)
	check(err)
	r, err := asset.DownloadExtension(ext)
	check(err)
	b, err := io.ReadAll(r)
	check(err)
	err = os.WriteFile(extid+".vsix", b, 0644)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	flag.BoolVar(&help, "h", false, "Show help")
	flag.Parse()
}
