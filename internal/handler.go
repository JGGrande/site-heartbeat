package internal

func RegistrarNovoSite(nome, url string) {
	siteUuid := CriarSiteNoBanco(nome, url)

	go MonitorarSite(siteUuid, url, nome)
}

func ConsultarLogDeUmSite(siteUuid string) []string {
	logs := ListarLogsDeUmSiteNoBanco(siteUuid)

	var logsFormatados []string

	for _, log := range logs {
		logsFormatados = append(logsFormatados, log.Texto)
	}

	return logsFormatados
}
