package model

type tModel struct {
	Model `bson:",inline"`

	Name string `json:"name,omitempty"`
}
