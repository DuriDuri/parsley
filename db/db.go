package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"os"
)

const (
	dbName            = "parsley"
	patientCollection = "patients"
)

type DB struct {
	Session *mgo.Session
}

func Init() (*DB, error) {
	// Database host from the environment variables
	host := os.Getenv("DB_HOST")
	if host == "" {
		// mongo instance for deployed version
		fmt.Print("Didnt find DB_HOST")
		host = "mongodb://parsley:qwerty123@ds147228.mlab.com:47228/parsley"
	}

	s, err := CreateSession(host)
	if err != nil {
		return nil, err
	}

	return &DB{
		Session: s,
	}, nil
}

func CreateSession(host string) (*mgo.Session, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)
	return session.Clone(), nil
}

func (db *DB) Collection() *mgo.Collection {
	return db.Session.DB(dbName).C(patientCollection)
}
