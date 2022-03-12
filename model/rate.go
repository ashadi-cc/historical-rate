package model

import (
	"encoding/xml"

	"history-rate/db/repo"
)

type Cube struct {
	XMLName  xml.Name `xml:"Cube"`
	Currency string   `xml:"currency,attr"`
	Rate     float64  `xml:"rate,attr"`
}

type Cubes struct {
	XMLName xml.Name `xml:"Cube"`
	Date    string   `xml:"time,attr"`
	Cubes   []Cube   `xml:"Cube"`
}

type RootCube struct {
	XMLName xml.Name `xml:"Cube"`
	Cubes   []Cubes  `xml:"Cube"`
}

type RateSender struct {
	XMLName xml.Name `xml:"Sender"`
	Name    string   `xml:"name"`
}

type Rate struct {
	XMLName xml.Name   `xml:"Envelope"`
	Subject string     `xml:"subject"`
	Sender  RateSender `xml:"Sender"`
	Cube    RootCube   `xml:"Cube"`
}

func (r Rate) ToModel() []repo.RateModel {
	rates := make([]repo.RateModel, 0)
	for _, cube := range r.Cube.Cubes {
		date := cube.Date
		for _, item := range cube.Cubes {
			rates = append(rates, repo.RateModel{
				Date:     date,
				Currency: item.Currency,
				Rate:     item.Rate,
			})
		}
	}

	return rates
}
