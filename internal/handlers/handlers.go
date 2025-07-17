package handlers

import (
	"io"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHtml(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "error: incorrect method", http.StatusInternalServerError)
		return
	}

	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "error when trying to open index.html", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, "error when trying to response", http.StatusInternalServerError)
		return
	} 
}

func Upload(w http.ResponseWriter, r *http.Request) {
	// скачиваем файл
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "error: incorrect method", http.StatusInternalServerError)
		return
	}
	r.ParseMultipartForm(10 << 20) // 10MB
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "error when receiving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()


	// считываем данные из локального файла
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "error when reading the file", http.StatusInternalServerError)
		return
	}

	// преобразуем полученные данные
	res, err := service.DefineText(string(data))
	if err != nil {
		http.Error(w, "error when handling the data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(res))
	if err != nil {
		http.Error(w, "error when trying to response", http.StatusInternalServerError)
		return
	} 
}