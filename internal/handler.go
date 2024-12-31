package internal

func RegistrarNovoSite(nome, url string) {
	siteUuid, err := CriarSiteNoBanco(nome, url)

	if err != nil {
		panic(err)
	}

	go MonitorarSite(siteUuid, url, nome)
}

func ConsultarLogDeUmSite(siteUuid string) []string {
	logs, err := ListarLogsDeUmSiteNoBanco(siteUuid)

	if err != nil {
		panic(err)
	}

	var logsFormatados []string

	for _, log := range logs {
		logsFormatados = append(logsFormatados, log.Texto)
	}

	return logsFormatados
}
