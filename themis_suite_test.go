package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

//var dbRunner *db.Runner
//var dbClient *db.Client

func TestThemis(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Themis Suite")
}

/*
var _ = BeforeSuite(func() {
    dbRunner = db.NewRunner()
    err := dbRunner.Start()
    Expect(err).NotTo(HaveOccurred())

    dbClient = db.NewClient()
    err = dbClient.Connect(dbRunner.Address())
    Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
    dbClient.Cleanup()
    dbRunner.Stop()
})
*/