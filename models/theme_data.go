package models

type ThemeDataStruct struct {
  MainFontFamily string `json:"mainFontFamily"`
  SecondaryFontFamily string `json:"secondaryFontFamily"`
  TertiaryFontFamily string `json:"tertiaryFontFamily"`
  PrimaryColor string `json:"primaryColor"`
  SecondaryColor string `json:"secondaryColor"`
  TertiaryColor string `json:"tertiaryColor"`
  AccentColor string `json:"accentColor"`
  SuccessColor string `json:"successColor"`
  AlertColor string `json:"alertColor"`
  DarkColor string `json:"darkColor"`
  MediumColor string `json:"mediumColor"`
  LightColor string `json:"lightColor"`

  Heading1FontStyle FontStyleStruct `json:"h1FontStyle"`
  Heading2FontStyle FontStyleStruct `json:"h2FontStyle"`
  Heading3FontStyle FontStyleStruct `json:"h3FontStyle"`
  Heading4FontStyle FontStyleStruct `json:"h4FontStyle"`
  ParagraphFontStyle FontStyleStruct `json:"paragraphFontStyle"`
  LinkFontStyle FontStyleStruct `json:"linkFontStyle"`
  ButtonFontStyle FontStyleStruct `json:"buttonFontStyle"`

}

type FontStyleStruct struct {
  FontColor string `json:"fontColor"`
  FontFamily string `json:"fontFamily"`

}
