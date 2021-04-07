package components

import (
  "github.com/Lumavate-Team/lumavate-go-common/properties"
"fmt"
  "os"
)

func GetFormItemsProperty() *properties.PropertyComponents {
  return &properties.PropertyComponents {
    &properties.PropertyBase{"formItems", "Form", "Field Types", "Field Types", "", ""},
    [] *properties.Component{}, &properties.PropertyOptionsComponents{[] string {"formText", "formDate", "formCheckbox", "formAddress", "formEmailAddress", "formHidden"}, [] *properties.Component {GetTextFormComponent(), GetDateFormComponent(), GetCheckboxFormComponent(), GetAddressFormComponent(), GetEmailFormComponent(), GetHiddenFormComponent()}},
  }
}

func GetTextFormComponent() *properties.Component {
  props := [] properties.PropertyType {}
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"textLabel", "", "", "Label", "", ""}, "", properties.PropertyOptionsText{}})
  props = append(props, &properties.PropertyText{
  		&properties.PropertyBase{"textFieldName", "", "", "Field Name", "", ""}, "", properties.PropertyOptionsText{}})
	props = append(props, &properties.PropertyText{
  		&properties.PropertyBase{"textRegExpression", "", "", "Regular Expression", "", ""}, "", properties.PropertyOptionsText{}})
	props = append(props, &properties.PropertyToggle{
	    &properties.PropertyBase{"textRequired", "", "", "Require Field", "", ""},
	    true})
	props = append(props, &properties.PropertyNumeric{
      &properties.PropertyBase{"textMaxLength", "", "", "Max Length", "", ""}, 0, properties.PropertyOptionsNumeric{ Min: 0, Max: 32}})
	props = append(props, &properties.PropertyText{
  		&properties.PropertyBase{"textErrorText", "", "", "Error Text", "", ""}, "", properties.PropertyOptionsText{}})
	props = append(props, &properties.PropertyText{
  		&properties.PropertyBase{"textInputType", "", "", "Input Type", "", ""}, "", properties.PropertyOptionsText{}})
	image := fmt.Sprintf("%v%v/static/images/form-text.svg", os.Getenv("WIDGET_URL_PREFIX"),os.Getenv("PUBLIC_KEY"))
  return &properties.Component{"formText", "", "formText", "Text", image, "Text", props, ""}
}

func GetDateFormComponent() *properties.Component {
  props := [] properties.PropertyType {}
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"dateLabel", "", "", "Label", "", ""}, "", properties.PropertyOptionsText{}})
  props = append(props, &properties.PropertyText{
  		&properties.PropertyBase{"dateFieldName", "", "", "Field Name", "", ""}, "", properties.PropertyOptionsText{}})
	props = append(props, &properties.PropertyToggle{
	    &properties.PropertyBase{"dateRequired", "", "", "Require Field", "", ""},
	    true})
	image := fmt.Sprintf("%v%v/static/images/form-date.svg", os.Getenv("WIDGET_URL_PREFIX"),os.Getenv("PUBLIC_KEY"))
  return &properties.Component{"formDate", "", "formDate", "Date", image, "Date", props, ""}
}

func GetDropDownFormComponent() *properties.Component {
  props := [] properties.PropertyType {}
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"dropDownLabel", "", "", "Label", "", ""}, "", properties.PropertyOptionsText{}})
  props = append(props, &properties.PropertyText{
  		&properties.PropertyBase{"dropDownFieldName", "", "", "Field Name", "", ""}, "", properties.PropertyOptionsText{}})
	props = append(props, &properties.PropertyToggle{
	    &properties.PropertyBase{"dropDownRequired", "", "", "Require Field", "", ""},
	    true})
	props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"dropDownItems", "", "", "Drop Down Items", "", ""}, "", properties.PropertyOptionsText{}})
	image := fmt.Sprintf("%v%v/static/images/form-dropdown.svg", os.Getenv("WIDGET_URL_PREFIX"),os.Getenv("PUBLIC_KEY"))
  return &properties.Component{"formDropDown", "", "formDropDown", "Drop Down", image, "Drop Down", props, ""}
}

func GetCheckboxFormComponent() *properties.Component {
  props := [] properties.PropertyType {}
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"checkboxLabel", "", "", "Label", "", ""}, "", properties.PropertyOptionsText{}})
  props = append(props, &properties.PropertyText{
  		&properties.PropertyBase{"checkboxFieldName", "", "", "Field Name", "", ""}, "", properties.PropertyOptionsText{}})
	image := fmt.Sprintf("%v%v/static/images/form-checkbox.svg", os.Getenv("WIDGET_URL_PREFIX"),os.Getenv("PUBLIC_KEY"))
  return &properties.Component{"formCheckbox", "", "formCheckbox", "Checkbox", image, "Checkbox", props, ""}
}

func GetAddressFormComponent() *properties.Component {
  props := [] properties.PropertyType {}
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"addressLabel", "", "", "Label", "", ""}, "", properties.PropertyOptionsText{}})
  props = append(props, &properties.PropertyText{
  		&properties.PropertyBase{"addressFieldName", "", "", "Field Name", "", ""}, "", properties.PropertyOptionsText{}})
  props = append(props, &properties.PropertyToggle{
    &properties.PropertyBase{"addressRequired", "", "", "Require Field", "", ""},
	    true})
	image := fmt.Sprintf("%v%v/static/images/form-address.svg", os.Getenv("WIDGET_URL_PREFIX"),os.Getenv("PUBLIC_KEY"))
  return &properties.Component{"formAddress", "", "formAddress", "Address", image, "Address", props, ""}
}

func GetEmailFormComponent() *properties.Component {
  props := [] properties.PropertyType {}
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"emailLabel", "", "", "Label", "", ""}, "", properties.PropertyOptionsText{}})
  props = append(props, &properties.PropertyText{
  		&properties.PropertyBase{"emailFieldName", "", "", "Field Name", "", ""}, "", properties.PropertyOptionsText{}})
  props = append(props, &properties.PropertyToggle{
	    &properties.PropertyBase{"emailRequired", "", "", "Require Field", "", ""},
	    true})
	image := fmt.Sprintf("%v%v/static/images/form-email.svg", os.Getenv("WIDGET_URL_PREFIX"),os.Getenv("PUBLIC_KEY"))
  return &properties.Component{"formEmailAddress", "", "formEmailAddress", "Email Address", image, "Email Address", props, ""}
}

func GetHiddenFormComponent() *properties.Component {
  props := [] properties.PropertyType {}
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"hiddenLabel", "", "", "Label", "", ""}, "", properties.PropertyOptionsText{}})
  props = append(props, &properties.PropertyText{
  		&properties.PropertyBase{"hiddenFieldName", "", "", "Field Name", "", ""}, "", properties.PropertyOptionsText{}})
	image := fmt.Sprintf("%v%v/static/images/form-hidden.svg", os.Getenv("WIDGET_URL_PREFIX"),os.Getenv("PUBLIC_KEY"))
  return &properties.Component{"formHidden", "", "formHidden", "Hidden", image, "Hidden", props, ""}
}
