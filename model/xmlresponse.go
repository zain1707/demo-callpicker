package model

import "encoding/xml"

type XMLResponse struct {
	XMLName  xml.Name      `xml:"Response"`
	Text     string        `xml:",chardata"`
	Elements []interface{} `xml:",any"`
}
