package auth

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/OpenPEC/config"
	"gopkg.in/guregu/null.v4"
)

//Testa a função de validação fornecendo diversos casos de testes dos dados dos usuários

func TestValidateRegex(t *testing.T) {

	//inicia um mock do banco de dados
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	srv := &config.Server{
		DB: db,
	}

	//Cria uma linha vazia no BD
	rowsVazia := sqlmock.NewRows([]string{"cpf", "email", "cns"})

	//Cria uma struct com todos os casos de teste
	tables := []struct {
		usuario *config.User
		saida   bool
	}{
		//padrão, tudo certo com apenas os campos obrigatórios
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
		}, true},
		//CPF tests
		{&config.User{
			CPF:        "999888777666", //12 caracteres
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
		}, false},
		{&config.User{
			CPF:        "9998a877766", //letra
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
		}, false},
		{&config.User{
			CPF:        "99988@77766", //símbolo
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
		}, false},
		//email tests
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpecemail.com", //sem @
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
		}, false},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@emailcom", //sem .
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
		}, false},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@.com", //sem email
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
		}, false},
		{&config.User{
			CPF:        "99988877766",
			Email:      "@email.com", //sem inicio
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
		}, false},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.", //sem fim
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
		}, false},
		{&config.User{
			CPF:        "99988877766",
			Email:      "open pec@email.com", //com espaço
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
		}, false},
		{&config.User{
			CPF:        "99988877766",
			Email:      "open!pec@email.com", //com símbolo inválido
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
		}, false},
		//Nome tests
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC da Silva", //mais de uma palavra
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
		}, false},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC1", //com número
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
		}, false},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "Open@PEC", //com símbolo
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
		}, false},
		//Sobrenome tests
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema de 10", //com número
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
		}, false},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema!", //com símbolo
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
		}, false},
		//CNS tests
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			CNS:        null.NewString("", "" != ""), //string vazia
		}, true},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			CNS:        null.NewString("999888777666555", "999888777666555" != ""), //string correta
		}, true},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			CNS:        null.NewString("999888", "999888" != ""), //quantidade de números incorreta
		}, false},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			CNS:        null.NewString("999@88777666555", "999@88777666555" != ""), //com símbolo
		}, false},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			CNS:        null.NewString("999888777a66555", "999888777a66555" != ""), //com letra
		}, false},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			CNS:        null.NewString("999888 77666555", "999888 77666555" != ""), //com espaço
		}, false},
		//cidade tests
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			Cidade:     null.NewString("São Paulo", "São Paulo" != ""), //string correta
		}, true},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			Cidade:     null.NewString("São 0Paulo", "São Paulo" != ""), //com número
		}, false},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			Cidade:     null.NewString("São Pau%lo", "São Paulo" != ""), //com símbolo
		}, false},
		//bairro tests
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			Bairro:     null.NewString("Novo Horizonte", "Novo Horizonte" != ""), //string correta
		}, true},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			Bairro:     null.NewString("Novo# Horizonte", "Novo Horizonte" != ""), //com símbolo
		}, false},
		//endereco tests
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			Endereco:   null.NewString("Rua 9 de Julho", "Rua 9 de Julho" != ""), //string correta
		}, true},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			Endereco:   null.NewString("Rua * de Julho", "Rua * de Julho" != ""), //com símbolo
		}, false},
		//Num tests
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			Num:        null.NewString("9988", "9988" != ""), //string correta
		}, true},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			Num:        null.NewString("9a88", "9a88" != ""), //com letra
		}, false},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			Num:        null.NewString("9&88", "9&88" != ""), //com símbolo
		}, false},
		//CEP tests
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			CEP:        null.NewString("12222555", "12222555" != ""), //string correta
		}, true},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			CEP:        null.NewString("122a2555", "122a2555" != ""), //com letra
		}, false},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			CEP:        null.NewString("12222-555", "12222-555" != ""), //com símbolo
		}, false},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			CEP:        null.NewString("122555", "122555" != ""), //com quantidade errada de caracteres
		}, false},
		//Tel tests
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			Tel:        null.NewString("11988887777", "11988887777" != ""), //string correta
		}, true},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			Tel:        null.NewString("119888877", "119888877" != ""), //com quantidade incorreta de caracteres
		}, false},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			Tel:        null.NewString("119888877ab", "119888877ab" != ""), //com letras
		}, false},
		{&config.User{
			CPF:        "99988877766",
			Email:      "openpec@email.com",
			Nome:       "OpenPEC",
			Sobrenome:  "Sistema",
			Sexo:       "Masculino",
			Estado:     "SP",
			Nascimento: "2000-01-01",
			Tel:        null.NewString("119888!7777", "119888!7777" != ""), //com símbolos
		}, false},
	}

	for _, table := range tables {
		ed := make(map[string]string)
		x, y, z := false, false, false

		//Evitar erro com o mock do banco de dados
		if _, err := validate(table.usuario, ed, srv); err != nil {
			if ed["CPF"] == "" {
				mock.ExpectPrepare("^SELECT (.+) FROM user WHERE cpf*").ExpectQuery().
					WithArgs(table.usuario.CPF).WillReturnRows(rowsVazia)
				_, err = validate(table.usuario, ed, srv)
				x = true
			}
			if ed["Email"] == "" {
				if x {
					mock.ExpectPrepare("^SELECT (.+) FROM user WHERE cpf*").ExpectQuery().
						WithArgs(table.usuario.CPF).WillReturnRows(rowsVazia)
				}
				mock.ExpectPrepare("^SELECT (.+) FROM user WHERE email*").ExpectQuery().
					WithArgs(table.usuario.Email).WillReturnRows(rowsVazia)
				_, err = validate(table.usuario, ed, srv)
				y = true
			}
			if ed["CNS"] == "" {
				if len(table.usuario.CNS.ValueOrZero()) == 15 { //verifica se string válida ou campo vazio
					z = true
				}
			}
		}

		if x {
			mock.ExpectPrepare("^SELECT (.+) FROM user WHERE cpf*").ExpectQuery().
				WithArgs(table.usuario.CPF).WillReturnRows(rowsVazia)
		}
		if y {
			mock.ExpectPrepare("^SELECT (.+) FROM user WHERE email*").ExpectQuery().
				WithArgs(table.usuario.Email).WillReturnRows(rowsVazia)
		}
		if z {
			mock.ExpectPrepare("^SELECT (.+) FROM user WHERE cns*").ExpectQuery().
				WithArgs(table.usuario.CNS).WillReturnRows(rowsVazia)
		}

		//Parte principal - testa as entradas e saídas
		val, err := validate(table.usuario, ed, srv)
		if val != table.saida {
			t.Log("Erro não esperado: ", err, "ou validação incorreta: ", ed)
			t.Errorf("Fail_test - Entrada %v não esperava a saída %t\n\n", table.usuario, table.saida)
		} else {
			t.Logf("Tudo conforme esperado com a entrada %v e a saída %t\n\n", table.usuario, table.saida)
		}

		//Verifica se o mock foi usado corretamente
		err = mock.ExpectationsWereMet()
		if err != nil {
			log.Println(err)
		}
	}

}
