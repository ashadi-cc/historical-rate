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

func (p *repoProvider) GetRateRepo() repo.Rate {
	return NewRateRepo(p.db)
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Rate{})
}

func init() {
	idb, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := autoMigrate(idb); err != nil {
		panic(err)
	}

	pr := &repoProvider{db: idb}
	db.Register("sqlite", pr)
}
