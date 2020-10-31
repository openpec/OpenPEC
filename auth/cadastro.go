package auth

import (
	"html/template"
	"log"
	"net/http"
	"regexp"

	"github.com/openpec/OpenPEC/config"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4"
)

//CadastroGet é o handler para a tela de cadastro
func CadastroGet(srv *config.Server) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		config.Render(w, "/templates/auth/cadastro.gohtml", nil)

	}
}

//CadastroPost é o handler de submissão do cadastro
func CadastroPost(srv *config.Server) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Println("Erro ao pegar os dados do formulário: ", err)
		}

		//Faz o tratamento dos valores que podem ser null. Utiliza o pacote null.v4
		ende := template.HTMLEscapeString(r.Form.Get("Endereco"))
		endereco := null.NewString(
			ende, ende != "")

		cns := template.HTMLEscapeString(r.Form.Get("CNS"))
		CNS := null.NewString(
			cns, cns != "")

		cid := template.HTMLEscapeString(r.Form.Get("Cidade"))
		cidade := null.NewString(
			cid, cid != "")

		n := template.HTMLEscapeString(r.Form.Get("Num"))
		num := null.NewString(
			n, n != "")

		ba := template.HTMLEscapeString(r.Form.Get("Bairro"))
		bairro := null.NewString(
			ba, ba != "")

		c := template.HTMLEscapeString(r.Form.Get("CEP"))
		cep := null.NewString(
			c, c != "")

		tel := template.HTMLEscapeString(r.Form.Get("Tel"))
		telefone := null.NewString(
			tel, tel != "")

		//informações do usuário
		user := &config.User{
			CPF:        template.HTMLEscapeString(r.Form.Get("CPF")),
			Email:      template.HTMLEscapeString(r.Form.Get("Email")),
			Nome:       template.HTMLEscapeString(r.Form.Get("Nome")),
			Sobrenome:  template.HTMLEscapeString(r.Form.Get("Sobrenome")),
			CNS:        CNS,
			Sexo:       r.Form.Get("Sexo"),
			Cidade:     cidade,
			Estado:     r.Form.Get("Estado"),
			Endereco:   endereco,
			Num:        num,
			Bairro:     bairro,
			CEP:        cep,
			Tel:        telefone,
			Nascimento: r.Form.Get("Nascimento"),
		}

		errosDados := make(map[string]string)

		//Chama a função para validação dos dados
		valida, err := validate(user, errosDados, srv)
		if err != nil {
			log.Panic(err)
			http.Error(w, "Desculpe, algo deu errado ", http.StatusInternalServerError)
			return
		}

		//Valida a Senha
		s := len(r.Form.Get("Senha")) >= 6 && r.Form.Get("Senha") == r.Form.Get("Senha2")

		if !s {
			errosDados["Senha"] = "As senhas não coincidem ou possuem menos de 6 caracteres."
		}

		//Verifica se os dados foram validados ou não
		if !valida || !s {
			config.Render(w, "/templates/auth/cadastro.gohtml", errosDados)
		} else {
			//Executa o cadastro

			//Encripta a senha
			hashPass, err := bcrypt.GenerateFromPassword([]byte(r.Form.Get("Senha")), bcrypt.DefaultCost) // generate bcrypt
			if err != nil {
				log.Panic("Erro ao encriptar a senha: ", err)
			}

			//Insere os dados no Banco de Dados
			stmt, err := srv.DB.Prepare("INSERT user SET cpf=?, pass=?, email=?, nome=?, sobrenome=?, cns=?, sexo=?, cidade=?, estado=?, endereco=?, numCasa=?, bairro=?, cep=?, telefone=?, nascimento=?, isAdmin=?")
			if err != nil {
				log.Panic("Erro ao preparar o banco de dados para inserção: ", err)
			}

			// executa o comando sql
			res, err := stmt.Exec(user.CPF, hashPass, user.Email, user.Nome, user.Sobrenome, user.CNS, user.Sexo, user.Cidade, user.Estado, user.Endereco, user.Num, user.Bairro, user.CEP, user.Tel, user.Nascimento, false)
			if err != nil {
				log.Panic("Erro ao inserir os dados no banco de dados: ", err)
			}

			uid, err := res.LastInsertId()
			if err != nil {
				log.Panic("Erro ao pegar o id da última linha inserida: ", err)
			}

			if uid == 1 { //O primeiro usuário criado é o administrador do sistema
				stmt, err := srv.DB.Prepare("UPDATE user SET isAdmin=? WHERE cod=?")
				if err != nil {
					log.Panic("Erro ao preparar o banco de dados para inserção: ", err)
				}

				// executa o comando sql
				_, err = stmt.Exec(true, uid)
				if err != nil {
					log.Panic("Erro ao atualizar os dados no banco de dados: ", err)
				}
			}
			log.Println("Cadastro efetuado com sucesso")

			//Sucesso no cadastro
			http.Redirect(w, r, "/", http.StatusFound) //Redireciona para a tela de login
		}

	}
}

