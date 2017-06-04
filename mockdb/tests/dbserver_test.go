package dbserver_test

import (
	"os"
	"time"
	"io/ioutil"

	"gopkg.in/mgo.v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"themis/mockdb"
)

var _ = Describe("Dbserver", func() {

	type M map[string]interface{}
	var oldCheckSessions string
	var server dbserver.DBServer

	BeforeEach(func() {
		oldCheckSessions = os.Getenv("CHECK_SESSIONS")
		os.Setenv("CHECK_SESSIONS", "1")
		dir, _ := ioutil.TempDir("", "themis_test")
		server.SetPath(dir)
		
		defer server.Stop()
	})

	AfterEach(func() {
		os.Setenv("CHECK_SESSIONS", oldCheckSessions)
	})

	Describe("Mongo database server", func() {
		Context("As a mock service", func() {
			It("should be able to wipe the database", func() {
				session := server.Session()
				err := session.DB("mydb").C("mycoll").Insert(M{"a": 1})
				session.Close()
				Expect(err).To(BeNil())

				server.Wipe()

				session = server.Session()
				names, err := session.DatabaseNames()
				session.Close()
				Expect(err).To(BeNil())
				for _, name := range names {
					if name != "local" && name != "admin" {
						Fail("Wipe should have removed this database: " + name)
					}
				}
			})

			It("should be able to stop", func() {
				// Server should not be running.
				process := server.ProcessTest()
				Expect(process).To(BeNil())

				session := server.Session()
				addr := session.LiveServers()[0]
				session.Close()

				// Server should be running now.
				process = server.ProcessTest()
				p, err := os.FindProcess(process.Pid)
				Expect(err).To(BeNil())
				p.Release()

				server.Stop()

				// Server should not be running anymore.
				session, err = mgo.DialWithTimeout(addr, 500*time.Millisecond)
				if session != nil {
					session.Close()
					Fail("Stop did not stop the server")
				}
			})

			It("be able to wipe with no check sessions", func() {
				os.Setenv("CHECK_SESSIONS", "0")

				// Should not panic, although it looks to Wipe like this session will leak.
				session := server.Session()
				defer session.Close()
				server.Wipe()
			})
		})
	})
})
