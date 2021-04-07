package models

type ThemeDataStruct struct {
  MainFontFamily string `json:"mainFontFamily"`
  PrimaryColor string `json:"primaryColor"`
  SecondaryColor string `json:"secondaryColor"`
  TertiaryColor string `json:"tertiaryColor"`
  AccentColor string `json:"accentColor"`
  SuccessColor string `json:"successColor"`
  AlertColor string `json:"alertColor"`
  DarkColor string `json:"darkColor"`
  MediumColor string `json:"mediumColor"`
  LightColor string `json:"lightColor"`
}
