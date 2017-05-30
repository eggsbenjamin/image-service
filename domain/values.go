package domain

type ImageResult struct {
	Title    string `json:"title,omitempty"`
	Link     string `json:"link,omitempty"`
	MimeType string `json:"mime,omitempty"`
	Image    *struct {
		Height          float32 `json:"height,omitempty"`
		Width           float32 `json:"width,omitempty"`
		ThumbnailLink   string  `json:"thumbnailLink,omitempty"`
		ThumbnailHeight float32 `json:"thumbnailHeight,omitempty"`
		ThumbnailWidth  float32 `json:"thumbnailWidth,omitempty"`
	} `json:"image"`
}

type GoogleImageFetcher interface {
	GetImages(srchTrm string, imgTyp string, n int) ([]*ImageResult, error)
}
