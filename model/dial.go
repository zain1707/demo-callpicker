package model

import "encoding/xml"

type Dial struct {
	XMLName           xml.Name `json:"-" bson:"-" xml:"Dial"`
	Step              string   `json:"step" bson:"step" xml:"loop,attr,omitempty" `
	DestinationNumber string   `json:"destination_number" bson:"destination_number" xml:",chardata"`
	Timeout           string   `json:"timeout" bson:"timeout" xml:"timeout,attr,omitempty"`
	CallbackUrl       string   `json:"callback_url"  bson:"callback_url" xml:"callbackUrl,attr,omitempty" `
	CallbackMethod    string   `json:"callback_method" bson:"callback_method" xml:"callbackMethod,attr,omitempty"`
	GroupConfirmKey   string   `json:"group_confirm_key" bson:"group_confirm_key" xml:"groupConfirmKey,attr,omitempty"`
	GroupConfirmFile  string   `json:"group_confirm_file" bson:"group_confirm_file" xml:"groupConfirmFile,attr,omitempty"`
	Action            string   `json:"action" bson:"action" xml:"action,attr,omitempty" `
	Method            string   `json:"method" bson:"method" xml:"method,attr,omitempty"`
	CallerID          string   `json:"caller_id" bson:"caller_id" xml:"callerid,attr,omitempty"`
	Record            string   `json:"record" bson:"record" xml:"record,attr,omitempty"`
	TimeLimit         string   `json:"time_limit" bson:"time_limit" xml:"timeLimit,attr,omitempty"`
	DialMusic         string   `json:"dial_music" bson:"dial_music" xml:"dialMusic,attr,omitempty"`
	RecordAction      string   `json:"record_action" bson:"record_action" xml:"recordAction,attr,omitempty"`
}
