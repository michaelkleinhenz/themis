package resources_test

import (
	"testing"

	"themis/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestResources(t *testing.T) {
	// setup logger
	utils.InitLogger()
	utils.SetLogFile("test.log")

	RegisterFailHandler(Fail)
	RunSpecs(t, "Resources Suite")
}
