package handlers

import (
	"fmt"
	"github.com/DuriDuri/parsley/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) CreatePatientRecord(gctx *gin.Context) {
	req := new(models.Patient)
	err := gctx.BindJSON(req)
	if err != nil {
		h.logger.Errorf("Error occurred when unmarshalling in CreateRecord: %s", err.Error())
		gctx.JSON(400, gin.H{"status": fmt.Sprintf("Invalid schema and/or data types: %s", err.Error())})
		return
	}

	err = req.Validate()
	if err != nil {
		gctx.JSON(400, gin.H{"status": fmt.Sprintf("Invalid data types: %s", err.Error())})
		return
	}

	err = h.Datastore.Collection().Insert(req)
	if err != nil {
		h.logger.Errorf("Error occurred when inserting record: %s", err.Error())
		gctx.JSON(422, gin.H{"status": fmt.Sprintf("Failure saving record: %s", err.Error())})
		return
	}

	gctx.JSON(201, gin.H{"status": "Saved."})
}

func (h *Handler) GetPatientRecord(gctx *gin.Context) {
	recordId := gctx.Param("id")
	patientRecord := make(map[string]interface{})
	err := h.Datastore.Collection().FindId(bson.ObjectIdHex(recordId)).One(&patientRecord)
	if err != nil {
		h.logger.Errorf("Error when searching record: %s", err.Error())
		gctx.JSON(422, gin.H{"status": fmt.Sprintf("Error finding record: %s", err.Error())})
		return
	}

	gctx.JSON(200, patientRecord)
}

func (h *Handler) UpdatePatientRecord(gctx *gin.Context) {
	req := new(models.Patient)
	err := gctx.BindJSON(req)
	if err != nil {
		h.logger.Errorf("Error occurred when unmarshalling in UpdateRecord: %s", err.Error())
		gctx.JSON(400, gin.H{"status": "Invalid schema."})
		return
	}

	err = req.Validate()
	if err != nil {
		gctx.JSON(400, gin.H{"status": fmt.Sprintf("Invalid data types: %s", err.Error())})
		return
	}

	recordId := gctx.Param("id")
	err = h.Datastore.Collection().UpdateId(bson.ObjectIdHex(recordId), req)
	if err != nil {
		h.logger.Errorf("Error occurred when updating record: %s", err.Error())
		gctx.JSON(422, gin.H{"status": fmt.Sprintf("Failure updating record: %s", err.Error())})
		return
	}

	gctx.JSON(200, gin.H{"status": "Updated."})
}

func (h *Handler) DeletePatientRecord(gctx *gin.Context) {
	recordId := gctx.Param("id")
	err := h.Datastore.Collection().RemoveId(bson.ObjectIdHex(recordId))

	if err != nil {
		h.logger.Errorf("Error occurred when deleting record: %s", err.Error())
		gctx.JSON(422, gin.H{"status": fmt.Sprintf("Failure removing record: %s", err.Error())})
		return
	}

	gctx.JSON(200, gin.H{"status": "Removed."})
}

func (h *Handler) ListPatientsRecords(gctx *gin.Context) {
	var patients []map[string]interface{}
	err := h.Datastore.Collection().Find(nil).All(&patients)
	if err != nil {
		h.logger.Errorf("Error occurred when listing records: %s", err.Error())
		gctx.JSON(422, gin.H{"status": fmt.Sprintf("Failure listing records: %s", err.Error())})
		return
	}

	if patients == nil {
		patients = []map[string]interface{}{}
	}

	gctx.JSON(200, patients)
}
