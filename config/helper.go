package config

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

//Render é usada para executar os templates
func Render(w http.ResponseWriter, filename string, data interface{}) {

	//Pega a pasta raiz do projeto
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.ParseFiles(wd + filename)
	if err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}

	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}

}
