package model

type Brand struct {
	Brandname string   `json:"brandname" bson:"brandname"`
	Models    []string `json:"models omitempty" bson:"models"`
}
