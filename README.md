# onCar Job Challenge

### Desenvolvido para o desafio proposto pela empresa onCar

<br>

<p align="left">Um sistema de listagem de veículos como possibilidade de seleção de veículos e envio de informações de contato. Todos os dados são salvos em um banco de dados SQL.</p>

<br>

## 🙅‍♂️ Autor: Silas Rafael Barreto Prado

<br>

## Tecnologias

#### Sistema desenvolvido com as seguintes tecnologias:

- **Backend:** [Go](https://go.dev/)
- **Frontend:** [JavaScript](https://developer.mozilla.org/pt-BR/docs/Web/JavaScript)
- **Banco de dados:** [MySQL](https://www.mysql.com/)

<br>

## ℹ️ Como rodar o sistema localmente

<p align="left">
    É necessário ter instalado em sua máquina o [Go](https://go.dev/) e [MySQL](https://www.mysql.com/). 
</p>

### 🎲 Clonando e rodando a aplicação

```bash
# Clone este repositório
$ git clone <https://github.com/silasprd/oncar-job-challenge>
# Acesse a pasta do projeto no terminal
$ cd oncar-job-challenge/api
# Dentro da pasta 'api', execute o comando no terminal
$ go run main.go
# O projeto deverá ser executado em alguns segundos.
```
<br>

#### Rodando os testes
```bash
# Ainda dentro da pasta api, acesse a pasta de testes
$ cd test
# Dentro da pasta test execute o comando
$ go test -v ./...
# Este comando irá executar todos os arquivos de teste
```