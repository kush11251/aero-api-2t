package controllers

import (
	"aero-api/src/models"
	"aero-api/src/utils"
)

func GetForecast() models.Weather {
	return models.Weather{Description: "clear skies", Temperature: 25.0}
}