package handlers

import (
	"github.com/DuriDuri/parsley/db"
	log "github.com/Sirupsen/logrus"
)

type Handler struct {
	Datastore *db.DB
	logger    *log.Logger
}

func Init() (*Handler, error) {
	db, err := db.Init()
	if err != nil {
		return nil, err
	}

	return &Handler{
		Datastore: db,
		logger:    log.New(),
	}, nil
}
