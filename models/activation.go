package models
import "time"

type ActivationStruct struct {
	Id int
	Key string
	CreatedAt time.Time
	CreatedBy int
	Extension map[string]interface{}
	LastModifiedAt time.Time
	LastModifiedBy int
  SerialNumber string
	SiteName string
}

