package dbserver_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMockdb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mockdb Suite")
}
