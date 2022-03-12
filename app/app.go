package app

import (
	"context"

	"history-rate/db"
	_ "history-rate/db/repo/sqlite"
	services "history-rate/service"
	"history-rate/service/api"
	"history-rate/service/rate"
	"history-rate/service/scheduler"
)

const dbDriver = "sqlite"

func Run(ctx context.Context) error {
	storage, err := db.Open(dbDriver)
	if err != nil {
		return err
	}

	scv := scheduler.NewService()
	rsv := rate.NewService(storage, scv)
	apisvc := api.NewService(scv, rsv)

	return services.RunServices(ctx, scv, apisvc)
}
