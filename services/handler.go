package services

import (
	"ascii-art-web/ascii_art/run"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func AsciiArtPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var tmplPath string
		// Определяем путь к шаблону в зависимости от запрошенного URL
		if r.URL.Path == "/ascii-art-page.html" {
			tmplPath = filepath.Join("templates", "ascii-art-page.html")
		} else if r.URL.Path == "/index.html" {
			tmplPath = filepath.Join("templates", "index.html")
		} else if r.URL.Path == "/" {
			tmplPath = filepath.Join("templates", "index.html")
		} else {
			// Указываем путь к файлу 404.html
			tmplPath = filepath.Join("templates", "404.html")
			// Открываем файл 404.html
			file, err := ioutil.ReadFile(tmplPath)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				log.Printf("Error reading file: %v", err)
				return
			}
			// Устанавливаем заголовок ответа для типа контента
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			// Отправляем статус 404 Not Found и содержимое файла 404.html
			w.WriteHeader(http.StatusNotFound)
			w.Write(file)
			return
		}
		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error parsing template: %v", err)
			return
		}
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error executing template: %v", err)
			return
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Обработка POST запросов
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		var requestDataStruct struct {
			Text   string `json:"text"`
			Banner string `json:"banner"`
		}
		if err := json.Unmarshal(body, &requestDataStruct); err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}
		asciiArt, status := run.Run(requestDataStruct.Text, requestDataStruct.Banner)
		if status == 400 {
			http.Error(w, "Failed to generate ASCII art", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, asciiArt)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
