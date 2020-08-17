package models

import (
    "encoding/json"
)

type WidgetData struct {
  Payload struct {
    Data struct {
      AuthData struct {
        Roles [] string `json:"roles"`
        Status string `json:"status"`
        User string `json:"user"`
      }
      ActivationData ActivationStruct
      Session map[string]interface{}
      PlatformVersion string 
      WidgetData json.RawMessage
      BrandingData json.RawMessage
      DomainData struct {
        Domain string `json:"domain"`
        IsTest bool `json:"isTest"`
        RuntimeData map[string]interface{} `json:"runtimeData"`
      }
      TokenData TokenDataStruct
      Resources struct {
        Pages [] struct {
          Id string
          Url string
        }
        Microservices [] struct {
          Id string
          Url string
        }
      }
    }
  }
}


