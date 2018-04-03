package component_data

import (
  "fmt"
)

type QuoteStruct struct {
  ComponentData struct {
    QuoteText string
    Color string
    FontSize int
    QuotationMarks bool
    ShowCard bool
    CardColor string
  }
}


func (this QuoteStruct) GetHtml() string {
  return fmt.Sprintf(`
    <lumavate-quote
        quote-text='%v'
        font-size='%vpt'
        show-card=%v
        card-color='%v'
        quotation-marks=%v
        color='%v'>
      </lumavate-quote>`,
    this.ComponentData.QuoteText,
    this.ComponentData.FontSize,
    this.ComponentData.ShowCard,
    this.ComponentData.CardColor,
    this.ComponentData.QuotationMarks,
    this.ComponentData.Color)
}

