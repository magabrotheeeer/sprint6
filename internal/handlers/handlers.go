package handlers

import (
	"net/http"
	"path/filepath"
	"os"
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

func SecondHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // 10MB
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "error when receiving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	os.Getwd()
}
