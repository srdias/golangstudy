# Teste com acesso ao banco de dados Postgres

### Pré-requisitos:
Docker Desktop instalado. Link para a instalação: [https://docs.docker.com/docker-for-windows/install/](https://docs.docker.com/docker-for-windows/install/).

### Instalando o package lib/pq para permitir o acesso ao banco de dados postgres
```
go get -u github.com/lib/pq
```

### Configurando o ambiente antes de executar o teste:

O arquivo docker-compose.yml inicia um ambiente com um servidor de banco de dados postgres e outro servidor com o PGAdmin.
O servidor do PG Admin é iniciado na porta 16543. Link: [http://localhost:16543/] (http://localhost:16543/). Usuário de acesso: **postgres**, senha: **admin**.
Ao iniciar o banco de dados, o comando init.sql é executado, criando a tabela pessoas, tabela que será usada no teste com o GO.

#### Comando para iniciar os dockers:
Precisa do docker instalado.
```
docker-compose up -d
```
OBS: Ao iniciar o docker o windows irá solicitar confirmação porque o docker-compose está utilizando compartilhamento de pastas entre windows e docker.

#### Setup do banco de dados para executar o teste:
1. Abrir o PG Admin conforme endereço/usuário/senha acima;
2. No grupo "Quick Links" acesse a opção "Add New Server";
3. Será aberta uma janela. Informe:
    - Guia 'Geral':
        - Campo 'Name' informe 'Postgres'
    - Guia 'Connection'
        - Campo 'Host' informe 'go-postgres'
        - Campo 'Username' informe 'postgres'
        - Campo 'Password' informe 'admin'
    - Clique no botão 'Save'

#### Executando o exemplo:
```
go run pg-teste.go
```
O exemplo (pg-teste.go) executa testes de select, insert, update e delete. (**Teste do CRUD**)
1. **Teste com o comando SELECT**: Lista os dados da tabela pessoas para confirmar que tabela está sem dados.
2. **Teste com o comando INSERT**: Insere 4 registros e em seguida lista todos os registros.
3. **Teste com o comando UPDATE**: Altera o nome da pessoa de código 2 e apresenta todos os registros para que seja confirmado a alteração.
4. **Teste com o comando DELETE**: Exclui a pessoa de codigo 3 e apresenta todos os registros para confirmar que a pessoa não está mais na tabela.
5. Executa outro comando DELETE para limpar a tabela e deixar pronto para uma nova execução do teste.
