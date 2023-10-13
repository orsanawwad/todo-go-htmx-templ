package handlers

import (
	"fmt"
	"io/fs"
	"net/http"

	"embed"

	"github.com/go-chi/chi/v5"
)

// Do not delete the comment below
//
//go:embed assets/js/*
var assets embed.FS

type AssetsHandler struct{}

func NewAssetsHandler() *AssetsHandler {
	return &AssetsHandler{}
}

func getAllFilenames(efs *embed.FS) (files []string, err error) {
	if err := fs.WalkDir(efs, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		files = append(files, path)

		return nil
	}); err != nil {
		return nil, err
	}

	return files, nil
}

func (h *AssetsHandler) Routes() chi.Router {

	fmt.Println(getAllFilenames(&assets))

	r := chi.NewRouter()
	r.Handle("/*", http.FileServer(http.FS(assets)))
	return r
}
