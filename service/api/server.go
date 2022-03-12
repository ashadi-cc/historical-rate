package api

import (
	"context"
	"net/http"

	"history-rate/service/rate"
	"history-rate/service/scheduler"

	"github.com/labstack/echo/v4"
)

const port = "8001"

func NewService(sc scheduler.ISchedulerService, rt rate.RateService) *apiServer {
	return &apiServer{
		sc: sc,
		rt: rt,
	}
}

type apiServer struct {
	sc scheduler.ISchedulerService
	rt rate.RateService
}

func (api *apiServer) Run(ctx context.Context) error {
	if err := api.rt.Run(ctx); err != nil {
		return err
	}

	return api.serveApi()
}

func (api *apiServer) serveApi() error {
	e := echo.New()
	e.GET("/rates/latest", func(c echo.Context) error {
		rates, err := api.rt.GetLatestRate()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, rates)
	})

	return e.Start(":" + port)
}
