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

  H1FontStyle *FontStyleStruct `json:"h1FontStyle"`
  H2FontStyle *FontStyleStruct `json:"h2FontStyle"`
  H3FontStyle *FontStyleStruct `json:"h3FontStyle"`
  H4FontStyle *FontStyleStruct `json:"h4FontStyle"`
  ParagraphFontStyle *FontStyleStruct `json:"paragraphFontStyle"`
  LinkFontStyle *FontStyleStruct `json:"linkFontStyle"`
  ButtonFontStyle *FontStyleStruct `json:"buttonFontStyle"`

}

type FontStyleStruct struct {
  FontColor string `json:"fontColor"`
  FontFamily string `json:"fontFamily"`
  FontSize int `json:"fontSize"`
  FontUnderline bool `json:"fontUnderline"`
}
type FontStyleDisplayStruct struct {
  Name string `json:"name"`
  FontColor string `json:"fontColor"`
  FontFamily string `json:"fontFamily"`
  FontSize string `json:"fontSize"`
  FontUnderline string `json:"fontUnderline"`
}
