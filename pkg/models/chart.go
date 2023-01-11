package models

type Chart struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	DefaultValues string `json:"default_values"`
}
