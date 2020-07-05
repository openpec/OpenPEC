package auth

import (
	"net/http"

	"github.com/OpenPEC/config"
)

//Logout faz a desconexão do usuário
func Logout(srv *config.Server) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		session, err := srv.Store.Get(r, "userInfo")
		if err != nil {
			http.Error(w, err.Error()+".\n\nVocê deve deletar o cookie do site localhost:9090 manualmente. \nLembre-se de clicar em 'Sair' quando desconectar do OpenPEC.", http.StatusInternalServerError)
			return
		}

		//Zera a estrutura user e expira o cookie
		session.Values["user"] = config.User{}
		session.Options.MaxAge = -1

		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)

	}
}
