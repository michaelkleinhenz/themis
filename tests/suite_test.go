package tests

import (
	"testing"
	"os"
	"io"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/mgo.v2"
	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go-adapter/gingonic"
	"gopkg.in/gin-gonic/gin.v1"

	"themis/schema"
	"themis/resources"
	"themis/models"
	"themis/database"
	"themis/mockdb"
	"themis/routes"
	"themis/utils"
)

type M map[string]interface{}
var dbServer dbserver.DBServer
var session *mgo.Session
var configuration utils.Configuration
var SpaceID string

func TestModels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suite")
}

var _ = BeforeSuite(func() {
	// setup logger
	utils.InitLogger()
	utils.SetLogFile("test.log")
	// test configuration
		configuration = utils.Configuration {
		ServiceURL: "http://localhost:8080",
    ServicePort: ":8080",
		ServiceMode: "production",
    DatabaseHost: "localhost",
		DatabasePort: 27017,
		DatabaseDatabase: "themis",
		DatabaseUser: "themis",
		DatabasePassword: "themis",
	}
	// launch test database instance
	os.Setenv("CHECK_SESSIONS", "1")
	dir, _ := ioutil.TempDir("", "themis_test")
	dbServer.SetPath(dir)
	session = dbServer.Session()
	db := session.DB(configuration.DatabaseDatabase)
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
	// Create the meta storage backend
	utils.NewDatabaseMeta(db)
	// setup test fixtures
	SpaceID = schema.SetupFixtureData(storageBackends)
	// launch service
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(utils.LogFile)
	r := gin.Default()
	r.Use(gin.Logger())
	api := api2go.NewAPIWithRouting(
		"api",
		api2go.NewStaticResolver(configuration.ServiceURL),
		gingonic.New(r),
	)
	r.StaticFile("/", "./static/index.html")
	api.AddResource(models.Space{}, resources.SpaceResource{SpaceStorage: storageBackends.Space})
	api.AddResource(models.WorkItem{}, resources.WorkItemResource{WorkItemStorage: storageBackends.WorkItem})
	api.AddResource(models.Area{}, resources.AreaResource{AreaStorage: storageBackends.Area})
	api.AddResource(models.Comment{}, resources.CommentResource{CommentStorage: storageBackends.Comment})
	api.AddResource(models.Iteration{}, resources.IterationResource{IterationStorage: storageBackends.Iteration})
	api.AddResource(models.Link{}, resources.LinkResource{LinkStorage: storageBackends.Link})
	api.AddResource(models.LinkCategory{}, resources.LinkCategoryResource{LinkCategoryStorage: storageBackends.LinkCategory})
	api.AddResource(models.LinkType{}, resources.LinkTypeResource{LinkTypeStorage: storageBackends.LinkType})
	api.AddResource(models.User{}, resources.UserResource{UserStorage: storageBackends.User, SpaceStorage: storageBackends.Space})
	api.AddResource(models.WorkItemType{}, resources.WorkItemTypeResource{WorkItemTypeStorage: storageBackends.WorkItemType})
	routes.Init(r)
	go r.Run(configuration.ServicePort)
})

var _ = AfterSuite(func() {	
	session.Close()
	dbServer.Wipe()
	dbServer.Stop()
	utils.CloseLogfile()
})
