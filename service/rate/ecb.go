package rate

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"history-rate/db/provider"
	"history-rate/db/repo"
	"history-rate/model"
	services "history-rate/service"
	"history-rate/service/scheduler"
	"history-rate/util"
)

const rateUrl = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml"

type RateService interface {
	services.Service
	GetLatestRate() ([]repo.RateModel, error)
}

func NewService(db provider.Provider, sc scheduler.ISchedulerService) RateService {
	return &ecb{
		db: db,
		sc: sc,
	}
}

type ecb struct {
	db provider.Provider
	sc scheduler.ISchedulerService
}

func (e *ecb) GetLatestRate() ([]repo.RateModel, error) {
	return e.db.GetRateRepo().Latest()
}

func (e *ecb) Run(ctx context.Context) error {
	e.sc.AddTask(e.saveRate, time.Second*10, 0)
	return nil
}

func (e *ecb) saveRate(ctx context.Context) error {
	log.Println("[RATE]", "request rate and store to database")
	rates, err := e.requestRate()
	if err != nil {
		return err
	}
	return e.db.GetRateRepo().Save(rates.ToModel())
}

func (e *ecb) requestRate() (model.Rate, error) {
	var rate model.Rate
	resp, err := doRequest(rateUrl)
	if err != nil {
		return rate, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return rate, err
	}

	return util.ParseRateFromXML(body)
}
