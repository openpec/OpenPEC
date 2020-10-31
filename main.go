package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/openpec/OpenPEC/config"
	"github.com/openpec/OpenPEC/routes"
)

func main() {

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

}

func run() error {
	db, err := config.Connect()
	if err != nil {
		return err
	}

	srv := &config.Server{
		DB:     db,
		Router: mux.NewRouter(),
	}

	srv.StartSession() //sessions

	routes.Routes(srv)

	//Para usar CSS
	srv.Router.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))

	log.Println("O servidor está rodando na porta 9090...")
	err = http.ListenAndServe(":9090", srv.Router) // setting listening port

	if err != nil {
		log.Panic("ListenAndServe: ", err)
	}
	return nil
}
