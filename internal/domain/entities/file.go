package entities

import "time"

type File struct {
	ID        int64     `json:"id"`
	Bucket    string    `json:"bucket"`
	S3Key     string    `json:"s3Key"`
	ExpiresAt time.Time `json:"expiresAt"`
}
