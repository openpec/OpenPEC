package routes

import (
	"github.com/openpec/OpenPEC/auth"
	"github.com/openpec/OpenPEC/config"
	"github.com/openpec/OpenPEC/modules"
	"github.com/openpec/OpenPEC/modules/administracao"
	"github.com/openpec/OpenPEC/modules/agenda"
	"github.com/openpec/OpenPEC/modules/atendimentos"
	"github.com/openpec/OpenPEC/modules/cidadao"
	"github.com/openpec/OpenPEC/modules/configuracoes"
	"github.com/openpec/OpenPEC/modules/relatorios"
)

//Routes define todas as rotas do webservice
func Routes(srv *config.Server) {

	//Autenticação
	srv.Router.Handle("/", auth.LoginGet(srv)).Methods("GET")
	srv.Router.Handle("/", auth.LoginPost(srv)).Methods("POST")
	srv.Router.Handle("/logout", auth.Logout(srv)).Methods("GET")
	srv.Router.Handle("/naologado", auth.NaoLogado(srv)).Methods("GET")
	srv.Router.Handle("/cadastro", auth.CadastroGet(srv)).Methods("GET")
	srv.Router.Handle("/cadastro", auth.CadastroPost(srv)).Methods("POST")

	srv.Router.Handle("/home", modules.HomeGet(srv)).Methods("GET")
	srv.Router.Handle("/sobre", modules.Sobre(srv)).Methods("GET")

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
