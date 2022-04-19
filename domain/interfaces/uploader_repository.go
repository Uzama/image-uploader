package interfaces

import (
	"context"
	"imageUploader/domain/entities"
)

// Dependency Injection
type UploaderRepository interface {
	UpdateFileInfo(ctx context.Context, info entities.FileInfo) error
}
