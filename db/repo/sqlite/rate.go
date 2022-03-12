package sqlite

import (
	"history-rate/db/repo"

	"gorm.io/gorm"
)

type rateRepo struct {
	db *gorm.DB
}

func NewRateRepo(db *gorm.DB) repo.Rate {
	r := &rateRepo{db: db}
	return r
}

func (repo *rateRepo) Save(rates []repo.RateModel) error {
	for _, r := range rates {
		rate := Rate{
			DateString: r.Date,
			DateInt:    r.DateInt(),
			Currency:   r.Currency,
			Rate:       r.Rate,
		}
		if err := repo.db.Where("date_int = ? and currency = ?", r.DateInt(), r.Currency).
			Attrs(rate).
			FirstOrCreate(&rate).Error; err != nil {
			return err
		}
	}
	return nil
}

func (repo *rateRepo) Latest() (results []repo.RateModel, err error) {
	rate := Rate{}
	if err := repo.db.Order("date_int desc").First(&rate).Error; err != nil {
		return nil, err
	}

	rates := []Rate{}
	if err := repo.db.Where("date_int = ?", rate.DateInt).Order("currency asc").Find(&rates).Error; err != nil {
		return nil, err
	}

	for _, r := range rates {
		results = append(results, r.ToRepoModel())
	}

	return results, nil
}