func validate(user *config.User, eD map[string]string, srv *config.Server) (bool, error) {

	//sexo, estado e nascimento não precisam de validação pois as escolhas no formulário são exatas

	//check CPF
	r, err := regexp.Compile(`^[\d]+$`)
	if err != nil {
		return false, errors.Wrap(err, "não foi possível criar a expressão regular")
	}
	match := r.MatchString(user.CPF)

	if !match || len(user.CPF) != 11 {
		eD["CPF"] = "Insira um CPF válido (Ex. 98765432199)."
	} else {
		//verifica se já existe no BD
		stmt, err := srv.DB.Prepare("SELECT * FROM user WHERE cpf=?")
		if err != nil {
			return false, errors.Wrap(err, "erro ao preparar o banco de dados:")
		}

		rows, err := stmt.Query(user.CPF)
		if err != nil {
			return false, errors.Wrap(err, "erro ao consultar o banco de dados:")
		}
		defer rows.Close()

		if rows.Next() { //verifica se existe
			eD["CPF"] = "Esse CPF já está cadastrado."
		}
		if rows.Err() != nil {
			return false, errors.Wrap(err, "falha na busca do CPF no banco de dados:")
		}

	}

	//check email
	r, err = regexp.Compile("^[a-zA-Z\\d\\.\\-_@]+@[a-zA-Z]+\\..+") //aceita letras, números e os símbolos - _ @ .
	if err != nil {
		return false, errors.Wrap(err, "não foi possível criar a expressão regular")
	}
	match = r.MatchString(user.Email)

	if !match {
		eD["Email"] = "Insira um e-mail válido (Ex. openpec@email.com)."
	} else {
		//verifica se já existe no BD
		stmt, err := srv.DB.Prepare("SELECT * FROM user WHERE email=?")
		if err != nil {
			return false, errors.Wrap(err, "erro ao preparar o banco de dados:")
		}

		rows, err := stmt.Query(user.Email)
		if err != nil {
			return false, errors.Wrap(err, "erro ao consultar o banco de dados:")
		}
		defer rows.Close()

		if rows.Next() { //verifica se existe
			eD["Email"] = "Esse e-mail já está cadastrado."
		}
		if rows.Err() != nil {
			return false, errors.Wrap(err, "falha na busca do email no banco de dados:")
		}

	}

	//check Nome
	r, err = regexp.Compile("^[A-Za-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ]+$")
	if err != nil {
		return false, errors.Wrap(err, "não foi possível criar a expressão regular")
	}
	match = r.MatchString(user.Nome)

	if !match {
		eD["Nome"] = "Insira um nome válido. Apenas o primeiro nome."
	}

	//check Sobrenome
	r, err = regexp.Compile("^[A-Za-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ ]+$")
	if err != nil {
		return false, errors.Wrap(err, "não foi possível criar a expressão regular")
	}
	match = r.MatchString(user.Sobrenome)

	if !match {
		eD["Sobrenome"] = "Insira um sobrenome válido (Ex. da Silva Souza)."
	}

	//check CNS
	r, err = regexp.Compile(`^[\d]+$`)
	if err != nil {
		return false, errors.Wrap(err, "não foi possível criar a expressão regular")
	}
	aux := user.CNS.ValueOrZero() //retorna o valor ou vazio
	match = r.MatchString(aux)

	if (!match || len(aux) != 15) && !user.CNS.IsZero() {
		eD["CNS"] = "Insira um CNS válido (São 15 números)."
	} else if !user.CNS.IsZero() {
		//verifica se já existe no BD
		stmt, err := srv.DB.Prepare("SELECT * FROM user WHERE cns=?")
		if err != nil {
			return false, errors.Wrap(err, "erro ao preparar o banco de dados:")
		}

		rows, err := stmt.Query(aux)
		if err != nil {
			return false, errors.Wrap(err, "erro ao consultar o banco de dados:")
		}
		defer rows.Close()

		if rows.Next() { //verifica se existe
			eD["CNS"] = "Esse CNS já está cadastrado."
		}
		if rows.Err() != nil {
			return false, errors.Wrap(err, "falha na busca do CNS no banco de dados:")
		}

	}

	//check Cidade
	r, err = regexp.Compile("^[A-Za-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ ]+$")
	if err != nil {
		return false, errors.Wrap(err, "não foi possível criar a expressão regular")
	}
	aux = user.Cidade.ValueOrZero() //retorna o valor ou vazio
	match = r.MatchString(aux)

	if !match && len(aux) > 0 {
		eD["Cidade"] = "Insira um nome de cidade válido."
	}

	//check Bairro
	r, err = regexp.Compile("^[A-Za-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ 0123456789]+$")
	if err != nil {
		return false, errors.Wrap(err, "não foi possível criar a expressão regular")
	}
	aux = user.Bairro.ValueOrZero() //retorna o valor ou vazio
	match = r.MatchString(aux)

	if !match && len(aux) > 0 {
		eD["Bairro"] = "Insira um nome de bairro válido."
	}

	//check Endereco
	r, err = regexp.Compile("^[A-Za-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑ 0123456789]+$")
	if err != nil {
		return false, errors.Wrap(err, "não foi possível criar a expressão regular")
	}
	aux = user.Endereco.ValueOrZero() //retorna o valor ou vazio
	match = r.MatchString(aux)

	if !match && len(aux) > 0 {
		eD["Endereco"] = "Insira um endereco válido."
	}

	//check Num
	r, err = regexp.Compile(`^[\d]+$`)
	if err != nil {
		return false, errors.Wrap(err, "não foi possível criar a expressão regular")
	}
	aux = user.Num.ValueOrZero() //retorna o valor ou vazio
	match = r.MatchString(aux)

	if !match && len(aux) > 0 {
		eD["Num"] = "Insira um número válido."
	}

	//check CEP
	r, err = regexp.Compile(`^[\d]+$`)
	if err != nil {
		return false, errors.Wrap(err, "não foi possível criar a expressão regular")
	}
	aux = user.CEP.ValueOrZero() //retorna o valor ou vazio
	match = r.MatchString(aux)

	if (!match || len(aux) != 8) && len(aux) > 0 {
		eD["CEP"] = "Insira um CEP válido (São 8 números)."
	}

	//check Tel
	r, err = regexp.Compile(`^[\d]+$`)
	if err != nil {
		return false, errors.Wrap(err, "não foi possível criar a expressão regular")
	}
	aux = user.Tel.ValueOrZero() //retorna o valor ou vazio
	match = r.MatchString(aux)

	if (!match || (len(aux) != 10 && len(aux) != 11)) && len(aux) > 0 {
		eD["Tel"] = "Insira um número de telefone válido (Digite apenas números)."
	}

	return len(eD) == 0, nil

}
