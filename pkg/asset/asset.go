package asset

import (
	"fmt"
	"io"
	"net/http"

	"github.com/0x1eef/offvsix/pkg/gallery"
)

func DownloadExtension(ext *gallery.Extension, version string) (io.Reader, error) {
	endpoint := ext.DownloadURL(version)
	fmt.Println(endpoint)
	res, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("download failure: %s", res.Status)
	}
	return res.Body, nil
}

func SaveExtension(ext *gallery.Extension, r io.Reader) error {
	return fmt.Errorf("not implemented")
}
