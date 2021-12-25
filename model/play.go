package model

import "encoding/xml"

type Play struct {
	XMLName xml.Name `json:"-" bson:"-" xml:"Play"`
	Loop    string   `json:"loop" bson:"loop" xml:"loop,attr,omitempty"`
	Value   string   `json:"value" bson:"value" xml:",chardata"`
	Step    string   `json:"step" bson:"step" xml:"-"`
}
