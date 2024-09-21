package main

import (
	"context"
	"file-cleaner/internal/application"
	"file-cleaner/internal/infrastructure/database"
	"file-cleaner/internal/infrastructure/storage"
	"file-cleaner/internal/lib/config"
	"file-cleaner/internal/lib/crontab"
	"file-cleaner/internal/lib/logger"

	"github.com/google/uuid"
)

type AWSConfig struct {
	Endpoint string `mapstructure:"endpoint"`
	Region   string `mapstructure:"region"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbName"`
}

func main() {
	// MySQL init
	dbConfig := config.Get[DBConfig]("mySQL")
	db, err := database.
		NewMySQLConnection(dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DbName)
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("Error connecting DB")
	}
	logger.Info().
		Msg("DB connected")

	// S3 init
	awsConfig := config.Get[AWSConfig]("aws")

	s3, err := storage.NewS3Client(awsConfig.Region, awsConfig.Endpoint)
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("Error connecting S3")
	}
	logger.Info().
		Msg("S3 connected")

	// init cleaner
	cleaner := application.NewCleaner(db, s3)

	crontab.RunEveryMinute(func() {
		traceId := uuid.New().String()
		ctx := context.WithValue(context.Background(), "traceId", traceId)

		err := cleaner.CleanExpiredFiles(ctx)
		if err != nil {
			logger.Error().Ctx(ctx).Err(err).Msg("Error cleaning expired files")
		}
	})

	// Keep app running (waits cronjobs)
	select {}
}
