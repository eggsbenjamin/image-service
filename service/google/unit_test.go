// +build unit

package google_test

import (
	"errors"

	"github.com/eggsbenjamin/image-service/domain"
	. "github.com/eggsbenjamin/image-service/service/google"
	"github.com/eggsbenjamin/image-service/service/mocks"
	"github.com/stretchr/testify/mock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Google Service", func() {
	It("returns image results correctly", func() {
		By("setup")
		cl := &mocks.HTTPClient{}
		mockBody := `{"kind":"customsearch#search","url":{"type":"application/json","template":"https://www.googleapis.com/customsearch/v1?q={searchTerms}&num={count?}&start={startIndex?}&lr={language?}&safe={safe?}&cx={cx?}&cref={cref?}&sort={sort?}&filter={filter?}&gl={gl?}&cr={cr?}&googlehost={googleHost?}&c2coff={disableCnTwTranslation?}&hq={hq?}&hl={hl?}&siteSearch={siteSearch?}&siteSearchFilter={siteSearchFilter?}&exactTerms={exactTerms?}&excludeTerms={excludeTerms?}&linkSite={linkSite?}&orTerms={orTerms?}&relatedSite={relatedSite?}&dateRestrict={dateRestrict?}&lowRange={lowRange?}&highRange={highRange?}&searchType={searchType}&fileType={fileType?}&rights={rights?}&imgSize={imgSize?}&imgType={imgType?}&imgColorType={imgColorType?}&imgDominantColor={imgDominantColor?}&alt=json"},"queries":{"request":[{"title":"GoogleCustomSearch-RightGuardTotalDefence5FreshAntiperspirantDeodorant250Ml","totalResults":"38200","searchTerms":"RightGuardTotalDefence5FreshAntiperspirantDeodorant250Ml","count":1,"startIndex":1,"inputEncoding":"utf8","outputEncoding":"utf8","safe":"off","cx":"004306976011237526387:u3abktktudg","fileType":"jpg","searchType":"image"}],"nextPage":[{"title":"GoogleCustomSearch-RightGuardTotalDefence5FreshAntiperspirantDeodorant250Ml","totalResults":"38200","searchTerms":"RightGuardTotalDefence5FreshAntiperspirantDeodorant250Ml","count":1,"startIndex":2,"inputEncoding":"utf8","outputEncoding":"utf8","safe":"off","cx":"004306976011237526387:u3abktktudg","fileType":"jpg","searchType":"image"}]},"context":{"title":"imageapi"},"searchInformation":{"searchTime":0.663327,"formattedSearchTime":"0.66","totalResults":"38200","formattedTotalResults":"38,200"},"items":[{"kind":"customsearch#result","title":"Right Guard Total Defence 5 Anti-Perspirant 250ml - Be Beautiful","htmlTitle":"\u003cb\u003eRightGuardTotalDefence5Anti-Perspirant250ml\u003c/b\u003e-BeBeautiful","link":"http://www.lifeandlooks.com/user/products/large/power%20and%20car%20250ml.jpg","displayLink":"www.lifeandlooks.com","snippet":"RightGuardTotalDefence5Anti-Perspirant250ml-BeBeautiful","htmlSnippet":"\u003cb\u003eRightGuardTotalDefence5Anti-Perspirant250ml\u003c/b\u003e-BeBeautiful","mime":"image/jpeg","image":{"contextLink":"http://www.lifeandlooks.com/right-guard-total-defence-5-anti-perspirant-250ml.html","height":640,"width":640,"byteSize":22953,"thumbnailLink":"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQXubhB8iS30CuqZzIfMfJb2B_Wp0zlxyJ7F6YwhStBUIkoV8L-AxyCp2Q_","thumbnailHeight":137,"thumbnailWidth":137}}]}`
		mockResponse := CreateMockResponse(200, mockBody)
		expected := []*domain.ImageResult{
			&domain.ImageResult{
				Title:    "Right Guard Total Defence 5 Anti-Perspirant 250ml - Be Beautiful",
				Link:     "http://www.lifeandlooks.com/user/products/large/power%20and%20car%20250ml.jpg",
				MimeType: "image/jpeg",
				Image: &struct {
					Height          float32 `json:"height,omitempty"`
					Width           float32 `json:"width,omitempty"`
					ThumbnailLink   string  `json:"thumbnailLink,omitempty"`
					ThumbnailHeight float32 `json:"thumbnailHeight,omitempty"`
					ThumbnailWidth  float32 `json:"thumbnailWidth,omitempty"`
				}{
					Height:          float32(640),
					Width:           float32(640),
					ThumbnailLink:   "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQXubhB8iS30CuqZzIfMfJb2B_Wp0zlxyJ7F6YwhStBUIkoV8L-AxyCp2Q_",
					ThumbnailWidth:  float32(137),
					ThumbnailHeight: float32(137),
				},
			},
		}
		cl.On("Do", mock.Anything).Return(mockResponse, nil)
		gSrv := NewGoogleImageSearchService(cl)

		By("making call")
		actual, err := gSrv.GetImages("", "", 1)

		By("assertion")
		Expect(err).NotTo(HaveOccurred())
		Expect(actual).To(Equal(expected))
	})

	Context("when a non-200 status code is returned", func() {
		It("returns the correct error", func() {
			By("setup")
			cl := &mocks.HTTPClient{}
			mockResponse := CreateMockResponse(500, `{}`)
			expectedError := errors.New("unable to get images")
			cl.On("Do", mock.Anything).Return(mockResponse, nil)
			gSrv := NewGoogleImageSearchService(cl)

			By("making call")
			actual, err := gSrv.GetImages("", "", 1)

			By("assertion")
			Expect(actual).To(BeNil())
			Expect(err).To(Equal(expectedError))
		})
	})

	Context("when a client error occurs", func() {
		It("returns the correct error", func() {
			By("setup")
			cl := &mocks.HTTPClient{}
			clientError := errors.New("client error")
			expectedError := errors.New("unable to get images")
			cl.On("Do", mock.Anything).Return(nil, clientError)
			gSrv := NewGoogleImageSearchService(cl)

			By("making call")
			actual, err := gSrv.GetImages("", "", 1)

			By("assertion")
			Expect(actual).To(BeNil())
			Expect(err).To(Equal(expectedError))
		})
	})
})
