package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"themis/client"

)

var _ = Describe("Space Service", func() {
		//var space models.Space

    BeforeEach(func() {
	    //space = *NewSpace()
    })

    Describe("Querying Space Service", func() {
        Context("With no parameters", func() {
            It("should return a valid space", func() {
							_, rawData, err := client.GetSpace(configuration.ServiceURL, SpaceID)
							//space = *thisSpace
							resultID := ((rawData["data"].(map[string]interface{})["attributes"]).(map[string]interface{})["ID"]).(string)
							Expect(resultID).Should(Equal(SpaceID))
							Expect(err).Should(BeNil())
            })
        })
    })
})
