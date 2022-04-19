package repositories

import (
	"context"
	"database/sql"
	"imageUploader/domain/entities"
)

// repository Layer
type UploaderRepository struct {
	db *sql.DB
}

// create new repository
func NewUploaderRepository(db *sql.DB) *UploaderRepository {
	repo := &UploaderRepository{
		db: db,
	}

	return repo
}

// update file information in database
// db-name: uploader, table: uploads
func (repo UploaderRepository) UpdateFileInfo(ctx context.Context, info entities.FileInfo) error {
	query := `INSERT INTO uploads (file_name, file_path, content_type, size) VALUES (?, ?, ?, ?);`

	// preparing the statement
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	// executing the statement
	_, err = stmt.ExecContext(ctx, info.FileName, info.Path, info.ContentType, info.Size)
	if err != nil {
		return err
	}

	return nil
}
