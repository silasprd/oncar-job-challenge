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

# É necessário ter instalado em sua máquina o [Go](https://go.dev/) e [MySQL](https://www.mysql.com/).

### 🎲 Clonando e configurando o banco de dados

```bash
# Em um terminal clone o repositório com o comando
$ git clone <https://github.com/silasprd/oncar-job-challenge>
```
<br>

#### Você precisará criar um banco de dados para o sistema.
```bash
# No MySQL Workbench ou alguma outra ferramenta gerenciadora de banco de dados, você poderá criar um banco de dados para a aplicação com o seguinte comando
$ create database database-name;
# Você pode substituir o database-name pelo nome que você quiser dar ao seu banco de dados
```

<br>

#### Dentro da pasta do projeto você precisará acessar o arquivo .env dentro da pasta api.
```bash
# Você pode acessar por linha de comando no próprio terminal com os seguintes comandos
$ cd api
# Você pode abrir a pasta onde contém o arquivo com o seguinte comando
$ start .
# Após acessar a pasta você precisará abrir o arquivo em um editor de textos sua escolha
```

<br>

#### Sera necessário configurar algumas credenciais de usuário no projeto
```bash
# As variáveis de ambiente presentes no arquivo representam as que você utilizar para conectar seu banco de dados local.
'DB_USER'='nome de usuário do banco'
'DB_PASSWORD'='senha do banco'
'DB_NAME'='nome do banco de dados criado para o sistema'
'DB_HOST'='IP host local'
'DB_PORT'='Porta em que o banco de dados está rodando localmente'
```
#### Após substituir os valores das variáveis para as suas credenciais, você pode salvar o arquivo e fechá-lo

<br>

### Após todas as configurações, ainda na pasta do projeto, siga estes passos para rodar a aplicação localmente
```bash
# Acesse a pasta api
$ cd api
# Rode o projeto com o comando
$ go run main.go
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