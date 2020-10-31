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
		log.Panic("Erro ao pegar a pasta raiz do projeto: ", err)
	}

	if wd == "/" {
		wd = "go/src/github.com/OpenPEC/"
	}

	t, err := template.ParseFiles(wd + filename)
	if err != nil {
		log.Println("Problema ao renderizar a página: ", err)
		http.Error(w, "Desculpe, algo deu errado ", http.StatusInternalServerError)
	}

	if err := t.Execute(w, data); err != nil {
		log.Println("Problema ao renderizar a página: ", err)
		http.Error(w, "Desculpe, algo deu errado ", http.StatusInternalServerError)
	}

}
