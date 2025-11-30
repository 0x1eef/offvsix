package gallery

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FindExtension(extensionID string, version string) (*Extension, error) {
	var ver Response
	req, err := newRequest(extensionID)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resBody, &ver)
	if err != nil {
		return nil, err
	}
	if len(ver.Results) == 0 {
		return nil, fmt.Errorf("extension not found")
	}
	if len(ver.Results[0].Extensions) == 0 {
		return nil, fmt.Errorf("extension not found")
	}
	return &ver.Results[0].Extensions[0], nil
}

func (ext *Extension) DownloadURL(version string) string {
	if version == "" {
		version = ext.LatestVersion()
	}
	var (
		scheme = "https"
		host   = fmt.Sprintf("%s.gallery.vsassets.io", ext.Publisher.PublisherName)
		path   = fmt.Sprintf("/_apis/public/gallery/publisher/%s/extension/%s/%s/assetbyname/Microsoft.VisualStudio.Services.VSIXPackage", ext.Publisher.PublisherName, "go", version)
	)
	return fmt.Sprintf("%s://%s%s", scheme, host, path)
}

func (ext *Extension) LatestVersion() string {
	return ext.Versions[0].Version
}
