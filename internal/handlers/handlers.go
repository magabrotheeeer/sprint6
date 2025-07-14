package handlers

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
	
	"github.com/magabrotheeeer/sprint6/internal/service"
)

func FirstHandler(w http.ResponseWriter, r *http.Request) {
	projectPath, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filename := filepath.Join(projectPath, "index.html")
	data, err := os.ReadFile(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // 10MB
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "error when receiving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	err = os.Mkdir("C:\\Users\\akhilgovmb\\Desktop\\sprint6\\uploads", 0755)
	if err != nil && !errors.Is(err, os.ErrExist) {
		http.Error(w, "error when creating the folder", http.StatusInternalServerError)
		return
	} 

	root, err := os.OpenRoot("C:\\Users\\akhilgovmb\\Desktop\\sprint6\\uploads")
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	defer root.Close()
	buffer, err := os.Create(filepath.Join("C:\\Users\\akhilgovmb\\Desktop\\sprint6\\uploads", time.Now().UTC().String()))
	if err != nil {
		http.Error(w, "error when creating the file", http.StatusInternalServerError)
	}
	defer buffer.Close()

	_, err = io.Copy(buffer, file)
	if err != nil {
		http.Error(w, "error when copying the file", http.StatusInternalServerError)
		return
	}
	data, err := os.ReadFile(buffer.Name())
	if err != nil {
		http.Error(w, "error when reading the file", http.StatusInternalServerError)
		return
	}
	res, err := service.DefineText(string(data))
	if err != nil {
		http.Error(w, "error when handling the data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}
