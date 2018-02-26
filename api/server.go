package api

import (
	"github.com/DuriDuri/parsley/handlers"
	"github.com/gin-gonic/gin"
)

type API struct {
	Port    string
	handler *handlers.Handler
}

func Init() (*API, error) {
	h, err := handlers.Init()
	if err != nil {
		return nil, err
	}

	return &API{
		Port:    "8080",
		handler: h,
	}, nil
}

func (a *API) GetServer() *gin.Engine {
	router := gin.New()

	router.POST("/patient", a.handler.CreatePatientRecord)
	router.GET("/patient/:id", a.handler.GetPatientRecord)
	router.PATCH("/patient/:id", a.handler.UpdatePatientRecord)
	router.DELETE("/patient/:id", a.handler.DeletePatientRecord)
	router.GET("/patients/list", a.handler.ListPatientsRecords)

	return router
}
