package tests

import (
	"testing"
	"os"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/mgo.v2"

	"themis/schema"
	"themis/database"
	"themis/mockdb"
	"themis/utils"
)

type M map[string]interface{}
var dbServer dbserver.DBServer
var session *mgo.Session

func TestModels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suite")
}

var _ = BeforeSuite(func() {
	// setup logger
	utils.InitLogger()
	// launch test database instance
	os.Setenv("CHECK_SESSIONS", "1")
	dir, _ := ioutil.TempDir("", "themis_test")
	dbServer.SetPath(dir)
	session = dbServer.Session()
	db := session.DB("themis_test")
	// creating all storage backends
	storageBackends := database.StorageBackends {
		Space: database.NewSpaceStorage(db),
		WorkItem: database.NewWorkItemStorage(db),
		WorkItemType: database.NewWorkItemTypeStorage(db),
		Area: database.NewAreaStorage(db),
		Comment: database.NewCommentStorage(db),
		Iteration: database.NewIterationStorage(db),
		LinkCategory: database.NewLinkCategoryStorage(db),
		Link: database.NewLinkStorage(db),
		LinkType: database.NewLinkTypeStorage(db),
		User: database.NewUserStorage(db),
	}
	// setup test fixtures
	schema.SetupFixtureData(storageBackends)
})

var _ = AfterSuite(func() {
	session.Close()
	dbServer.Wipe()
	dbServer.Stop()
})
