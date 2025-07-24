# Bank-DB

Repositório para o projeto de Banco de Dados relacional, desenvolvido para fins acadêmicos usando **PostgreSQL** e **Go**, totalmente automatizado com Docker Compose.

---

## Descrição Geral

Este projeto simula um sistema bancário completo, abrangendo:

- Estruturação de tabelas com chaves primárias, estrangeiras e relacionamentos complexos
- População automática de dados iniciais para todos os cenários (clientes, funcionários, contas, cartões, transações etc.)
- Operações CRUD para todas as entidades principais
- Demonstração dos CRUDs via aplicação Go que executa e comprova a persistência dos dados

O objetivo é demonstrar, na prática, domínio de modelagem, normalização, integração e automação de um sistema de banco de dados relacional realista.

---

## Tecnologias Utilizadas

- **PostgreSQL** — Banco de dados relacional
- **Go** (Golang) — Aplicação para testes CRUD e integração
- **Docker** e **Docker Compose** — Orquestração, automação e reprodutibilidade do ambiente

---

## Estrutura do Projeto

```
Bank-DB/
│   README.md
│   docker-compose.yml
│   Dockerfile
│   main.go
│   crudbanco.go
│
└───sql/
    │   tabelas.sql
    │   popular_exemplo.sql (ignorado no deploy)
    │   consultas_exemplo.sql (ignorado no deploy)
```

- `sql/tabelas.sql` — Criação de todas as tabelas e constraints do modelo relacional
- `sql/popular.sql` — População inicial do banco com dados realistas de exemplo (não é executado automaticamente)
- `crudbanco.go` — Funções de CRUD (Create, Read, Update, Delete) para todas as entidades
- `main.go` — Execução dos testes CRUD (prova de funcionamento prático)
- `docker-compose.yml` e `Dockerfile` — Automatizam toda a infraestrutura
- `sql/consultas_exemplo.sql` — Consultas extras para estudo (não é executado automaticamente)

---

## Como Executar (Ambiente Totalmente Automatizado)

1. **Clone o repositório:**

   ```bash
   git clone <url-do-repositorio>
   cd Bank-DB
   ```

2. **Suba toda a stack automaticamente:**

   ```bash
   docker-compose up --build
   ```

   Isso irá:

   - Subir o container do PostgreSQL já com os scripts de criação e população de tabelas aplicados
   - Instalar todas as dependências Go e compilar o app
   - Executar automaticamente o código Go que testa todos os CRUDs e mostra os resultados no terminal

3. **Resultado esperado:**

   - O terminal do container do app Go mostrará a execução das operações CRUD, servindo como prova prática de funcionamento.
   - O banco pode ser inspecionado usando qualquer client SQL (DBeaver, PgAdmin, etc) apontando para `localhost:5432`, usuário/senha definidos no `docker-compose.yml`.

---

## Exemplo de uso (após subir)

Veja prints ou logs similares a:

```
=== FUNCIONARIO ===
Criado: 10
Funcionario: 10 Func Teste ...
Telefone atualizado
...
=== CLIENTE ===
Cliente criado: 5
Cliente: 5 Rua Teste, 123 81994442222
...
--- Teste CRUD COMPLETO Finalizado ---
```

---

## Observações Importantes

- O projeto já está pronto para rodar em qualquer ambiente com Docker.
- Para modificar as credenciais do banco, altere o `docker-compose.yml` e ajuste o `connStr` no código Go se necessário.

---