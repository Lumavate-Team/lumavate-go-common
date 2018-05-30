package models

import (
  "github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
)

type CommonWidgetStruct struct {
  InstanceName string `json:"instance__name"`
  PageType component_data.PageTypeStruct `json:"pageType"`
  component_data.NavBarContainerStruct
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
  Token string `json:"token"`
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
