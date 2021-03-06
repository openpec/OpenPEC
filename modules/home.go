package modules

import (
	"log"
	"net/http"

	"github.com/openpec/OpenPEC/config"
)

//HomeGet é o handler da pagina principal
func HomeGet(srv *config.Server) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		session, err := srv.Store.Get(r, "userInfo")
		if err != nil {
			log.Println("Conflito com os cookies. É preciso deletar o cookie de sessão passada para gerar um novo cookie.")
			http.Error(w, err.Error()+".\n\nVocê deve deletar o cookie do site localhost:9090 manualmente. \nLembre-se de clicar em 'Sair' quando desconectar do OpenPEC.", http.StatusInternalServerError)
			return
			//chrome://settings/cookies/detail?site=localhost
			//link no chrome pra gerenciar cookie do localhost
		}

		user := config.GetUser(session)

		//Verifica se está logado
		if auth := user.Authenticated; !auth {
			http.Redirect(w, r, "/naologado", http.StatusFound)
		}

		config.Render(w, "/templates/home.gohtml", user.Nome)
	}
}
