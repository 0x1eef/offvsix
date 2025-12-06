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
	say("find extension %q", extid)
	ext, err := gallery.FindExtension(extid)
	if err != nil {
		return err
	}
	if version == "" {
		version = ext.LatestVersion()
	}
	p := fmt.Sprintf("%s-%s.vsix", extid, version)
	_, err = os.Stat(p)
	if err == nil {
		say("extension %q already exists on disk", p)
		return nil
	}
	say("download version %q", version)
	r, l, err := gallery.DownloadExtension(ext, version)
	if err != nil {
		return err
	}
	defer r.Close()
	b, err := read(r, l)
	if err != nil {
		return err
	}
	fmt.Println()
	say("save extension to disk")
	err = os.WriteFile(p, b, 0644)
	if err != nil {
		return err
	}
	say("extension saved to %q", p)
	return nil
}

func read(res io.ReadCloser, len int64) ([]byte, error) {
	b := make([]byte, len)
	t := 0
	for {
		n, err := res.Read(b)
		t += n
		if err == io.EOF || n == 0 {
			break
		} else if err != nil {
			return b, err
		}
		fmt.Printf("\033[0K\roffvsix: %d/%d bytes", t, len)
	}
	return b, nil
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

func say(m string, f ...any) {
	msg := fmt.Sprintf("offvsix: %s", m)
	fmt.Fprintf(os.Stdout, msg+"\n", f...)
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
