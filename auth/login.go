package auth

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/OpenPEC/config"
	"golang.org/x/crypto/bcrypt"
)

//LoginGet é o handler do login com método get
func LoginGet(srv *config.Server) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		//Pega a pasta raiz do projeto
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		//template html
		t, err := template.ParseFiles(wd + "/templates/login.gohtml")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, nil)
	}
}

//LoginPost é o handler do login com método post
func LoginPost(srv *config.Server) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		//Checa dados
		stmt, err := srv.DB.Prepare("SELECT * FROM user WHERE CPF=?")
		if err != nil {
			log.Fatal(err)
		}

		row, err := stmt.Query(template.HTMLEscapeString(r.Form.Get("enter_cpf")))
		if err != nil {
			fmt.Println(err)
		}

		var id int
		var cpf string
		var pass string

		row.Next()
		err = row.Scan(&id, &cpf, &pass)
		if err != nil {
			fmt.Println("Dados errados ou faltando: ", err)
		}

		err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(template.HTMLEscapeString(r.Form.Get("enter_pass")))) // validating password

		if err == nil {
			http.Redirect(w, r, "home", 302)
		} else {
			fmt.Println("CPF ou senha invalidos: ", err)
		}

	}
}
