package auth

import (
	"log"
	"net/http"

	"github.com/openpec/OpenPEC/config"
)

//Logout faz a desconexão do usuário
func Logout(srv *config.Server) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		session, err := srv.Store.Get(r, "userInfo")
		if err != nil {
			log.Println("Conflito com os cookies. É preciso deletar o cookie de sessão passada para gerar um novo cookie.")
			http.Error(w, err.Error()+".\n\nVocê deve deletar o cookie do site localhost:9090 manualmente. \nLembre-se de clicar em 'Sair' quando desconectar do OpenPEC.", http.StatusInternalServerError)
			return
		}

		//Zera a estrutura user e expira o cookie
		session.Values["user"] = config.User{}
		session.Options.MaxAge = -1

		err = session.Save(r, w)
		if err != nil {
			log.Println("Erro ao salvar a sessão")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("Logout efetuado com sucesso")

		http.Redirect(w, r, "/", http.StatusFound)

	}
}

//NaoLogado é a página que o usuário é redirecionado quando nao está logado
func NaoLogado(srv *config.Server) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		config.Render(w, "/templates/auth/naologado.gohtml", nil)

	}
}
