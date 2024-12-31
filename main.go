package main

import (
	"fmt"
	"os"
	"sitehb/internal"
)

func pedirDadosSite() (string, string) {
	var nome, url string

	fmt.Print("Digite o nome do site: ")
	fmt.Scan(&nome)

	fmt.Print("Digite a URL do site: ")
	fmt.Scan(&url)

	return nome, url
}

func pedirSiteParaLog() internal.Site {
	sites, err := internal.ListarSitesDoBanco()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println("Sites disponíveis:")

	for i, site := range sites {
		fmt.Printf("%d - %s\n", i, site.Nome)
	}

	var siteEscolhido int
	fmt.Print("Digite o número do site: ")
	fmt.Scan(&siteEscolhido)

	return sites[siteEscolhido]
}

func lerComando() int {
	var comando int

	fmt.Print("Digite o comando: ")

	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi:", comando)

	return comando
}

func exibeMenu() {
	fmt.Println("---------------------------------------")
	fmt.Println("| 1 - Registrar novo site             |")
	fmt.Println("| 2 - Consultar logs                  |")
	fmt.Println("| 0 - Sair do Programa                |")
	fmt.Println("---------------------------------------")
}

func main() {
	fmt.Println("Inicializando Banco de Dados...")
	err := internal.IniciarBancoDeDados("sitehb.db")

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println("Banco de Dados inicializado com sucesso.")

	for {
		exibeMenu()

		comando := lerComando()

		switch comando {
		case 1:
			nome, url := pedirDadosSite()

			internal.RegistrarNovoSite(nome, url)

			fmt.Println("Site registrado com sucesso.")
		case 2:
			site := pedirSiteParaLog()

			sites := internal.ConsultarLogDeUmSite(site.Uuid)

			fmt.Println(sites)
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido.")
			os.Exit(-1)
		}

		fmt.Println("***************************************")
	}
}
