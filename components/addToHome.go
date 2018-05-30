package components

import (
  "github.com/Lumavate-Team/lumavate-go-common/properties"
)

func GetAddToHomeProperties() [] properties.PropertyType {
  props := [] properties.PropertyType {}

  props = append(props, &properties.PropertyToggle{
	    &properties.PropertyBase{"showAddToHome", "Add To Home Screen", "Enable Add To Home Screen", "Add To Home Screen Prompt", ""},
	    true})
  props = append(props, &properties.PropertyToggle{
	    &properties.PropertyBase{"skipFirst", "Add To Home Screen", "On First Visit", "Skip first visit", "Show only to returning visitors (ie: skip the first time you visit)"},
	    true})

  props = append(props, &properties.PropertyNumeric{
      &properties.PropertyBase{"startDelay", "Add To Home Screen", "Numeric Options", "Delay Time", "Display the message after that many seconds from page load"}, 0, properties.PropertyOptionsNumeric{ Min: 0, Max: 32}})

  props = append(props, &properties.PropertyNumeric{
      &properties.PropertyBase{"lifespan", "Add To Home Screen", "Numeric Options", "Lifespan", "Life of the message in seconds"}, 0, properties.PropertyOptionsNumeric{ Min: 0, Max: 32}})

  props = append(props, &properties.PropertyNumeric{
      &properties.PropertyBase{"displayCount", "Add To Home Screen", "Numeric Options", "Display Count", "Message will display once per day until end user adds to home screen. Display Count represents the number of days the message will be shown. (0: no limit)"}, 0, properties.PropertyOptionsNumeric{ Min: 0, Max: 32}})

  props = append(props, &properties.PropertyText{
    	&properties.PropertyBase{"message", "Add To Home Screen", "Prompt Message", "Message", "Message to be displayed on prompt."},
    	"Delay", properties.PropertyOptionsText{}})

  return props
}