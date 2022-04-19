package container

import (
	"database/sql"
	"imageUploader/domain/interfaces"
)

type Containers struct {
	Adapters     Adapters
	Repositories Repositories
}

type Adapters struct {
	Db *sql.DB
}

type Repositories struct {
	Uploader interfaces.UploaderRepository
}
