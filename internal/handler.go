package internal

import (
	"fmt"
	"sync"
)

var (
	monitoramentoAtivo = make(map[string]chan bool)
	mutex              sync.Mutex
)

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

func IniciarMonitores() {
	sites, err := ListarSitesDoBanco()

	if err != nil {
		panic(err)
	}

	for _, site := range sites {
		func(site Site) {
			mutex.Lock()
			defer mutex.Unlock()

			if _, ativo := monitoramentoAtivo[site.Uuid]; ativo {
				fmt.Printf("Monitoramento já ativo para o site %s\n", site.Nome)
				return
			}

			stopChan := make(chan bool)
			monitoramentoAtivo[site.Uuid] = stopChan

			go func() {
				fmt.Printf("Monitorando o site %s [uuid] %s\n", site.Nome, site.Uuid)

				select {
				case <-stopChan:
					fmt.Printf("Monitoramento do site %s encerrado.\n", site.Nome)
					return
				default:
					MonitorarSite(site.Uuid, site.Url, site.Nome)
				}

			}()

		}(site)
	}
}

func ExcluirMonitoramento(siteUuid string) {
	mutex.Lock()
	defer mutex.Unlock()

	if stopChan, ativo := monitoramentoAtivo[siteUuid]; ativo {
		close(stopChan)
		delete(monitoramentoAtivo, siteUuid)

		err := ExcluirSiteDoBanco(siteUuid)

		if err != nil {
			panic(err)
		}

		fmt.Printf("Monitoramento do site [UUID] %s excluído.\n", siteUuid)
	} else {
		fmt.Printf("Nenhum monitoramento ativo encontrado para o site ID %s.\n", siteUuid)
	}
}
