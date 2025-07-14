package handlers

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHtml(w http.ResponseWriter, r *http.Request) {

	filename := filepath.Join("C:\\Users\\akhilgovmb\\Desktop\\sprint6\\", "index.html")
	data, err := os.ReadFile(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	// скачиваем файл
	r.ParseMultipartForm(10 << 20) // 10MB
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "error when receiving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// создаем файл
	err = os.Mkdir("C:\\Users\\akhilgovmb\\Desktop\\sprint6\\uploads", 0755)
	if err != nil && !errors.Is(err, os.ErrExist) {
		http.Error(w, "error when creating the folder", http.StatusInternalServerError)
		return
	} 
	
	// ограничиваем доступ к файловой системе
	root, err := os.OpenRoot("C:\\Users\\akhilgovmb\\Desktop\\sprint6\\uploads")
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	defer root.Close()

	// создаем локальный файл
	buffer, err := os.Create(filepath.Join("C:\\Users\\akhilgovmb\\Desktop\\sprint6\\uploads\\", handler.Filename))
	if err != nil {
		http.Error(w, "error when creating the file", http.StatusInternalServerError)
		return
	}
	defer buffer.Close()

	// копируем скачанный файл в локальный файл
	_, err = io.Copy(buffer, file)
	if err != nil {
		http.Error(w, "error when copying the file", http.StatusInternalServerError)
		return
	}

	// считываем данные из локального файла
	data, err := os.ReadFile(buffer.Name())
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

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}