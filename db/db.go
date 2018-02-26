package db

import (
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
	host := os.Getenv("DB_HOST")
	if host == "" {
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
	return session, nil
}

func (db *DB) Collection() *mgo.Collection {
	return db.Session.DB(dbName).C(patientCollection)
}
