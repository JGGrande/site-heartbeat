package internal

import (
	"fmt"
	"net/http"
	"time"
)

const delay = 5 * time.Second

func MonitorarSite(siteUuid string, url string, nome string) {
	for {
		res, err := http.Get(url)

		temErro := err != nil

		if temErro {
			fmt.Println(err)
			time.Sleep(delay)
			continue
		}

		estaOnline := res.StatusCode == 200

		status := "offline"

		if estaOnline {
			status = "online"
		}

		dataAtualFormatada := time.Now().Format("02/01/2006 15:04:05")

		textoLog := fmt.Sprintf("[%s] - está %s", dataAtualFormatada, status)

		CriarLogNoBanco(siteUuid, textoLog)

		time.Sleep(delay)
	}
}
