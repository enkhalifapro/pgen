package utilities_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/enkhalifapro/pgen/utilities"
)

var _ = Describe("SimilarTechUtil", func() {
	Describe("GetTechnologies", func() {
		It("should return an instance of SimilarTechResponse with a good response and error should be nil", func() {
			similarTechUtil := &utilities.SimilarTechUtil{}
			resp, err := similarTechUtil.GetTechnologies("jet.com")
			Expect(resp).NotTo(BeNil())
			Expect(err).To(BeNil())
			Expect(resp.Found).To(Equal(true))
			Expect(resp.Technologies[0].ID).To(Equal(1779))
			Expect(len(resp.Technologies)).To(Equal(1))
		})
	})
})
