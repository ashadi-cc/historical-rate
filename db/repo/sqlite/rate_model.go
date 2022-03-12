package sqlite

import (
	"history-rate/db/repo"

	"gorm.io/gorm"
)

type Rate struct {
	gorm.Model
	DateInt    int `gorm:"index"`
	DateString string
	Currency   string `gorm:"index"`
	Rate       float64
}

func (r Rate) ToRepoModel() repo.RateModel {
	return repo.RateModel{
		Date:     r.DateString,
		Currency: r.Currency,
		Rate:     r.Rate,
	}
}
