package repositories

import (
	"context"
	"database/sql"
	"imageUploader/domain/entities"
)

type UploaderRepository struct {
	db *sql.DB
}

func NewUploaderRepository(db *sql.DB) *UploaderRepository {
	repo := &UploaderRepository{
		db: db,
	}

	return repo
}

func (repo UploaderRepository) UpdateFileInfo(ctx context.Context, info entities.FileInfo) error {
	query := `INSERT INTO uploads (file_name, file_path, content_type, size) VALUES (?, ?, ?, ?);`

	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, info.FileName, info.Path, info.ContentType, info.Size)
	if err != nil {
		return err
	}

	return nil
}
