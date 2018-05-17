package component_data

import (
)

type AddToHomeStruct struct {
	ShowAddToHome bool `json:"showAddToHome"`
	SkipFirst bool `json:"skipFirst"`
	StartDelay int `json:"startDelay"`
	Lifespan int `json:"lifespan"`
	DisplayCount int `json:"displayCount"`
	Message string `json:"message"`
}