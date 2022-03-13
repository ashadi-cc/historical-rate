package provider

import "history-rate/db/repo"

type Provider interface {
	Init() error
	GetRateRepo() repo.Rate
}
