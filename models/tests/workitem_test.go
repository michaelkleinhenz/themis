package models_test

import (
	. "themis/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Workitem", func() {
		var (
        workItem  WorkItem
    )

    BeforeEach(func() {
			workItem = *NewWorkItem()
    })

    Describe("Initializing new WorkItems", func() {
        Context("With correct dates", func() {
            It("should have a set create date", func() {
                Expect(workItem.CreatedAt).Should(Not(BeNil()))
            })
            It("should have a set update date", func() {
                Expect(workItem.UpdatedAt).Should(Not(BeNil()))
            })
        })
    })
})
