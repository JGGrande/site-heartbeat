package main

import (
	"log"
	"net/http"
	"sitehb/internal"
)

func main() {
	internal.IniciarBancoDeDados("sitehb.db")

	internal.IniciarMonitoresHandler()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("resources"))))

	http.HandleFunc("/criar-monitoramento", internal.CriarMonitoramento)

	http.HandleFunc("/historico/", internal.RenderHistorico)

	http.HandleFunc("/parar-monitoramento", internal.PararMonitoramento)

	http.HandleFunc("/", internal.RenderHome)

	log.Println("Servidor iniciado em http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
