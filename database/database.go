package database

import (
	"fmt"

	"gopkg.in/mgo.v2"

	"themis/utils"
)

// Connect connects to the database, returning a database handle.
func Connect(configuration utils.Configuration) (*mgo.Session, *mgo.Database) {
	session, err := mgo.DialWithInfo(&mgo.DialInfo {
		Addrs:    []string { configuration.DatabaseHost },
		Username: configuration.DatabaseUser,
		Password: configuration.DatabasePassword,
		Database: configuration.DatabaseDatabase,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected to database at %v!\n", session.LiveServers())
	return session, session.DB(configuration.DatabaseDatabase)
}

// Close closes the database connection.
func Close(database *mgo.Database) {
	database.Session.Close()
}
