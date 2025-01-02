package internal

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
)

func RenderHome(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("resources/html/home.html")

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

func CriarMonitoramento(w http.ResponseWriter, r *http.Request) {
	type NovoSiteRequest struct {
		Nome string `json:"nome"`
		URL  string `json:"url"`
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req NovoSiteRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Dados inválidos no corpo da requisição", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if req.Nome == "" || req.URL == "" {
		http.Error(w, "Campos 'nome' e 'url' são obrigatórios", http.StatusBadRequest)
		return
	}

	err = RegistrarNovoSiteHandler(req.Nome, req.URL)

	if err != nil {
		http.Error(w, "Erro ao registrar o site", http.StatusInternalServerError)
		log.Println("Erro ao registrar o site:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func PararMonitoramento(w http.ResponseWriter, r *http.Request) {
	siteUuid := r.URL.Query().Get("site")

	if siteUuid == "" {
		http.Error(w, "Site não encontrado", http.StatusNotFound)
		return
	}

	ExcluirMonitoramentoHandler(siteUuid)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
