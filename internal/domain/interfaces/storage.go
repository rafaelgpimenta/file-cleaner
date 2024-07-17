package interfaces

type Storage interface {
	DeleteFile(bucket string, key string) error
}
