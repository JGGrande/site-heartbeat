package internal

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

var sitesDatabase []Site
var logsDatabase []Log

var db *sql.DB

func IniciarBancoDeDados(caminho string) error {
	var err error
	db, err = sql.Open("sqlite", caminho)
	if err != nil {
		return fmt.Errorf("erro ao abrir o banco de dados: %w", err)
	}

	// Criação das tabelas
	criarTabelas := `
	CREATE TABLE IF NOT EXISTS sites (
		uuid TEXT PRIMARY KEY,
		nome TEXT NOT NULL,
		url TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS logs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		site_uuid TEXT NOT NULL,
		texto TEXT NOT NULL,
		FOREIGN KEY(site_uuid) REFERENCES sites(uuid)
	);`

	_, err = db.Exec(criarTabelas)
	if err != nil {
		return fmt.Errorf("erro ao criar tabelas: %w", err)
	}

	return nil
}

func CriarSiteNoBanco(nome string, url string) string {
	uuidCriado := uuid.New().String()

	site := Site{
		Uuid: uuidCriado,
		Nome: nome,
		Url:  url,
	}

	sitesDatabase = append(sitesDatabase, site)

	return uuidCriado
}

func ListarSitesDoBanco() []Site {
	return sitesDatabase
}

func CriarLogNoBanco(siteUuid string, texto string) {
	log := Log{
		SiteUuid: siteUuid,
		Texto:    texto,
	}

	logsDatabase = append(logsDatabase, log)
}

func ListarLogsDeUmSiteNoBanco(siteUuid string) []Log {
	var logs []Log

	for _, log := range logsDatabase {
		if log.SiteUuid == siteUuid {
			logs = append(logs, log)
		}
	}

	return logs
}
