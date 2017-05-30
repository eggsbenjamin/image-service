package google

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/eggsbenjamin/image-service/domain"
	"github.com/spf13/viper"
)

type GoogleImageService struct {
	cl HTTPClient
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type noopCloser struct {
	io.Reader
}

func (n *noopCloser) Close() error {
	return nil
}

func CreateMockResponse(sc int, bStr string) *http.Response {
	return &http.Response{
		StatusCode: sc,
		Body:       &noopCloser{bytes.NewBufferString(bStr)},
	}
}

type ImageResponse struct {
	Items []*domain.ImageResult `json:"items"`
}

func (g *GoogleImageService) GetImages(srchTrm string, imgTyp string, n int) ([]*domain.ImageResult, error) {
	url := fmt.Sprintf(
		"%s?key=%s&cx=%s&q=%s&searchType=image&fileType=%s&num=%d",
		viper.GetString("search_api_base_url"),
		viper.GetString("search_api_key"),
		viper.GetString("search_api_cx"),
		srchTrm,
		imgTyp,
		n,
	)
	rErr := errors.New("unable to get images")
	rq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to create request - %v", err)
		return nil, rErr
	}
	rsp, err := g.cl.Do(rq)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to make request - %v", err)
		return nil, rErr
	}
	raw, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if rsp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "error from remote server - %s", string(raw))
		return nil, rErr
	}
	imgRsp := &ImageResponse{}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error from reading response body - %v", err)
		return nil, rErr
	}
	if err = json.Unmarshal(raw, imgRsp); err != nil {
		fmt.Fprintf(os.Stderr, "error marshaling json - %v", err)
		return nil, rErr
	}
	return imgRsp.Items, nil
}

// constructor
func NewGoogleImageSearchService(cl HTTPClient) *GoogleImageService {
	return &GoogleImageService{
		cl: cl,
	}
}
