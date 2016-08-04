package url_test

import (
	neturl "net/url"

	"github.com/challiwill/burly/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Url", func() {
	Describe("Parse", func() {
		Context("when given a valid interface", func() {
			var (
				urlStruct   testURLStruct
				expectedURL *neturl.URL
				err         error
			)

			BeforeEach(func() {
				urlStruct = testURLStruct{
					thing0: "https",
					thing1: "mydomain.com",
					thing2: "my/special/path",
					thing3: "one/value",
					thing4: "two-value",
				}

				expectedURL, err = neturl.Parse("https://mydomain.com/my/special/path?firstparam=one%2Fvalue&secondparam=two-value")
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns a properly constructed net/url.URL struct", func() {
				actualURL, err := url.Parse(urlStruct)
				Expect(actualURL.Scheme).To(Equal(expectedURL.Scheme))
				Expect(err).NotTo(HaveOccurred())
				Expect(actualURL.Host).To(Equal(expectedURL.Host))
				Expect(actualURL.Path).To(Equal(expectedURL.Path))
				Expect(actualURL.RawQuery).To(Equal(expectedURL.RawQuery))
			})
		})

	})
})

// thing1/thing2?firstparam=thing3&secondparam=thing4
type testURLStruct struct {
	thing0 string `url:"protocol"`
	thing1 string `url:"domain"`
	thing2 string `url:"path"`
	thing3 string `url:"firstparam"`
	thing4 string `url:"secondparam"`
}
