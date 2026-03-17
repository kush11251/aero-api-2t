package models

type Weather struct {
	Description string `json:"description"`
	Temperature float64 `json:"temperature"`
}
