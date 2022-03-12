package provider

import "history-rate/db/repo"

type Provider interface {
	GetRateRepo() repo.Rate
}
