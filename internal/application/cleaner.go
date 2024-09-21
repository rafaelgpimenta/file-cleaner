package application

import (
	"context"
	"file-cleaner/internal/domain/interfaces"
	"file-cleaner/internal/lib/logger"
)

type Cleaner struct {
	db interfaces.Database
	s3 interfaces.Storage
}

func NewCleaner(db interfaces.Database, s3 interfaces.Storage) *Cleaner {
	return &Cleaner{
		db: db,
		s3: s3,
	}
}

func (c *Cleaner) CleanExpiredFiles(ctx context.Context) error {
	files, err := c.db.GetExpiredFiles()
	if err != nil {
		return err
	}

	logger.Info().Ctx(ctx).
		Int("files", len(files)).
		Msg("Total expired files")

	for _, file := range files {
		logger.Info().Ctx(ctx).
			Int64("ID", file.ID).
			Str("Bucket", file.Bucket).
			Str("s3Key", file.S3Key).
			Msg("Removing expired file path")

		err = c.s3.DeleteFile(file.Bucket, file.S3Key)
		if err != nil {
			return err
		}

		logger.Info().Ctx(ctx).
			Int64("ID", file.ID).
			Str("Bucket", file.Bucket).
			Str("s3Key", file.S3Key).
			Msg("Removed expired file path from S3")

		err = c.db.DeleteFileRecord(file.ID)
		if err != nil {
			return err
		}

		logger.Info().Ctx(ctx).
			Int64("ID", file.ID).
			Str("Bucket", file.Bucket).
			Str("s3Key", file.S3Key).
			Msg("Removed expired file path from Database")
	}

	return nil
}
