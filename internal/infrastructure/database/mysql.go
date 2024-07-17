package database

import (
	"database/sql"
	"file-cleaner/internal/domain/entities"
	"time"

	"github.com/go-sql-driver/mysql"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQLConnection(host, user, password, dbName string) (*MySQL, error) {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   user,
		Passwd: password,
		Net:    "tcp",
		Addr:   host,
		DBName: dbName,
		Params: map[string]string{
			"allowNativePasswords": "true",
			"parseTime":            "true",
		},
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	return &MySQL{db: db}, nil
}

func (m *MySQL) GetExpiredFiles() ([]entities.File, error) {
	now := time.Now()
	query := "SELECT id, bucket, s3_key, expires_at FROM files WHERE expires_at < ?"
	rows, err := m.db.Query(query, now)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []entities.File
	for rows.Next() {
		var file entities.File
		err := rows.Scan(&file.ID, &file.Bucket, &file.S3Key, &file.ExpiresAt)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	return files, nil
}

func (m *MySQL) DeleteFileRecord(id int64) error {
	_, err := m.db.Exec("DELETE FROM files WHERE id = ?", id)
	return err
}
