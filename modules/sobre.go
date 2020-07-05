package modules

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/OpenPEC/config"
)

//Sobre define a página de sobre
func Sobre(srv *config.Server) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		session, err := srv.Store.Get(r, "userInfo")
		if err != nil {
			http.Error(w, err.Error()+".\n\nVocê deve deletar o cookie do site localhost:9090 manualmente. \nLembre-se de clicar em 'Sair' quando desconectar do OpenPEC.", http.StatusInternalServerError)
			return
		}

		user := config.GetUser(session)

		//Verifica se está logado
		if auth := user.Authenticated; !auth {
			http.Redirect(w, r, "/naologado", http.StatusFound)
		}

		//Pega a pasta raiz do projeto
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		//template html
		t, err := template.ParseFiles(wd + "/templates/sobre.gohtml")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, user.FirstName)
	}
}
