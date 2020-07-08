package config

import (
	"encoding/gob"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"gopkg.in/guregu/null.v4"
)

//User define a estrutura para a sessão
type User struct {
	Authenticated bool
	CPF           string
	Nome          string
	Sobrenome     string
	Email         string
	IsAdmin       bool
	CNS           null.String
	Sexo          string
	Cidade        null.String
	Estado        string
	Endereco      null.String
	Num           null.String
	Bairro        null.String
	CEP           null.String
	Tel           null.String
	Nascimento    string
}

//StartSession inicializa as configurações pra sessão
func (srv *Server) StartSession() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	srv.Store = sessions.NewCookieStore(authKeyOne, encryptionKeyOne)

	srv.Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 12, //12 horas
		HttpOnly: true,
	}

	gob.Register(User{})
}

//GetUser retorna o usuário da sessão
func GetUser(s *sessions.Session) User {
	val := s.Values["user"]
	var user = User{}
	user, ok := val.(User)
	if !ok {
		return User{Authenticated: false}
	}
	return user
}
