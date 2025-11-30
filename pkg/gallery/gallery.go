package gallery

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FindExtension(extensionID string) (*Extension, error) {
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

func (ext *Extension) DownloadURL() (string, error) {
	var url string
	ver := ext.Versions[0]
	for _, f := range ver.Files {
		if f.AssetType == "Microsoft.VisualStudio.Services.VSIXPackage" {
			url = f.Source
		}
	}
	if url == "" {
		return url, fmt.Errorf("download URL not found")
	}
	return url, nil
}
