package sqlite

import (
	"history-rate/db"
	"history-rate/db/repo"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbName = "./data/rate.db"

type repoProvider struct {
	db *gorm.DB
}

func (p *repoProvider) Init() error {
	if p.db != nil {
		return nil
	}

	idb, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := autoMigrate(idb); err != nil {
		return err
	}

	p.db = idb

	return nil
}

func (p *repoProvider) GetRateRepo() repo.Rate {
	return NewRateRepo(p.db)
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Rate{})
}

func init() {
	pr := &repoProvider{}
	db.Register("sqlite", pr)
}
