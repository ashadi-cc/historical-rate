package sqlite

import (
	"history-rate/db"
	"history-rate/db/repo"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbName = "./data/rate.db"

type repoProvider struct {
	db          *gorm.DB
	rateRepo    repo.Rate
	isConnected bool
}

func (p *repoProvider) Init() error {
	if p.isConnected {
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
	p.rateRepo = NewRateRepo(p.db)
	p.isConnected = true

	return nil
}

func (p *repoProvider) GetRateRepo() repo.Rate {
	return p.rateRepo
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Rate{})
}

func init() {
	pr := &repoProvider{}
	db.Register("sqlite", pr)
}
