package repo

import (
	"strconv"
	"strings"
)

type Rate interface {
	Save(rates []RateModel) error
	Latest() ([]RateModel, error)
}

type RateModel struct {
	Date     string  `json:"date"`
	Currency string  `json:"currency"`
	Rate     float64 `json:"rate"`
}

func (r RateModel) DateInt() int {
	v := strings.ReplaceAll(r.Date, "-", "")
	n, _ := strconv.Atoi(v)
	return n
}
