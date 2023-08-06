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

### É necessário ter instalado em sua máquina o [Go](https://go.dev/) e [MySQL](https://www.mysql.com/).

### 🎲 Clonando e configurando o banco de dados

#### Clone o projeto
```bash
# Em um terminal clone o repositório com o comando
$ git clone <https://github.com/silasprd/oncar-job-challenge>
```

#### Você precisará criar um banco de dados para o sistema.
```bash
# No MySQL Workbench ou alguma outra ferramenta gerenciadora de banco de dados, execute o comando
$ create database database-name;
# Você pode substituir o database-name pelo nome que você quiser dar ao seu banco de dados
```

#### Dentro da pasta do projeto você precisará acessar o arquivo .env dentro da pasta api.
```bash
# Você pode acessar por linha de comando no próprio terminal com os seguintes comandos
$ cd api
# Você pode abrir a pasta onde contém o arquivo com o seguinte comando
$ start .
# Após acessar a pasta você precisará abrir o arquivo em um editor de textos sua escolha
```

#### Sera necessário configurar algumas credenciais de usuário no projeto
```bash
# Essas variáveis de ambiente representam as que você utiliza para conectar ao seu banco de dados local.
DB_USER='nome de usuário do banco'
DB_PASSWORD='senha do banco'
DB_NAME='nome do banco de dados criado para o sistema'
DB_HOST='IP host local'
DB_PORT='Porta em que o banco de dados está rodando localmente'
```
#### Substitua os valores das variáveis pelas suas credenciais, salve o arquivo e pode fechá-lo

### 👨‍💻 Após todas as configurações, ainda na pasta do projeto, rode a aplicação

#### Rodando a aplicação
```bash
# Acesse a pasta api
$ cd api
# Rode o projeto com o comando
$ go run main.go
# A aplicação estará acessível em http://localhost:3000
# Certifique-se de não ter nada rodando localmente na porta 3000.
```

#### Rodando os testes
```bash
# Ainda dentro da pasta api, acesse a pasta de testes
$ cd test
# Dentro da pasta test execute o comando
$ go test -v ./...
# Este comando irá executar todos os arquivos de teste
```

<details>
    <summary>Estrutura das pastas e arquivos na raiz</summary>
    <b style="font-style: bold">/.env:</b><span> Arquivo de definição das variáveis globais.</span><br>
    <span style="font-style: bold">/.gitgnore:</span><span> Arquivo de configuração do rastreamento de controle de versão do git.</span><br>
    <span style="font-style: bold">/main.go:</span><span> Arquivo principal onde são executados os servidores web e api.</span><br>
    <span style="font-style: bold">/api:</span><span> Onde estão toda a estrutura e os arquivos da API desenvolvida na linguagem Go.</span><br>
    <span style="font-style: bold">/api/core/:</span><span> Modelos, serviços e controladores desenvolvidos para atender as requisições.</span><br>
    <span style="font-style: bold">/api/core/model/:</span><span> Modelo dos dados utilizados na API.</span><br>
    <span style="font-style: bold">/api/core/service/:</span><span> Toda a lógica do negócio, onde são executadas as querys para manipulação dos dados no banco.</span><br>
    <span style="font-style: bold">/api/core/controller/:</span><span> Onde estão os controladores, responsáveis por fazer as requisições http.</span><br>
    <span style="font-style: bold">/api/db/:</span><span> Aqui são feitas as configurações e conexão com o banco de dados, e também a auto migração das tabelas.</span><br>
    <span style="font-style: bold">/api/routes/:</span><span> Todas as definições de rotas utilizadas na aplicação.</span><br>
    <span style="font-style: bold">/api/test/:</span><span> Esta pasta contém todos os arquivos de teste. A pasta tem a mesma estrutura da pasta 'api'. Esta pasta deve simular a pasta api.</span><br>
    
</details>