# Site-Heartbeat

O **Site-Heartbeat** é uma aplicação simples desenvolvida em Go para verificar a disponibilidade de sites. Ele realiza uma "heartbeat" para monitorar se os sites estão respondendo.

## Pré-requisitos

Certifique-se de que você possui os seguintes pré-requisitos instalados:

- [Go](https://golang.org/dl/) (versão 1.19 ou superior)
- Acesso ao terminal ou linha de comando

## Instalação

1. Clone este repositório para sua máquina local:

    ```bash
    git clone https://github.com/JGGrande/site-heartbeat.git
    cd site-heartbeat
    ```
2. Instale as dependências necessárias:

    ```bash
    go mod tidy
    ```

3. Execute o servidor WEB:
    
    ```bash
    go run .
    ```

4. Tudo pronto! agora a aplicação está disponível em [http://localhost:8080](http://localhost:8080)