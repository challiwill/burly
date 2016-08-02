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
					thing3: "one-value",
					thing4: "two-value",
				}

				expectedURL, err = neturl.Parse("https://mydomain.com/my/special/path?firstparam=one-value&secondparam=two-value")
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns a properly constructed net/url.URL struct", func() {
				actualURL, err := url.Parse(urlStruct)
				Expect(err).NotTo(HaveOccurred())
				Expect(*actualURL).To(Equal(*expectedURL))
			})
		})

	})
})

// thing1/thing2?firstparam=thing3&secondparam=thing4
type testURLStruct struct {
	thing0 string `comp:"protocol"`
	thing1 string `comp:"domain"`
	thing2 string `comp:"path"`
	thing3 string `comp:"firstparam"`
	thing4 string `comp:"secondparam"`
}
