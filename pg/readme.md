# Teste com acesso ao banco de dados

### Instalando o package lib/pq para permitir o acesso ao banco de dados postgres
```
go get -u github.com/lib/pq
```

### Configurando o ambiente antes de executar o teste:

Na pasta raiz deste projeto, tem o docker-compose.yml que inicia um ambiente com um servidor de banco de dados postgres e outro servidor com o PGAdmin.
O servidor do PG Admin é iniciado na porta 16543. Link: [http://localhost:16543/] (http://localhost:16543/). Usuário de acesso: **postgres**, senha: **admin**.

Comando para iniciar os dockers:
```
docker-compose up -d
```

Executando o exemplo:
```
go run primeiro.go
```
