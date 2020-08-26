package model

type Brand struct {
	Brandname string   `json:"brandname" bson:"brandname"`
	Models    []string `json:"automodels omitempty" bson:"automodels"`
}
