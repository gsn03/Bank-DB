package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/lib/pq"
)

const connStr = "user=usuario password=senha dbname=db host=db sslmode=disable"

func main() {
    db, err := sql.Open("postgres", connStr)
    if err != nil { log.Fatal(err) }
    defer db.Close()

    fmt.Println("\n=== FUNCIONARIO ===")
    mat, _ := CriarFuncionario(db, "Func Teste", "81990001111", "Gerente", "2024-07-24", "func.teste", "senhateste")
    fmt.Println("Criado:", mat)
    rows, _ := BuscarFuncionarios(db)
    for rows.Next() {
        var m, n, t, c, d, l string
        rows.Scan(&m, &n, &t, &c, &d, &l)
        fmt.Println("Funcionario:", m, n, t, c, d, l)
    }
    AtualizarFuncionarioTelefone(db, mat, "81990001122")
    fmt.Println("Telefone atualizado")

    fmt.Println("\n=== GERENTE ===")
    CriarGerente(db, mat, true)
    fmt.Println("Gerente criado")
    AtualizarGerenteLimite(db, mat, false)
    fmt.Println("Limite de aprovação atualizado")

    fmt.Println("\n=== AGENCIA ===")
    numAg, _ := CriarAgencia(db, "End Teste Ag", "81333112233", mat)
    fmt.Println("Agencia criada:", numAg)
    rows, _ = BuscarAgencias(db)
    for rows.Next() {
        var num int
        var end, tel string
        var gerente int
        rows.Scan(&num, &end, &tel, &gerente)
        fmt.Println("Agencia:", num, end, tel, gerente)
    }
    AtualizarAgenciaTelefone(db, numAg, "8133334444")
    fmt.Println("Telefone agência atualizado")

    fmt.Println("\n=== ADMINISTRATIVO ===")
    matAdm, _ := CriarFuncionario(db, "Func Adm", "81990001133", "Administrativo", "2024-07-24", "adm.teste", "senhaadm")
    CriarAdministrativo(db, matAdm, "RH", 2, mat, numAg)
    fmt.Println("Administrativo criado")
    AtualizarAdministrativoSetor(db, matAdm, "Financeiro")
    fmt.Println("Setor administrativo atualizado")

    fmt.Println("\n=== CLIENTE ===")
    idCli, _ := CriarCliente(db, "Rua Teste, 123", "81994442222")
    fmt.Println("Cliente criado:", idCli)
    rows, _ = BuscarClientes(db)
    for rows.Next() {
        var id int
        var end, tel string
        rows.Scan(&id, &end, &tel)
        fmt.Println("Cliente:", id, end, tel)
    }
    AtualizarClienteTelefone(db, idCli, "81993334444")
    fmt.Println("Telefone cliente atualizado")

    fmt.Println("\n=== PESSOA_FISICA ===")
    CriarPessoaFisica(db, idCli, "999.111.222-55", "Teste PF", "2000-01-01")
    fmt.Println("Pessoa Física criada")
    AtualizarPessoaFisicaNome(db, idCli, "Nome PF Atualizado")
    fmt.Println("Nome PF atualizado")

    fmt.Println("\n=== PESSOA_JURIDICA ===")
    idCli2, _ := CriarCliente(db, "Av PJ, 987", "81995556666")
    CriarPessoaJuridica(db, idCli2, "11.222.333/0001-55", "Empresa Teste")
    fmt.Println("Pessoa Jurídica criada")
    AtualizarPessoaJuridicaRazao(db, idCli2, "Empresa Teste 2")
    fmt.Println("Razão Social PJ atualizada")

    fmt.Println("\n=== DEPENDENTE ===")
    idDep, _ := CriarDependente(db, "Dep Teste", "Rua Dep, 1", "81990001234", idCli)
    fmt.Println("Dependente criado:", idDep)
    AtualizarDependenteTelefone(db, idDep, "81991111234")
    fmt.Println("Telefone dependente atualizado")

    fmt.Println("\n=== CONTA ===")
    numConta, _ := CriarConta(db, "corrente", 2000, 500, "ativa", "2024-07-24", "senhaConta", mat, idCli, numAg)
    fmt.Println("Conta criada:", numConta)
    AtualizarContaSaldo(db, numConta, 999.99)
    fmt.Println("Saldo conta atualizado")

    err = CriarContaCorrente(db, numConta)
    if err != nil { fmt.Println("Erro criando Conta Corrente:", err) } else { fmt.Println("Conta Corrente vinculada") }

    numContaP, _ := CriarConta(db, "poupanca", 500, 300, "ativa", "2024-07-24", "senhaCP", mat, idCli, numAg)
    err = CriarContaPoupanca(db, numContaP)
    if err != nil { fmt.Println("Erro criando Conta Poupança:", err) } else { fmt.Println("Conta Poupança criada") }

    fmt.Println("\n=== CARTAO ===")
    CriarCartao(db, "8888999911112222", "123", 2000, "2030-01-01", "ativo", "Visa", "credito", numConta)
    fmt.Println("Cartão criado")
    AtualizarCartaoStatus(db, "8888999911112222", "bloqueado")
    fmt.Println("Status cartão atualizado")

    fmt.Println("\n=== TRANSACAO ===")
    idTrans, _ := CriarTransacao(db, "2024-07-25", "14:00:00", 1000.0, "deposito", nil, &numConta)
    fmt.Println("Transação depósito criada:", idTrans)
    idSaque, _ := CriarTransacao(db, "2024-07-25", "15:00:00", 200.0, "saque", &numConta, nil)
    fmt.Println("Transação saque criada:", idSaque)
    idTransf, _ := CriarTransacao(db, "2024-07-25", "16:00:00", 300.0, "transferencia", &numConta, &numContaP)
    fmt.Println("Transação transferência criada:", idTransf)

    // Listar transações
    rows, _ = BuscarTransacoes(db)
    for rows.Next() {
        var id int
        var data, hora, tipo string
        var valor float64
        rows.Scan(&id, &data, &hora, &valor, &tipo)
        fmt.Println("Transação:", id, data, hora, valor, tipo)
    }

    // ===== DEPENDENTE_CONTA =====
    fmt.Println("\n=== DEPENDENTE_CONTA ===")
    DarAcessoDependenteConta(db, idDep, numConta)
    fmt.Println("Dependente teve acesso à conta vinculado")
    RemoverAcessoDependenteConta(db, idDep, numConta)
    fmt.Println("Acesso dependente/conta removido")

    // ===== LISTAGENS FINAIS DE TUDO =====
    fmt.Println("\nFuncionários:")
    rows, _ = BuscarFuncionarios(db)
    for rows.Next() {
        var m, n, t, c, d, l string
        rows.Scan(&m, &n, &t, &c, &d, &l)
        fmt.Println(m, n, t, c, d, l)
    }
    fmt.Println("Clientes:")
    rows, _ = BuscarClientes(db)
    for rows.Next() {
        var id int
        var end, tel string
        rows.Scan(&id, &end, &tel)
        fmt.Println(id, end, tel)
    }
    fmt.Println("Contas:")
    rows, _ = BuscarContas(db)
    for rows.Next() {
        var num int
        var tipo string
        var limite, saldo float64
        var status string
        rows.Scan(&num, &tipo, &limite, &saldo, &status)
        fmt.Println(num, tipo, limite, saldo, status)
    }
    fmt.Println("Cartões:")
    rows, _ = BuscarCartoes(db)
    for rows.Next() {
        var num, bandeira, tipo string
        var limite float64
        var status string
        rows.Scan(&num, &bandeira, &tipo, &limite, &status)
        fmt.Println(num, bandeira, tipo, limite, status)
    }
    fmt.Println("Dependentes:")
    rows, _ = BuscarDependentes(db)
    for rows.Next() {
        var id, idCli int
        var nome string
        rows.Scan(&id, &nome, &idCli)
        fmt.Println(id, nome, idCli)
    }
    
    // ===== EXCLUSÕES FINAIS =====
    fmt.Println("\n--- LIMPEZA FINAL ---")

    DeletarTransacao(db, idTrans)
    DeletarTransacao(db, idSaque)
    DeletarTransacao(db, idTransf)
    DeletarCartao(db, "8888999911112222")

    DeletarContaCorrente(db, numConta)
    DeletarContaPoupanca(db, numContaP)

    DeletarConta(db, numConta)
    DeletarConta(db, numContaP)

    DeletarDependente(db, idDep)

    DeletarPessoaFisica(db, idCli)
    DeletarPessoaJuridica(db, idCli2)

    DeletarCliente(db, idCli)
    DeletarCliente(db, idCli2)

    DeletarAdministrativo(db, matAdm)
    DeletarAgencia(db, numAg)

    DeletarGerente(db, mat)

    DeletarFuncionario(db, matAdm)
    DeletarFuncionario(db, mat)

    fmt.Println("--- Teste CRUD COMPLETO Finalizado ---")
}
