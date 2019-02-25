package models

import (
		"encoding/json"
)

type WidgetData struct {
	Payload struct {
		Data struct {
      AuthData struct {
        Roles [] string
				Status string
				User string
      }
			ActivationData ActivationStruct
			Session map[string]interface{}
			WidgetData json.RawMessage
			DomainData struct {
				Domain string
				RuntimeData map[string]interface{}
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


