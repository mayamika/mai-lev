package model

type Entry struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Surname    string `json:"surname,omitempty"`
	University string `json:"university,omitempty"`
}
