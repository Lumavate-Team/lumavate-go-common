package component_data

import (
  "fmt"
  "reflect"
  "encoding/json"
  _"html/template"
)

type FormItemStruct struct {
	ComponentData struct {
		Item ComponentData
	}
}

type FormStruct struct {
	FormItems []FormItemStruct
}

type FormTextStruct struct {
	ComponentData struct {
		TextLabel string `json:"textLabel"`
		TextFieldName string `json:"textFieldName"`
		TextRegExpression string `json:"textRegExpression"`
		TextRequired bool `json:"textRequired"`
		TextMaxLength int `json:"textMaxLength"`
		TextErrorText string `json:"textErrorText"`
		TextInputType string `json:"textInputType"`
	}
}

type DateStruct struct {
	ComponentData struct {
		DateLabel string `json:"dateLabel"`
		DateFieldName string `json:"dateFieldName"`
		DateRequired bool `json:"dateRequired"`
	}
}

type DropDownStruct struct {
	ComponentData struct {
		DropDownLabel string `json:"dropDownLabel"`
		DropDownFieldName string `json:"dropDownFieldName"`
		DropDownRequired bool `json:"dropDownRequired"`
		DropDownItems string //not really sure about this
	}
}

type CheckboxStruct struct {
	ComponentData struct {
		CheckboxLabel string `json:"checkboxLabel"`
		CheckboxFieldName string `json:"checkboxFieldName"`
	}
}

type AddressStruct struct {
	ComponentData struct {
		AddressLabel string `json:"addressLabel"`
		AddresFieldName string `json:"addressFieldName"`
		AddressRequired bool `json:"addressRequired"`
	}
}

type EmailStruct struct {
	ComponentData struct {
		EmailLabel string `json:"emailLabel"`
		EmailFieldName string `json:"emailFieldName"`
		EmailRequired bool `json:"emailRequired"`
	}
}

type HiddenStruct struct {
	ComponentData struct {
		HiddenLabel string `json:"hiddenLabel"`
		HiddenFieldName string `json:"hiddenFieldName"`
	}
}

func (this FormStruct) GetHtml() string {
	var html string

	for _, element := range this.FormItems {
		html = html + fmt.Sprintf(`
			<div class="nav-item nav-tile">
				%v
			</div>`,
			element.ComponentData.Item.GetHtml())
	}

	return fmt.Sprintf(`
		<form action="">
			%v
		</form>`,
		html)
}

func (this FormTextStruct) GetHtml() string {
	var required = ""
	if this.ComponentData.TextRequired == true {
		required = "required"
	}
	return fmt.Sprintf(`
		<div>
			<input class="pure-input-1" aria-label="%v" id="%v" name="%v" type="%v" placeholder="%v" %v />
		</div>`,
		this.ComponentData.TextLabel,
		this.ComponentData.TextFieldName,
		this.ComponentData.TextFieldName,
		this.ComponentData.TextInputType,
		this.ComponentData.TextLabel,
		required)
}

func (this DateStruct) GetHtml() string {
	return fmt.Sprintf(`
		<div>
			<input aria-label="%v" id="%v" name="%v" type="date">
		</div>`,
		this.ComponentData.DateLabel,
		this.ComponentData.DateFieldName,
		this.ComponentData.DateFieldName)
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
		this.ComponentData.DropDownFieldName,
		this.ComponentData.DropDownFieldName,
		this.ComponentData.DropDownLabel)
}

func (this CheckboxStruct) GetHtml() string {
	return fmt.Sprintf(`
		<div>
			<input type="checkbox" name="%v" value="%v">%v
		</div>`,
		this.ComponentData.CheckboxFieldName,
		this.ComponentData.CheckboxFieldName,
		this.ComponentData.CheckboxLabel)
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
		this.ComponentData.EmailFieldName)
}

func (this HiddenStruct) GetHtml() string {
	return fmt.Sprintf(`

		`)
}

func (fs *FormItemStruct) UnmarshalJSON(data []byte) error {
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

	fs.ComponentData.Item = component

	return nil
}

func UnmarshalCustomValue(data []byte, typeField, resultField string, customTypes map[string]reflect.Type) (ComponentData, error) {
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
