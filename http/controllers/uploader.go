package controllers

import (
	"html/template"
	"imageUploader/domain/entities"
	"imageUploader/domain/globals"
	"imageUploader/domain/usecases"
	"imageUploader/util/container"
	"io"
	"net/http"
	"os"
	"strings"
)

// controller layer
type UploaderController struct {
	usecase usecases.UploaderUsecase
}

// create new controller
func NewUploaderController(ctr container.Containers) UploaderController {
	usecase := usecases.NewUploaderUsecase(ctr.Repositories.Uploader)

	ctl := UploaderController{
		usecase: usecase,
	}

	return ctl
}

// handler for upload image
func (ctl UploaderController) UploadImage(w http.ResponseWriter, r *http.Request) {

	// if method is GET, then return the form
	if r.Method == "GET" {

		t, err := template.ParseFiles("html/upload.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		t.Execute(w, globals.AUTH_KEY)

		// if method is POST, process the uploaded file
	} else if r.Method == "POST" {

		ctx := r.Context()

		// parsing response file
		t, err := template.ParseFiles("html/response.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		r.ParseMultipartForm(8 << 20)

		// get the file & other info
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			t.Execute(w, "invalid file")
			return
		}

		// get the content type
		cType := handler.Header["Content-Type"]

		// validate to allow only image file
		key := strings.Split(cType[0], "/")[0]
		if key != "image" {
			w.WriteHeader(http.StatusForbidden)
			t.Execute(w, "only image file is supported")
			return
		}

		// validate to allow only 8MB
		if handler.Size > 8*1024*1024 {
			w.WriteHeader(http.StatusForbidden)
			t.Execute(w, "maximum 8MB is supported")
			return
		}

		defer file.Close()

		// creating a file in local file system
		f, err := os.OpenFile("images/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			t.Execute(w, "internal server error")
			return
		}

		defer f.Close()

		// copy uploded file into above created file
		_, err = io.Copy(f, file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			t.Execute(w, "internal server error")
			return
		}

		// prepare file information for update into database
		info := entities.FileInfo{
			FileName:    handler.Filename,
			Path:        "images/" + handler.Filename,
			Size:        handler.Size,
			ContentType: cType[0],
		}

		// update into database
		err = ctl.usecase.UpdateFileInfo(ctx, info)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			t.Execute(w, "internal server error")
			return
		}

		// success message
		w.WriteHeader(http.StatusAccepted)
		t.Execute(w, "file is uploaded")
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}
