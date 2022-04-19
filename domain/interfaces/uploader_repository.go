package interfaces

import (
	"context"
	"imageUploader/domain/entities"
)

type UploaderRepository interface {
	UpdateFileInfo(ctx context.Context, info entities.FileInfo) error
}
