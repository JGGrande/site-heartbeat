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
		ativo BOOLEAN NOT NULL,
		data TEXT NOT NULL,
		FOREIGN KEY(site_uuid) REFERENCES sites(uuid)
	);`

	_, err = db.Exec(criarTabelasSql)
	if err != nil {
		return fmt.Errorf("erro ao criar tabelas: %w", err)
	}

	return nil
}

func CriarSiteNoBanco(nome string, url string) (Site, error) {
	uuid := uuid.New().String()

	sql := `INSERT INTO sites (uuid, nome, url) VALUES (?, ?, ?)`

	_, err := db.Exec(sql, uuid, nome, url)

	if err != nil {
		return Site{}, fmt.Errorf("erro ao criar site: %w", err)
	}

	novoSite := Site{
		Uuid: uuid,
		Nome: nome,
		Url:  url,
	}

	return novoSite, nil
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

func ExcluirSiteDoBanco(siteUuid string) error {
	sql := `DELETE FROM logs WHERE site_uuid = ?`

	_, err := db.Exec(sql, siteUuid)

	if err != nil {
		return fmt.Errorf("erro ao excluir logs: %w", err)
	}

	sql = `DELETE FROM sites WHERE uuid = ?`

	_, err = db.Exec(sql, siteUuid)

	if err != nil {
		return fmt.Errorf("erro ao excluir site: %w", err)
	}

	return nil
}

func CriarLogNoBanco(siteUuid string, texto string, ativo bool, data string) error {
	sql := `INSERT INTO logs (site_uuid, texto, ativo, data) VALUES (?, ?, ?, ?)`

	_, err := db.Exec(sql, siteUuid, texto, ativo, data)

	if err != nil {
		return fmt.Errorf("erro ao criar log: %w", err)
	}

	return nil
}

func ListarLogsDeUmSiteNoBanco(siteUuid string) ([]Log, error) {
	sql := `
		SELECT site_uuid, texto, ativo, data 
		FROM logs
		WHERE site_uuid = ?
		ORDER By texto ASC
	`

	rows, err := db.Query(sql, siteUuid)

	if err != nil {
		return nil, fmt.Errorf("erro ao listar logs: %w", err)
	}

	defer rows.Close()

	var logs []Log

	for rows.Next() {
		var log Log

		err := rows.Scan(&log.SiteUuid, &log.Texto, &log.Ativo, &log.Data)

		if err != nil {
			return nil, fmt.Errorf("erro ao escanear log: %w", err)
		}

		logs = append(logs, log)
	}

	return logs, nil
}
