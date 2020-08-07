# Teste com acesso ao banco de dados Postgres

### Instalando o package lib/pq para permitir o acesso ao banco de dados postgres
```
go get -u github.com/lib/pq
```

### Configurando o ambiente antes de executar o teste:

Na pasta raiz deste projeto, tem o docker-compose.yml que inicia um ambiente com um servidor de banco de dados postgres e outro servidor com o PGAdmin.
O servidor do PG Admin é iniciado na porta 16543. Link: [http://localhost:16543/] (http://localhost:16543/). Usuário de acesso: **postgres**, senha: **admin**.

#### Comando para iniciar os dockers:
```
docker-compose up -d
```

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
4. Após esses passos, abrirá uma lista no estilo treeview. Abra os itens até o 4 nivel, em cima do nome do banco de dados 'postgres'. Use o botão do mouse para acessar a opção 'Query Tool...'
5. Na caixa de comando, execute os comandos abaixo:
```
CREATE SCHEMA golang;
ALTER SCHEMA golang OWNER TO POSTGRES;
create table golang.pessoas(
	i_pessoas integer,
	nome character(64),
	primary key(i_pessoas)
);
commit;
```

#### Executando o exemplo:
```
go run primeiro.go
```
O exemplo (primeiro.go) executa testes de select, insert, update e delete. (**Teste do CRUD**)
1. **Teste com o comando SELECT**: Lista os dados da tabela pessoas para confirmar que tabela está sem dados.
2. **Teste com o comando INSERT**: Insere 4 registros e em seguida lista todos os registros.
3. **Teste com o comando UPDATE**: Altera o nome da pessoa de código 2 e apresenta todos os registros para que seja confirmado a alteração.
4. **Teste com o comando DELETE**: Exclui a pessoa de codigo 3 e apresenta todos os registros para confirmar que a pessoa não está mais na tabela.
5. Executa outro comando DELETE para limpar a tabela e deixar pronto para uma nova execução do teste.