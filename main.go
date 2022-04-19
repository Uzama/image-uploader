package main

import (
	"context"
	"imageUploader/bootstrap"
)

func main() {
	ctx := context.Background()

	bootstrap.Start(ctx)
}
