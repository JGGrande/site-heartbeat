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

	http.HandleFunc("/", internal.RenderHome)

	http.HandleFunc("/criar-monitoramento", internal.CriarMonitoramento)

	http.HandleFunc("/parar-monitoramento", internal.PararMonitoramento)

	log.Println("Servidor iniciado em http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
