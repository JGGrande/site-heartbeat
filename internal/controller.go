package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"
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

	siteExiste := VerificarSeSiteExisteHandler(req.URL)

	if !siteExiste {
		http.Error(w, "URL inválida.", http.StatusBadRequest)
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

	w.WriteHeader(http.StatusOK)
}

func RenderHistorico(w http.ResponseWriter, r *http.Request) {
	siteUuid := strings.TrimPrefix(r.URL.Path, "/historico/")

	if siteUuid == "" {
		http.Error(w, "UUID inválido", http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles("resources/html/historico.html")

	if err != nil {
		http.Error(w, "Erro ao carregar o template", http.StatusInternalServerError)
		log.Println("Erro ao carregar o template:", err)
		return
	}

	logs, err := ListarLogsDeUmSiteNoBanco(siteUuid)

	if err != nil {
		http.Error(w, "Erro ao listar os logs", http.StatusNotFound)
		log.Println("Erro ao listar os logs:", err)
		return
	}

	type HistoricoTemplateData struct {
		Logs       []Log
		DataLabels []string
		DataPoints []int
	}

	var labels []string
	var data []int

	for _, log := range logs {
		t, err := time.Parse("02/01/2006 15:04:05", log.Data)

		if err != nil {
			fmt.Println("Erro ao processar data:", err)
			continue
		}

		label := t.Format("2006-01-02 15:04:05")

		labels = append(labels, label)

		if log.Ativo {
			data = append(data, 100)
		} else {
			data = append(data, 0)
		}
	}

	templateData := HistoricoTemplateData{
		Logs:       logs,
		DataLabels: labels,
		DataPoints: data,
	}

	err = tmpl.Execute(w, templateData)

	if err != nil {
		http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
		log.Println("Erro ao renderizar o template:", err)
	}
}
