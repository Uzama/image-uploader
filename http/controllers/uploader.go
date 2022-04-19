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

type UploaderController struct {
	usecase usecases.UploaderUsecase
}

func NewUploaderController(ctr container.Containers) UploaderController {
	usecase := usecases.NewUploaderUsecase(ctr.Repositories.Uploader)

	ctl := UploaderController{
		usecase: usecase,
	}

	return ctl
}

func (ctl UploaderController) UploadImage(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		t, err := template.ParseFiles("html/upload.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		t.Execute(w, globals.AUTH_KEY)

	} else if r.Method == "POST" {

		ctx := r.Context()

		t, err := template.ParseFiles("html/response.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		r.ParseMultipartForm(8 << 20)

		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			t.Execute(w, "invalid file")
			return
		}

		cType := handler.Header["Content-Type"]

		key := strings.Split(cType[0], "/")[0]

		if key != "image" {
			w.WriteHeader(http.StatusForbidden)
			t.Execute(w, "only image file is supported")
			return
		}

		if handler.Size > 8*1024*1024 {
			w.WriteHeader(http.StatusForbidden)
			t.Execute(w, "maximum 8MB is supported")
			return
		}

		defer file.Close()

		f, err := os.OpenFile("images/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			t.Execute(w, "internal server error")
			return
		}

		defer f.Close()

		_, err = io.Copy(f, file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			t.Execute(w, "internal server error")
			return
		}

		info := entities.FileInfo{
			FileName:    handler.Filename,
			Path:        "images/" + handler.Filename,
			Size:        handler.Size,
			ContentType: cType[0],
		}

		err = ctl.usecase.UpdateFileInfo(ctx, info)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			t.Execute(w, "internal server error")
			return
		}

		w.WriteHeader(http.StatusAccepted)
		t.Execute(w, "file is uploaded")
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}
