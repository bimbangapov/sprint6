package handlers

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

type Handler struct {
	Logger *log.Logger
}

func (h *Handler) HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.Logger.Printf("Error Method not allowed: %s", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := os.Mkdir("uploads", 0755)
	// если директория существует, то никак не реагируем
	if err != nil && !errors.Is(err, os.ErrExist) {
		h.Logger.Println("Internal server error: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	filename, _, err := r.FormFile("myFile")
	if err != nil {
		h.Logger.Println("Error: ", err)
		http.Error(w, ("File must be at key myFile " + err.Error()), http.StatusBadRequest)
		return
	}
	defer filename.Close()

	filebody, err := io.ReadAll(filename)
	if err != nil {
		h.Logger.Println("Internal server error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	convertFilebody := service.Convert(string(filebody))

	uploadsFile, err := os.Create("uploads/" + time.Now().UTC().String())
	if err != nil {
		h.Logger.Println("Internal server error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer uploadsFile.Close()

	_, err = uploadsFile.Write([]byte(convertFilebody))
	if err != nil {
		h.Logger.Println("Internal server error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	h.Logger.Println("POST /upload")
	outputString := fmt.Sprintf("Содержимое файла: \n%s", convertFilebody)
	fmt.Println("file body", string(filebody))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outputString))

}

func (h *Handler) GetHome(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("index.html")
	if err != nil {
		h.Logger.Println("Internal server error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := ""
	for scanner.Scan() {
		result += scanner.Text()
	}
	h.Logger.Println("GET /")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
