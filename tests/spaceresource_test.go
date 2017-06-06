package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"themis/client"

)

var _ = Describe("Space Service", func() {
    BeforeEach(func() {
	    //space = *NewSpace()
    })

    Describe("Querying Space Service", func() {
        Context("With no parameters", func() {
            It("should return a valid space", func() {
							resultSpace, rawData, err := client.GetSpace(configuration.ServiceURL, SpaceID)
							Expect(err).Should(BeNil())
							resultID := ((rawData["data"].(map[string]interface{})["id"])).(string)
							Expect(resultID).Should(Equal(SpaceID))
							Expect(resultSpace.ID.Hex()).Should(Equal(SpaceID))
            })
        })
    })
})
