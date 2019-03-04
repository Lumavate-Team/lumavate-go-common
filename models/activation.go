package models
import "time"

type ActivationStruct struct {
  Id int `json:"id"`
  Key string `json:"key"`
  CreatedAt time.Time `json:"createdAt"`
  CreatedBy int `json:"createdBy"`
  Extension map[string]interface{} `json:"extension"`
  LastModifiedAt time.Time `json:"lastModifiedAt"`
  LastModifiedBy int `json:"lastModifiedBy"`
  SerialNumber string `json:"serialNumber"`
  SiteName string `json:"siteName"`
}

