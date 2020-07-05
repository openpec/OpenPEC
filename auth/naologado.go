package auth

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/OpenPEC/config"
)

//NaoLogado é a página que o usuário é redirecionado quando nao está logado
func NaoLogado(srv *config.Server) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		//Pega a pasta raiz do projeto
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		//template html
		t, err := template.ParseFiles(wd + "/templates/naologado.gohtml")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, nil)

	}
}
