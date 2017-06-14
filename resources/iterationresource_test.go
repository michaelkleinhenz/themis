package resources

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/manyminds/api2go"
)

var _ = Describe("Iteration Resource", func() {

	var iterationResource IterationResource
	var request api2go.Request

	BeforeEach(func() {
		var iterationStorageMock IterationStorageMock
		var workItemStorageMock WorkItemStorageMock
		iterationResource = IterationResource{IterationStorage: iterationStorageMock, WorkItemStorage: workItemStorageMock}
		request = NewRequest(make(map[string][]string, 1), make(map[string]string, 1))
	})

	Describe("Querying Iteration Resource", func() {
		Context("With no parameters", func() {
			It("should return a valid Iteration", func() {
				resultIteration, err := iterationResource.FindOne("594009713bf1116a59b12b4b", request)
				Expect(err).Should(BeNil())
				Expect(resultIteration).ShouldNot(BeNil())
				// Expect(resultID).Should(Equal(SpaceID))
				// Expect(resultSpace.ID.Hex()).Should(Equal(SpaceID))
			})
		})
	})
})


