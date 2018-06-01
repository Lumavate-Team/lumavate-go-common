package models

import (
  "github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
)

type CommonWidgetStruct struct {
  InstanceName string `json:"instance__name"`
  PageType component_data.PageTypeStruct `json:"pageType"`
  component_data.NavBarContainerStruct
  component_data.AddToHomeStruct
}


type ValidatePhoneResponse struct {
  Payload struct {
    Data struct {
      CountryCode string `json:"countryCode"`
      FormattedPhone string `json:"formattedPhone"`
      PhoneNumber string `json:"phoneNumber"`
    } `json:"data"`
  } `json:"payload"`
}

type SingleUseToken struct {
  Payload struct {
    Data struct {
      Token string `json:"token"`
    } `json:"data"`
  } `json:"payload"`
}

type PersonCategories struct {
  Payload struct {
    Data [] struct {
      Id int `json:"id"`
      Name string `json:"name"`
    } `json:"data"`
  } `json:"payload"`
}

type PersonResponse struct {
  Payload struct {
    EmailAddress string  `json:"emailAddress"`
    Id float64 `json:"id"`
    Data map[string]interface{} `json:"data"`
  } `json:"payload"`
}

type ErrorPayload struct {
    Error string `json:"error"`
    ApiField string `json:"apiField"`
}
type  ErrorResponse struct {
  Payload ErrorPayload `json:"payload"`
  CallType string `json:"callType"`
}
