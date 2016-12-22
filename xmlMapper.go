package eurate

import (
	"encoding/xml"
	"fmt"
)



type Query struct {
	DayList []DailyRate `xml:"Cube>Cube"`
}

type DailyRate struct {
	Date string `xml:"time,attr"`
	RatesList []Rate `xml:"Cube"`
}

type Rate struct {
	Currency string `xml:"currency,attr"`
	Rate float64 `xml:"rate,attr"`
}


type RateForCurrencies map[string]float64
type RatesForDates map[string]RateForCurrencies

func XmlExtractor(xmlString string) []DailyRate {
	q := Query{}
	err := xml.Unmarshal([]byte(xmlString), &q)

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	return q.DayList
}

func XmlMapper(xmlString string) RatesForDates {
	result :=  make(RatesForDates)

	dayList := XmlExtractor(xmlString)

	for _, day := range dayList {
		dailyRates := make(RateForCurrencies)
		result[day.Date] = dailyRates
		for _, rate := range day.RatesList {
			dailyRates[rate.Currency] = rate.Rate
		}
	}

	return result
}
