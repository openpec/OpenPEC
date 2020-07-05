package routes

import (
	"github.com/OpenPEC/auth"
	"github.com/OpenPEC/config"
	"github.com/OpenPEC/modules"
	"github.com/OpenPEC/modules/administracao"
	"github.com/OpenPEC/modules/agenda"
	"github.com/OpenPEC/modules/atendimentos"
	"github.com/OpenPEC/modules/cidadao"
	"github.com/OpenPEC/modules/configuracoes"
	"github.com/OpenPEC/modules/relatorios"
)

//Routes define todas as rotas do webservice
func Routes(srv *config.Server) {

	srv.Router.Handle("/", auth.LoginGet(srv)).Methods("GET")
	srv.Router.Handle("/", auth.LoginPost(srv)).Methods("POST")
	srv.Router.Handle("/logout", auth.Logout(srv)).Methods("GET")
	srv.Router.Handle("/naologado", auth.NaoLogado(srv)).Methods("GET")
	srv.Router.Handle("/home", modules.HomeGet(srv)).Methods("GET")

	//modulo atendimentos
	srv.Router.Handle("/atendimentos", atendimentos.HomeGet(srv)).Methods("GET")

	//modulo agenda
	srv.Router.Handle("/agenda", agenda.HomeGet(srv)).Methods("GET")

	//modulo administracao
	srv.Router.Handle("/administracao", administracao.HomeGet(srv)).Methods("GET")

	//modulo cidadao
	srv.Router.Handle("/cidadao", cidadao.HomeGet(srv)).Methods("GET")

	//modulo configuracoes
	srv.Router.Handle("/configuracoes", configuracoes.HomeGet(srv)).Methods("GET")

	//modulo relatorios
	srv.Router.Handle("/relatorios", relatorios.HomeGet(srv)).Methods("GET")

}
