package util

import (
	"encoding/xml"

	"history-rate/model"
)

// ParseRateFromXML convert XML to Rate data model
func ParseRateFromXML(data []byte) (model.Rate, error) {
	var rate model.Rate
	if err := xml.Unmarshal(data, &rate); err != nil {
		return rate, err
	}

	return rate, nil
}
