package main

import (
	"log"
	"net/http"
	"sitehb/internal"
)

func main() {
	internal.IniciarBancoDeDados("sitehb.db")

	internal.IniciarMonitores()

	http.HandleFunc("/", internal.RenderHome)

	http.HandleFunc("/parar-monitoramento", internal.PararMonitoramento)

	log.Println("Servidor iniciado em http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
