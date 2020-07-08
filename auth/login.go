package auth

import (
	"html/template"
	"log"
	"net/http"

	"github.com/OpenPEC/config"
	"golang.org/x/crypto/bcrypt"
)

//LoginGet é o handler do login com método get
func LoginGet(srv *config.Server) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		session, err := srv.Store.Get(r, "userInfo")
		if err != nil {
			http.Error(w, err.Error()+".\n\nVocê deve deletar o cookie do site localhost:9090 manualmente. \nLembre-se de clicar em 'Sair' quando desconectar do OpenPEC.", http.StatusInternalServerError)
			return
		}

		user := config.GetUser(session)

		if user.Authenticated {
			config.Render(w, "/templates/auth/login.gohtml", user)
		} else {
			config.Render(w, "/templates/auth/login.gohtml", nil)
		}

	}
}

//LoginPost é o handler do login com método post
func LoginPost(srv *config.Server) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		session, err := srv.Store.Get(r, "userInfo")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		r.ParseForm()

		//Checa dados
		stmt, err := srv.DB.Prepare("SELECT * FROM user WHERE cpf=?")
		if err != nil {
			log.Fatal(err)
		}

		rows, err := stmt.Query(template.HTMLEscapeString(r.Form.Get("CPF")))
		if err != nil {
			log.Panic(err)

		}
		defer rows.Close()

		var id int
		var pass string

		//Para validação
		errosDados := make(map[string]string)

		user := new(config.User) //Inicia a struct para pegar as informações do BD

		if !rows.Next() {
			if rows.Err() == nil {
				//CPF não existe no BD
				errosDados["CPF"] = "Esse CPF não está cadastrado."
				config.Render(w, "/templates/auth/login.gohtml", errosDados)

			} else {
				log.Panic("Erro em pegar o rows.next: ", rows.Err())
			}
		} else {

			//Recebe os valores do BD
			err = rows.Scan(&id, &user.CPF, &pass, &user.Nome, &user.Sobrenome, &user.Email, &user.CNS, &user.Sexo, &user.Cidade, &user.Estado, &user.Endereco, &user.Num, &user.Bairro, &user.CEP, &user.Tel, &user.Nascimento, &user.IsAdmin)
			if err != nil {
				log.Panic("Erro no rows.scan: ", err)

			}

			//verifica se a senha está correta para esse CPF
			err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(template.HTMLEscapeString(r.Form.Get("Senha")))) // validating password

			if err == nil { //Login aceito

				user.Authenticated = true

				session.Values["user"] = user //a sessão recebe os dados do usuário
				err = session.Save(r, w)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				http.Redirect(w, r, "home", http.StatusFound) //Redireciona para página principal

			} else {
				errosDados["Pass"] = "A senha está incorreta."
				config.Render(w, "/templates/auth/login.gohtml", errosDados)
			}

		}

	}
}
