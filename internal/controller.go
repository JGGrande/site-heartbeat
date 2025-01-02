package internal

import (
	"log"
	"net/http"
	"text/template"
)

func RenderHome(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")

	if err != nil {
		http.Error(w, "Erro ao carregar o template", http.StatusInternalServerError)
		log.Println("Erro ao carregar o template:", err)
		return
	}

	sites, err := ListarSitesDoBanco()

	if err != nil {
		http.Error(w, "Erro ao listar os sites", http.StatusNotFound)
		log.Println("Erro ao listar os sites:", err)
		return
	}

	type HomeTemplateData struct {
		Sites []Site
	}

	data := HomeTemplateData{
		Sites: sites,
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
		log.Println("Erro ao renderizar o template:", err)
	}
}

func PararMonitoramento(w http.ResponseWriter, r *http.Request) {
	siteUuid := r.URL.Query().Get("site")

	if siteUuid == "" {
		http.Error(w, "Site não encontrado", http.StatusNotFound)
		return
	}

	ExcluirMonitoramento(siteUuid)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
