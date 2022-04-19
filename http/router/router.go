package router

import (
	"imageUploader/http/controllers"
	"imageUploader/util/container"

	"github.com/gorilla/mux"
)

// initiate router
func Init(ctr container.Containers) *mux.Router {

	r := mux.NewRouter()

	ctl := controllers.NewUploaderController(ctr)

	r.HandleFunc("/upload", ctl.UploadImage)

	return r
}
