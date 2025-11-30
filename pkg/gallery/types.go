package gallery

type Response struct {
	Results []Result `json:"results"`
}

type Result struct {
	Extensions []Extension `json:"extensions"`
}

type Extension struct {
	Publisher        Publisher   `json:"publisher"`
	ExtensionId      string      `json:"extensionId"`
	ExtensionName    string      `json:"extensionName"`
	DisplayName      string      `json:"displayName"`
	ShortDescription string      `json:"shortDescription"`
	Versions         []Version   `json:"versions"`
	Statistics       []Statistic `json:"statistics"`
}

type Publisher struct {
	PublisherId   string `json:"publisherId"`
	PublisherName string `json:"publisherName"`
	DisplayName   string `json:"displayName"`
}

type Version struct {
	Version    string     `json:"version"`
	Files      []File     `json:"files"`
	Properties []Property `json:"properties"`
}

type File struct {
	AssetType string `json:"assetType"`
	Source    string `json:"source"`
}

type Property struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Statistic struct {
	StatisticName string  `json:"statisticName"`
	Value         float64 `json:"value"`
}

type Request struct {
	Filters []Filter `json:"filters"`
	Flags   int      `json:"flags"`
}

type Filter struct {
	Criteria []Criteria `json:"criteria"`
}

type Criteria struct {
	FilterType int    `json:"filterType"`
	Value      string `json:"value"`
}
