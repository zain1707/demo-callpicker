package model

import "encoding/xml"

type Say struct {
	XMLName xml.Name `json:"-"  xml:"Say"`
	Loop    string   `json:"loop" bson:"loop" xml:"loop,attr,omitempty"`
	Voice   string   `json:"voice" bson:"voice" xml:"voice,attr,omitempty" `
	Value   string   `json:"value" bson:"value" xml:",chardata"`
	Step    string   `json:"step" bson:"step" xml:"-"`
}
