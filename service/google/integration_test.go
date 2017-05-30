// +build integration

package google_test

import (
	"net/http"

	. "github.com/eggsbenjamin/image-service/service/google"
	"github.com/spf13/viper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("google")
}

var _ = Describe("Google Service", func() {
	Context("when calling the Google Search API", func() {
		It("returns the results correctly", func() {
			By("setup")
			cl := &http.Client{}
			gSrv := NewGoogleImageSearchService(cl)

			By("making call")
			actual, err := gSrv.GetImages("test", "jpg", 2)

			By("assertion")
			Expect(err).NotTo(HaveOccurred())
			Expect(actual).To(HaveLen(2))
		})
	})
})
