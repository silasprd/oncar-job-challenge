# onCar Job Challenge

### Desenvolvido para o desafio proposto pela empresa onCar

<br>

<p align="left">Um sistema de listagem de veículos com possibilidade de seleção de veículos e o envio de informações de contato, bem como a consulta dos dados do veículo e dos contatos enviados. Estes dados são todos salvos em um banco de dados SQL.</p>

<br>

## 🙅‍♂️ Autor: Silas Rafael Barreto Prado

<br>

## Tecnologias

#### Sistema desenvolvido com as seguintes tecnologias:

- **Backend:** [Go](https://go.dev/)
- **Frontend:** [JavaScript](https://developer.mozilla.org/pt-BR/docs/Web/JavaScript)
- **Banco de dados:** [MySQL](https://www.mysql.com/)

<br>

<details>
    <summary style="font-size: 18px; font-weight: bolder;">Decisões técnicas</summary>
    <ul style="list-style: none;">
        <li>
            <b>Arquitetura MSC:</b><span> Foi utilizado a Arquitetura MSC na criação da api, para deixar o código bem estruturado, facilitando a manutenção e escalabilidade do código.</span>
        </li>
        <li>
            <b>ORM:</b><span> Utilizamos a biblioteca GORM da linguagem Go, para nos ajudar no mapeamento do objeto-relacional e no uso do banco de dados.</span>
        </li>
        <li>
            <b>Mock de dados com sqlmock:</b><span> Nos testes utilizamos a biblioteca sqlmock para mockagem dos dados, assim sendo possível simular um banco de dados real para realização dos teste unitários.</span>
        </li>
        <li>
            <b>Material Design:</b><span> No frontend, a biblioteca do material io foi utilizada para fornecer acesso a alguns ícones utilizados na parte web.</span>
        </li>
    </ul>
</details>

<br>

<details>
    <summary style="font-size: 18px; font-weight: bolder;">Estrutura das pastas</summary>
    <ul style="list-style: none;">
        <legend style="font-weight: bolder">BACKEND:</legend>
        <li>
            <b>/.env:</b><span> Arquivo de definição das variáveis globais.</span>
        </li>
        <li>
            <b>/.gitgnore:</b><span> Arquivo de configuração do rastreamento de controle de versão do git.</span>
        </li>
        <li>
            <b>/main.go:</b><span> Arquivo principal onde são executados os servidores web e api.</span>
        </li>
        <li>
            <b>/api:</b><span> Onde estão toda a estrutura e os arquivos da API desenvolvida na linguagem Go.</span>
        </li>
        <li>
            <b>/api/core/:</b><span> Modelos, serviços e controladores desenvolvidos para atender as requisições.</span>
        </li>
        <li>
            <b>/api/core/model/:</b><span> Modelo dos dados utilizados na API.</span>
        </li>
        <li>
            <b>/api/core/service/:</b><span> Toda a lógica do negócio, onde são executadas as querys para manipulação dos dados no banco.</span>
        </li>
        <li>
            <b>/api/core/controller/:</b><span> Onde estão os controladores, responsáveis por fazer as requisições http.</span>
        </li>
        <li>
            <b>/api/db/:</b><span> Aqui são feitas as configurações e conexão com o banco de dados, e também a auto migração das tabelas.</span>
        </li>
        <li>
            <b>/api/routes/:</b><span> Todas as definições de rotas utilizadas na aplicação.</span>
        </li>
        <li>
            <b>/api/test/:</b><span> Esta pasta contém todos os arquivos de teste. A pasta tem a mesma estrutura da pasta 'api'. Esta pasta deve simular a api para realização dos testes.</span>
        </li>
    </ul> 
    <br>
    <ul style="list-style: none;">
        <legend style="font-weight: bolder;">FRONTEND</legend>
        <li>
            <b>/web/:</b><span> Aqui estão todos os arquivos utilizados para criação da página web.</span>
        <li>
        <li>
            <b>/web/css:</b><span> Arquivo de estilização da página.</span>
        </li>
        <li>
            <b>/web/pages:</b><span> Arquivos HTML renderizados para a página.</span>
        </li>
        <li>
            <b>/web/script:</b><span> Aqui é onde está a lógica por trás do frontend, onde são feitas as chamadas para a api.</span>
        </li>
    </ul>  
</details>

<br>

## ℹ️ Como rodar o sistema localmente

#### É necessário ter instalado em sua máquina o [Go](https://go.dev/) e o [MySQL](https://www.mysql.com/), também é desejável um bom editor de textos como por exemplo o [Visual Studio Code](https://code.visualstudio.com/).

### 🎲 Clonando o projeto e configurando o banco de dados

#### Clone o projeto
```bash
# Em um terminal clone o repositório com o comando
$ git clone https://github.com/silasprd/oncar-job-challenge
```

#### Você precisará criar um banco de dados para o sistema.
```bash
# No MySQL Workbench ou alguma outra ferramenta gerenciadora do banco de dados, execute o comando
$ create database database-name;
# Você pode substituir o database-name pelo nome que você quiser dar ao seu banco de dados.
```

#### Dentro da pasta do projeto você precisará acessar o arquivo .env dentro da pasta api.
```bash
# Você pode acessar por linha de comando no próprio terminal com o seguinte comando
$ cd api
# Você pode abrir a pasta onde contém o arquivo com o seguinte comando
$ start .
# Após acessar a pasta você precisará abrir o arquivo em um editor de textos de sua escolha
```

#### Será necessário configurar algumas credenciais do banco de dados no projeto.
```bash
# Essas variáveis de ambiente representam as que você utiliza para conectar ao seu banco de dados local.
DB_USER='nome de usuário do banco'
DB_PASSWORD='senha do banco'
DB_NAME='nome do banco de dados criado para o sistema'
DB_HOST='IP host local(Geralmente: 127.0.0.1)'
DB_PORT='Porta em que o banco de dados está rodando localmente'
```
#### Substitua os valores das variáveis pelas credenciais do seu banco de dados, salve o arquivo e pode fechá-lo.

### 👨‍💻 Após estas configurações, ainda na pasta do projeto, rode a aplicação.

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


