package model

import "encoding/xml"

type Record struct {
	XMLName     xml.Name `json:"-"  xml:"Record"`
	Loop        string   `json:"loop" bson:"loop" xml:"loop,attr,omitempty"`
	Value       string   `json:"value" bson:"value" xml:",chardata"`
	MaxLength   string   `json:"max_length" bson:"max_length" xml:"maxLength,attr,omitempty"`
	Action      string   `json:"action" bson:"action" xml:"action,attr,omitempty"`
	Method      string   `json:"method" bson:"method" xml:"method,attr,omitempty"`
	PlayBeep    string   `json:"play_beep" bson:"play_beep" xml:"playBeep,attr,omitempty"`
	Timeout     string   `json:"timeout" bson:"timeout" xml:"timeout,attr,omitempty"`
	FinishOnKey string   `json:"finish_on_key" bson:"finish_on_key" xml:"finishOnKey,attr,omitempty"`
	TrimSilence string   `json:"trim_silence" bson:"trim_silence" xml:"trimSilence,attr,omitempty"`
	Background  string   `json:"background" bson:"background" xml:"background,attr,omitempty"`
	FileFormat  string   `json:"file_format" bson:"file_format" xml:"fileFormat,attr,omitempty"`
	Step        string   `json:"step" bson:"step" xml:"-"`
}
