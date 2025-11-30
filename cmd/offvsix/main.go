package main

import (
	"flag"
	"fmt"

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
	if err != nil {
		panic(err)
	}
	fmt.Println(ext)
}

func init() {
	flag.BoolVar(&help, "h", false, "Show help")
	flag.Parse()
}
