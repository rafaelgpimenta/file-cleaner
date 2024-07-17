package crontab

import (
	"file-cleaner/internal/lib/logger"

	"github.com/robfig/cron/v3"
)

func RunEveryMinute(job func()) {
	c := cron.New()
	_, err := c.AddFunc("*/1 * * * *", job)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Error while adding cron task")
	}

	c.Start()
}
