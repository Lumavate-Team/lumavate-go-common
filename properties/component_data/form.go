package component_data

import (
  "fmt"
  "reflect"
  "encoding/json"
  _"html/template"
)

type FormItemsStruct struct {
	ComponentData struct {
		Forms []ComponentData
	}
}

type FormStruct struct {
	FormItems FormItemsStruct
}

type FormContainerStruct struct {
	Form FormStruct
	FormItems FormItemsStruct
}

type FormTextStruct struct {
	TextLabel string `json:"textLabel"`
	TextFieldName string `json:"textFieldName"`
	TextRegExpression string `json:"textRegExpression"`
	TextRequired bool `json:"textRequired"`
	TextMaxLength int `json:"textMaxLength"`
	TextErrorText string `json:"textErrorText"`
	TextInputType string `json:"textInputType"`
}

type DateStruct struct {
	DateLabel string `json:"dateLabel"`
	DateFieldName string `json:"dateFieldName"`
	DateRequired bool `json:"dateRequired"`
}

type DropDownStruct struct {
	DropDownLabel string `json:"dropDownLabel"`
	DropDownFieldName string `json:"dropDownFieldName"`
	DropDownRequired bool `json:"dropDownRequired"`
	DropDownItems string //not really sure about this
}

type CheckboxStruct struct {
	CheckboxLabel string `json:"checkboxLabel"`
	CheckboxFieldName string `json:"checkboxFieldName"`
}

type AddressStruct struct {
	AddressLabel string `json:"addressLabel"`
	AddresFieldName string `json:"addressFieldName"`
	AddressRequired bool `json:"addressRequired"`
}

type EmailStruct struct {
	EmailLabel string `json:"emailLabel"`
	EmailFieldName string `json:"emailFieldName"`
	EmailRequired bool `json:"emailRequired"`
}

type HiddenStruct struct {
	HiddenLabel string `json:"hiddenLabel"`
	HiddenFieldName string `json:"hiddenFieldName"`
}

func (this FormTextStruct) GetHtml() string {
	return fmt.Sprintf(`
		<div>
			<input class="pure-input-1" aria-label="%v" id="%v" name="%v" type="%v" placeholder="%v" 
			{{if .TextRequired}}
				required
			{{end}}>
		</div>`,
		this.TextLabel,
		this.TextFieldName,
		this.TextFieldName,
		this.TextInputType,
		this.TextLabel)
}

func (this DateStruct) GetHtml() string {
	return fmt.Sprintf(`
		<div>
			<input aria-label="%v" id="%v" name="%v" type="date">
		</div>`,
		this.DateLabel,
		this.DateFieldName,
		this.DateFieldName)
}

func (this DropDownStruct) GetHtml() string {
	return fmt.Sprintf(`
		<div>
			<select name="%v">
				{{range item := .dropDownItems}}
					<option value="%v">%v</option>
				{{end}}
			</select>
		</div>`,
		this.DropDownFieldName,
		this.DropDownFieldName,
		this.DropDownLabel)
}

func (this CheckboxStruct) GetHtml() string {
	return fmt.Sprintf(`
		<div>
			<input type="checkbox" name="%v" value="%v">%v
		</div>`,
		this.CheckboxFieldName,
		this.CheckboxFieldName,
		this.CheckboxLabel)
}

func (this AddressStruct) GetHtml() string {
	return fmt.Sprintf(`
		<div>
		</div>`)
}

func (this EmailStruct) GetHtml() string {
	return fmt.Sprintf(`
		<div>
			<input type="email" name="%v">
		</div>`,
		this.EmailFieldName)
}

func (this HiddenStruct) GetHtml() string {
	return fmt.Sprintf(`

		`)
}

func (lc *FormStruct) UnmarshalJSON(data []byte) error {
	fmt.Println("IN UNMARSHAL")
	//Extract LayoutProperties from underlying Component
	// var tmp tmpLayoutStruct
	// if err := json.Unmarshal(data, &tmp); err != nil {
	// 	return err
	// }
	// Instantiate proper Component
	component, err := UnmarshalCustomValue(data, "componentType", "componentData",
		map[string]reflect.Type{
			"formText": reflect.TypeOf(FormTextStruct{}),
			"formDate": reflect.TypeOf(DateStruct{}),
			"formDropDown": reflect.TypeOf(DropDownStruct{}),
			"formCheckbox": reflect.TypeOf(CheckboxStruct{}),
			"formAddress": reflect.TypeOf(AddressStruct{}),
			"formEmailAddress": reflect.TypeOf(EmailStruct{}),
			"formHidden": reflect.TypeOf(HiddenStruct{}),
		})
	if err != nil {
		return err
	}

	// lc.Item = component
	fmt.Println(component)

	return nil
}

func UnmarshalCustomValue(data []byte, typeField, resultField string, customTypes map[string]reflect.Type) (ComponentData, error) {
	fmt.Println("IN CUSTOM UNMARSHAL")
	m := map[string]interface{}{}
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	//fmt.Println(m)
	valueBytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	typeName := m[typeField].(string)
	switch typeName {
		case "formText":
			var newObj FormTextStruct
			if err = json.Unmarshal(valueBytes, &newObj); err != nil {
				return nil, err
			}
			return newObj, nil
		case "formDate":
			var newObj DateStruct
			if err = json.Unmarshal(valueBytes, &newObj); err != nil {
				return nil, err
			}
			return newObj, nil
		case "formDropdown":
			var newObj DropDownStruct
			if err = json.Unmarshal(valueBytes, &newObj); err != nil {
				return nil, err
			}
			return newObj, nil
		case "formCheckbox":
			var newObj CheckboxStruct
			if err = json.Unmarshal(valueBytes, &newObj); err != nil {
				return nil, err
			}
			return newObj, nil
		case "formAddress":
			var newObj AddressStruct
			if err = json.Unmarshal(valueBytes, &newObj); err != nil {
				return nil, err
			}
			return newObj, nil
		case "formEmailAddress":
			var newObj EmailStruct
			if err = json.Unmarshal(valueBytes, &newObj); err != nil {
				return nil, err
			}
			return newObj, nil
		case "formHidden":
			var newObj HiddenStruct
			if err = json.Unmarshal(valueBytes, &newObj); err != nil {
				return nil, err
			}
			return newObj, nil
		}
	//var newObj component_data.ComponentData
	//if ty, found := customTypes[typeName]; found {
	//	newObj = reflect.New(ty).Interface().(component_data.ComponentData)
	//}
	return nil, nil
}