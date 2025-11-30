package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/0x1eef/offvsix/pkg/gallery"
)

var help bool
var version string
var file string

func main() {
	args := flag.Args()
	if file == "" {
		extid := args[0]
		err := save(extid)
		check(err)
	} else {
		err := saveAll(file)
		check(err)
	}
}

func save(extid string) error {
	fmt.Printf("offvsix: find extension %q\n", extid)
	ext, err := gallery.FindExtension(extid)
	if err != nil {
		return err
	}
	if version == "" {
		version = ext.LatestVersion()
	}
	fmt.Printf("offvsix: download version %q\n", version)
	r, err := gallery.DownloadExtension(ext, version)
	if err != nil {
		return err
	}
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	fmt.Printf("offvsix: save extension to disk\n")
	file := fmt.Sprintf("%s-%s.vsix", extid, version)
	err = os.WriteFile(file, b, 0644)
	if err != nil {
		return err
	}
	fmt.Printf("offvsix: extension saved to %q\n", file)
	return nil
}

func saveAll(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	buf, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	for _, line := range strings.Split(string(buf), "\n") {
		extid := strings.TrimSpace(line)
		if extid == "" {
			continue
		}
		err := save(extid)
		if err != nil {
			return err
		} else {
			version = ""
		}
	}
	return nil
}

func showHelp() {
	fmt.Println("Usage: offvsix [options] extension")
	flag.PrintDefaults()
}

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "offvsix: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	flag.BoolVar(&help, "h", false, "Show help")
	flag.StringVar(&version, "v", "", "Set extension version")
	flag.StringVar(&file, "f", "", "Path to a text file with extensions to download, one per line")
	flag.Parse()
	if file == "" && len(flag.Args()) != 1 {
		showHelp()
		os.Exit(1)
	} else if help {
		showHelp()
		os.Exit(0)
	}
}
