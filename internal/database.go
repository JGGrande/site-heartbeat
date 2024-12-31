package internal

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func IniciarBancoDeDados(caminho string) error {
	var err error
	db, err = sql.Open("sqlite3", caminho)

	if err != nil {
		return fmt.Errorf("erro ao abrir o banco de dados: %w", err)
	}

	criarTabelasSql := `
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

	_, err = db.Exec(criarTabelasSql)
	if err != nil {
		return fmt.Errorf("erro ao criar tabelas: %w", err)
	}

	return nil
}

func CriarSiteNoBanco(nome string, url string) (string, error) {
	uuid := uuid.New().String()

	sql := `INSERT INTO sites (uuid, nome, url) VALUES (?, ?, ?)`

	_, err := db.Exec(sql, uuid, nome, url)

	if err != nil {
		return "", fmt.Errorf("erro ao criar site: %w", err)
	}

	return uuid, nil
}

func ListarSitesDoBanco() ([]Site, error) {
	sql := `SELECT uuid, nome, url FROM sites`

	rows, err := db.Query(sql)

	if err != nil {
		return nil, fmt.Errorf("erro ao listar sites: %w", err)
	}

	defer rows.Close()

	var sites []Site

	for rows.Next() {
		var site Site

		err := rows.Scan(&site.Uuid, &site.Nome, &site.Url)

		if err != nil {
			return nil, fmt.Errorf("erro ao escanear site: %w", err)
		}

		sites = append(sites, site)
	}

	return sites, nil
}

func CriarLogNoBanco(siteUuid string, texto string) error {
	sql := `INSERT INTO logs (site_uuid, texto) VALUES (?, ?)`

	_, err := db.Exec(sql, siteUuid, texto)

	if err != nil {
		return fmt.Errorf("erro ao criar log: %w", err)
	}

	return nil
}

func ListarLogsDeUmSiteNoBanco(siteUuid string) ([]Log, error) {
	sql := `SELECT site_uuid, texto FROM logs WHERE site_uuid = ?`

	rows, err := db.Query(sql, siteUuid)

	if err != nil {
		return nil, fmt.Errorf("erro ao listar logs: %w", err)
	}

	defer rows.Close()

	var logs []Log

	for rows.Next() {
		var log Log

		err := rows.Scan(&log.SiteUuid, &log.Texto)

		if err != nil {
			return nil, fmt.Errorf("erro ao escanear log: %w", err)
		}

		logs = append(logs, log)
	}

	return logs, nil
}
