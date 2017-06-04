package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestThemis(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Themis Suite")
}
