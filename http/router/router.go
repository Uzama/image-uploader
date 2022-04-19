package router

import (
	"imageUploader/http/controllers"
	"imageUploader/util/container"
	"net/http"

	"github.com/gorilla/mux"
)

func Init(ctr container.Containers) *mux.Router {

	r := mux.NewRouter()

	ctl := controllers.NewUploaderController(ctr)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//
	}).Methods(http.MethodGet)

	r.HandleFunc("/upload", ctl.UploadImage)

	return r
}
