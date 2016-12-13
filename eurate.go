package eurate

import (
	"encoding/xml"
	"fmt"
)

type Query struct {
	RatesList []Rate `xml:"Cube>Cube>Cube"`
}

type Rate struct {
	Currency string `xml:"currency,attr"`
    Rate float64 `xml:"rate,attr"`
}

func eurateXmlMapper(xmlString string) map[string]float64 {
	result :=  make(map[string]float64)
	q := Query{}
	err := xml.Unmarshal([]byte(xmlString), &q)

	if err != nil {
		fmt.Printf("error: %v", err)
		return result
	}

	for _, element := range q.RatesList {
		result[element.Currency] = element.Rate
	}

	return result
}
