package models

type Stat struct {
	Name string `json:"name"`
	Value int `json:"value"`
	Modifier int `json:"mod"`
}