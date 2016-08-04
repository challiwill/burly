package url_test

import (
	"github.com/challiwill/burly/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Url", func() {
	Describe("Parse", func() {
		Context("when given a valid interface", func() {
			var urlStruct encodedURLStruct

			BeforeEach(func() {
				urlStruct = encodedURLStruct{
					thing0: "https",
					thing1: "mydomain.com",
					thing2: "my/special/path",
					thing3: "one/value",
					thing4: "two-value",
				}
			})

			It("returns a properly constructed net/url.URL struct", func() {
				actualURL, err := url.Parse(urlStruct)
				Expect(err).NotTo(HaveOccurred())
				Expect(actualURL.Scheme).To(Equal("https"))
				Expect(actualURL.Host).To(Equal("mydomain.com"))
				Expect(actualURL.Path).To(Equal("/my/special/path"))
				Expect(actualURL.RawQuery).To(Equal("firstparam=one%2Fvalue&secondparam=two-value"))
			})

			Context("when given an interface that doesn't want to be encoded", func() {
				var urlStruct unencodedURLStruct

				BeforeEach(func() {
					urlStruct = unencodedURLStruct{
						thing0: "https",
						thing1: "mydomain.com",
						thing2: "my/special/path",
						thing3: "one/value",
						thing4: "two-value",
					}
				})

				It("returns a properly constructed net/url.URL struct", func() {
					actualURL, err := url.Parse(urlStruct)
					Expect(err).NotTo(HaveOccurred())
					Expect(actualURL.Scheme).To(Equal("https"))
					Expect(actualURL.Host).To(Equal("mydomain.com"))
					Expect(actualURL.Path).To(Equal("/my/special/path"))
					Expect(actualURL.RawQuery).To(Equal("firstparam=one/value&secondparam=two-value"))
				})
			})
		})

		Context("when not pased a struct", func() {
			It("returns an appropriate error", func() {
				_, err := url.Parse("this is not a struct")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("must be a struct"))
			})
		})
	})
})

type encodedURLStruct struct {
	thing0 string `url:"protocol"`
	thing1 string `url:"domain"`
	thing2 string `url:"path"`
	thing3 string `url:"firstparam"`
	thing4 string `url:"secondparam"`
}

type unencodedURLStruct struct {
	thing0 string `url:"protocol"`
	thing1 string `url:"domain"`
	thing2 string `url:"path"`
	thing3 string `url:"firstparam" encode:"false"`
	thing4 string `url:"secondparam"`
}
