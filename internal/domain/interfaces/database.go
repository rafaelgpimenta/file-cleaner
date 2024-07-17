package interfaces

import "file-cleaner/internal/domain/entities"

type Database interface {
	GetExpiredFiles() ([]entities.File, error)
	DeleteFileRecord(key int64) error
}
