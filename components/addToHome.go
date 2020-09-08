package components

import (
  "github.com/Lumavate-Team/lumavate-go-common/properties"
)

func GetAddToHomeProperties() [] properties.PropertyType {
  props := [] properties.PropertyType {}

  props = append(props, &properties.PropertyToggle{
	    &properties.PropertyBase{"showAddToHome", "Add To Home Screen", "Add To Home Screen Prompt", "Enable Prompt", "Display a message prompt to users on iOS devices that explains how to add this app to their home screen.", ""},
	    false})
  props = append(props, &properties.PropertyToggle{
	    &properties.PropertyBase{"skipFirst", "Add To Home Screen", "Prompt Action on First Visit", "Skip first visit", "This determines if the prompt will be displayed on the users first visit. If you would like to prompt new and returning users, then toggle off. If you would like to only prompt returning users, then toggle on.", ""},
	    true})

  props = append(props, &properties.PropertyNumeric{
      &properties.PropertyBase{"startDelay", "Add To Home Screen", "Prompt Time Configuration", "Delay Time", "Add to Home Screen Prompt will display X seconds after the page has fully loaded. If set to 0, the prompt will display on page load.", ""}, 0, properties.PropertyOptionsNumeric{ Min: 0, Max: 30}})

  props = append(props, &properties.PropertyNumeric{
      &properties.PropertyBase{"lifespan", "Add To Home Screen", "Prompt Time Configuration", "Lifespan", "Lifespan is the duration of the prompt message in seconds. We recommend giving your users enough time to read your prompt message; max duration is 30 seconds. If you would like the message to persist with no time expiration, set the duration to 0; (0 = infinity).", ""}, 0, properties.PropertyOptionsNumeric{ Min: 0, Max: 30}})

  props = append(props, &properties.PropertyNumeric{
      &properties.PropertyBase{"displayCount", "Add To Home Screen", "Prompt Time Configuration", "Display Count", "Display Count is the number of days you would like this prompt to appear until a user saves the app to their home screen. By default, the prompt message will display once per day until the user saves the app to their home screen. If you would like the prompt message to display until action is taken, set Display Count to 0.", ""}, 0, properties.PropertyOptionsNumeric{ Min: 0, Max: 30}})

  props = append(props, &properties.PropertyText{
    	&properties.PropertyBase{"message", "Add To Home Screen", "Prompt Message", "Message", "This is the message that will be displayed to your users. Your message should be concise and explain how to add the app to the home screen. The prompt message will display your fav icon—configured in the App Manifest on your Home Page.", ""},
    	"Add this app to your home screen!", properties.PropertyOptionsText{}})

  return props
}
