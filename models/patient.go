package models

import (
	"gopkg.in/go-playground/validator.v9"
	"time"
)

type Patient struct {
	FirstName  string `json:"firstName" validate:"alpha"`
	MiddleName string `json:"middleName" validate:"alpha"`
	LastName   string `json:"lastName validate:"alpha"`
	Phones     []struct {
		Type   string `json:"type" validate:"alpha"`
		Number string `json:"number" validate:"numeric"`
	} `json:"phones"`
	Email           string    `json:"email" validate:"email"`
	Dob             string    `json:"dob"`
	Age             int       `json:"age" validate:"gte=0,lte=100"`
	Gender          string    `json:"gender" validate:"alpha"`
	Status          string    `json:"status" validate:"alpha"`
	TermsAccepted   bool      `json:"termsAccepted"`
	TermsAcceptedAt time.Time `json:"termsAcceptedAt"`
	Address         struct {
		Line1 string `json:"line1"`
		Line2 string `json:"line2"`
		City  string `json:"city" validate:"alpha"`
		State string `json:"state" validate:"len=2,alpha"`
		Zip   string `json:"zip" validate:"numeric,len=5"`
	} `json:"address"`
}

func (p *Patient) Validate() error {
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		return err
	}

	return nil
}
