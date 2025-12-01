package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/0x1eef/offvsix/pkg/gallery"
)

var (
	help    bool
	version string
	file    string
)

func main() {
	var err error
	args := flag.Args()
	if file == "" {
		err = save(args[0], version)
	} else {
		err = saveAll(file, version)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "offvsix: %v\n", err)
		os.Exit(1)
	}
}

func save(extid string, version string) error {
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
	defer r.Close()
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	fmt.Printf("offvsix: save extension to disk\n")
	p := fmt.Sprintf("%s-%s.vsix", extid, version)
	err = os.WriteFile(p, b, 0644)
	if err != nil {
		return err
	}
	fmt.Printf("offvsix: extension saved to %q\n", p)
	return nil
}

func saveAll(file string, version string) error {
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
		err := save(extid, version)
		if err != nil {
			return err
		}
	}
	return nil
}

func showHelp() {
	fmt.Println("Usage: offvsix [options] extension")
	flag.PrintDefaults()
}

func init() {
	flag.StringVar(&version, "v", "", "Set extension version")
	flag.StringVar(&file, "f", "", "Path to a text file with extensions to download, one per line")
	flag.BoolVar(&help, "h", false, "Show help")
	flag.Parse()
	if file == "" && len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "offvsix: please provide either an extension or a file containing extensions")
		os.Exit(1)
	} else if help {
		showHelp()
		os.Exit(0)
	}
}
