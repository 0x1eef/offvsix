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
	help       bool
	pkgVersion string
	file       string
)

func main() {
	var err error
	args := flag.Args()
	if file == "" {
		err = save(args[0], pkgVersion)
	} else {
		err = saveAll(file, pkgVersion)
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
	pkg := fmt.Sprintf("%s-%s.vsix", extid, version)
	_, err = os.Stat(pkg)
	if err == nil {
		say("extension %q already exists on disk", pkg)
		return nil
	}
	say("download version %q", version)
	body, length, err := gallery.DownloadExtension(ext, version)
	if err != nil {
		return err
	}
	defer body.Close()
	result, err := read(body, length)
	if err != nil {
		return err
	}
	fmt.Println()
	err = os.WriteFile(pkg, result, 0644)
	if err != nil {
		return err
	}
	say("save %q", pkg)
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
	lines := strings.Split(string(buf), "\n")
	for index, line := range lines {
		extid := strings.TrimSpace(line)
		if index > 0 && index < len(lines)-1 {
			fmt.Println()
		}
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

func read(res io.ReadCloser, len int64) ([]byte, error) {
	result := make([]byte, len)
	total := 0
	width := 50
	for {
		n, err := res.Read(result)
		total += n
		if err == io.EOF || n == 0 {
			break
		} else if err != nil {
			return result, err
		}
		percent := float64(total) / float64(len)
		arrows := int(percent * float64(width))
		fmt.Printf(
			"\033[0K\roffvsix: [%s>%s] %d%%",
			strings.Repeat("=", arrows),
			strings.Repeat(" ", width-arrows),
			int(percent*100),
		)
	}
	return result, nil
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
	flag.StringVar(&pkgVersion, "v", "", "Set extension version")
	flag.StringVar(&file, "f", "", "Path to a text file with extensions to download, one per line")
	flag.BoolVar(&help, "h", false, "Show help")
	flag.Parse()
	if file == "" && len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "offvsix: provide either an extension or a file containing extensions")
		os.Exit(1)
	} else if help {
		showHelp()
		os.Exit(0)
	}
}
