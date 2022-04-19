package usecases

import (
	"context"
	"imageUploader/domain/entities"
	"imageUploader/domain/interfaces"
)

// usecase layer
type UploaderUsecase struct {
	repo interfaces.UploaderRepository
}

// create new usecase
func NewUploaderUsecase(repo interfaces.UploaderRepository) UploaderUsecase {
	usecase := UploaderUsecase{
		repo: repo,
	}

	return usecase
}

func (usecase UploaderUsecase) UpdateFileInfo(ctx context.Context, info entities.FileInfo) error {
	return usecase.repo.UpdateFileInfo(ctx, info)
}
