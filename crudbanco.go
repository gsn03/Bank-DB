package main

import (
    "database/sql"
    _ "github.com/lib/pq"
)

const connStr = "user=seuusuario password=suasenha dbname=seudb sslmode=disable"

// ===== FUNCIONARIO CRUD =====
func CriarFuncionario(db *sql.DB, nome, telefone, cargo, dataAdmissao, login, senha string) (int, error) {
    var matricula int
    err := db.QueryRow(`INSERT INTO Funcionario (Nome, Telefone, Cargo, Data_admissao, Login, Senha) VALUES ($1, $2, $3, $4, $5, $6) RETURNING Matricula`, nome, telefone, cargo, dataAdmissao, login, senha).Scan(&matricula)
    return matricula, err
}

func BuscarFuncionarios(db *sql.DB) (*sql.Rows, error) {
    return db.Query(`SELECT Matricula, Nome, Telefone, Cargo, Data_admissao, Login FROM Funcionario`)
}

func AtualizarFuncionarioTelefone(db *sql.DB, matricula int, novoTelefone string) error {
    _, err := db.Exec(`UPDATE Funcionario SET Telefone = $1 WHERE Matricula = $2`, novoTelefone, matricula)
    return err
}

func DeletarFuncionario(db *sql.DB, matricula int) error {
    _, err := db.Exec(`DELETE FROM Funcionario WHERE Matricula = $1`, matricula)
    return err
}

// ===== GERENTE CRUD =====
func CriarGerente(db *sql.DB, matricula int, aprovarLimite bool) error {
    _, err := db.Exec(`INSERT INTO Gerente (Matricula, Aprovar_limite) VALUES ($1, $2)`, matricula, aprovarLimite)
    return err
}

func AtualizarGerenteLimite(db *sql.DB, matricula int, novo bool) error {
    _, err := db.Exec(`UPDATE Gerente SET Aprovar_limite = $1 WHERE Matricula = $2`, novo, matricula)
    return err
}

func DeletarGerente(db *sql.DB, matricula int) error {
    _, err := db.Exec(`DELETE FROM Gerente WHERE Matricula = $1`, matricula)
    return err
}

// ===== ADMINISTRATIVO CRUD =====
func CriarAdministrativo(db *sql.DB, matricula int, setor string, nivelAcesso int, gerenteResp int, numAgencia int) error {
    _, err := db.Exec(`INSERT INTO Administrativo VALUES ($1, $2, $3, $4, $5)`, matricula, setor, nivelAcesso, gerenteResp, numAgencia)
    return err
}

func AtualizarAdministrativoSetor(db *sql.DB, matricula int, novoSetor string) error {
    _, err := db.Exec(`UPDATE Administrativo SET Setor = $1 WHERE Matricula = $2`, novoSetor, matricula)
    return err
}

func DeletarAdministrativo(db *sql.DB, matricula int) error {
    _, err := db.Exec(`DELETE FROM Administrativo WHERE Matricula = $1`, matricula)
    return err
}

// ===== CLIENTE CRUD =====
func CriarCliente(db *sql.DB, endereco, telefone string) (int, error) {
    var id int
    err := db.QueryRow(`INSERT INTO Cliente (Endereco, Telefone) VALUES ($1, $2) RETURNING Id_cliente`, endereco, telefone).Scan(&id)
    return id, err
}

func BuscarClientes(db *sql.DB) (*sql.Rows, error) {
    return db.Query(`SELECT Id_cliente, Endereco, Telefone FROM Cliente`)
}

func AtualizarClienteTelefone(db *sql.DB, id int, telefone string) error {
    _, err := db.Exec(`UPDATE Cliente SET Telefone = $1 WHERE Id_cliente = $2`, telefone, id)
    return err
}

func DeletarCliente(db *sql.DB, id int) error {
    _, err := db.Exec(`DELETE FROM Cliente WHERE Id_cliente = $1`, id)
    return err
}

// ===== PESSOA FISICA/JURIDICA CRUD =====
func CriarPessoaFisica(db *sql.DB, idCliente int, cpf, nome, dataNasc string) error {
    _, err := db.Exec(`INSERT INTO Pessoa_Fisica VALUES ($1, $2, $3, $4)`, idCliente, cpf, nome, dataNasc)
    return err
}

func CriarPessoaJuridica(db *sql.DB, idCliente int, cnpj, razao string) error {
    _, err := db.Exec(`INSERT INTO Pessoa_Juridica VALUES ($1, $2, $3)`, idCliente, cnpj, razao)
    return err
}

func BuscarPessoasFisicas(db *sql.DB) (*sql.Rows, error) {
    return db.Query(`SELECT Id_cliente, CPF, Nome_completo, Data_nasc FROM Pessoa_Fisica`)
}

func BuscarPessoasJuridicas(db *sql.DB) (*sql.Rows, error) {
    return db.Query(`SELECT Id_cliente, CNPJ, Razao_social FROM Pessoa_Juridica`)
}

func AtualizarPessoaFisicaNome(db *sql.DB, idCliente int, nome string) error {
    _, err := db.Exec(`UPDATE Pessoa_Fisica SET Nome_completo = $1 WHERE Id_cliente = $2`, nome, idCliente)
    return err
}

func AtualizarPessoaJuridicaRazao(db *sql.DB, idCliente int, razao string) error {
    _, err := db.Exec(`UPDATE Pessoa_Juridica SET Razao_social = $1 WHERE Id_cliente = $2`, razao, idCliente)
    return err
}

func DeletarPessoaFisica(db *sql.DB, idCliente int) error {
    _, err := db.Exec(`DELETE FROM Pessoa_Fisica WHERE Id_cliente = $1`, idCliente)
    return err
}

func DeletarPessoaJuridica(db *sql.DB, idCliente int) error {
    _, err := db.Exec(`DELETE FROM Pessoa_Juridica WHERE Id_cliente = $1`, idCliente)
    return err
}

// ===== DEPENDENTE CRUD =====
func CriarDependente(db *sql.DB, nome, endereco, telefone string, idCliente int) (int, error) {
    var id int
    err := db.QueryRow(`INSERT INTO Dependente (Nome_dependente, Endereco, Telefone, Id_cliente) VALUES ($1, $2, $3, $4) RETURNING Id_dependente`, nome, endereco, telefone, idCliente).Scan(&id)
    return id, err
}

func BuscarDependentes(db *sql.DB) (*sql.Rows, error) {
    return db.Query(`SELECT Id_dependente, Nome_dependente, Id_cliente FROM Dependente`)
}

func AtualizarDependenteTelefone(db *sql.DB, id int, telefone string) error {
    _, err := db.Exec(`UPDATE Dependente SET Telefone = $1 WHERE Id_dependente = $2`, telefone, id)
    return err
}

func DeletarDependente(db *sql.DB, id int) error {
    _, err := db.Exec(`DELETE FROM Dependente WHERE Id_dependente = $1`, id)
    return err
}

// ===== AGENCIA CRUD =====
func CriarAgencia(db *sql.DB, endereco, telefone string, gerenteAgencia int) (int, error) {
    var num int
    err := db.QueryRow(`INSERT INTO Agencia (Endereco, Telefone, Gerente_agencia) VALUES ($1, $2, $3) RETURNING Num_agencia`, endereco, telefone, gerenteAgencia).Scan(&num)
    return num, err
}

func BuscarAgencias(db *sql.DB) (*sql.Rows, error) {
    return db.Query(`SELECT Num_agencia, Endereco, Telefone, Gerente_agencia FROM Agencia`)
}

func AtualizarAgenciaTelefone(db *sql.DB, num int, telefone string) error {
    _, err := db.Exec(`UPDATE Agencia SET Telefone = $1 WHERE Num_agencia = $2`, telefone, num)
    return err
}

func DeletarAgencia(db *sql.DB, num int) error {
    _, err := db.Exec(`DELETE FROM Agencia WHERE Num_agencia = $1`, num)
    return err
}

// ===== CONTA CRUD =====
func CriarConta(db *sql.DB, tipo string, limite, saldo float64, status, dataAbertura, senha string, abertaPor, idCliente, numAgencia int) (int, error) {
    var num int
    err := db.QueryRow(`INSERT INTO Conta (Tipo, Limite, Saldo, Status, Data_abertura, Senha, Aberta_por, Id_cliente, Num_agencia) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING Num_conta`, tipo, limite, saldo, status, dataAbertura, senha, abertaPor, idCliente, numAgencia).Scan(&num)
    return num, err
}

func BuscarContas(db *sql.DB) (*sql.Rows, error) {
    return db.Query(`SELECT Num_conta, Tipo, Limite, Saldo, Status FROM Conta`)
}

func AtualizarContaSaldo(db *sql.DB, num int, saldo float64) error {
    _, err := db.Exec(`UPDATE Conta SET Saldo = $1 WHERE Num_conta = $2`, saldo, num)
    return err
}

func DeletarConta(db *sql.DB, num int) error {
    _, err := db.Exec(`DELETE FROM Conta WHERE Num_conta = $1`, num)
    return err
}

// ===== CARTAO CRUD =====
func CriarCartao(db *sql.DB, numCartao, codVerif string, limite float64, validade, status, bandeira, tipo string, numConta int) error {
    _, err := db.Exec(`INSERT INTO Cartao VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, numCartao, codVerif, limite, validade, status, bandeira, tipo, numConta)
    return err
}

func BuscarCartoes(db *sql.DB) (*sql.Rows, error) {
    return db.Query(`SELECT Num_cartao, Bandeira, Tipo, Limite, Status FROM Cartao`)
}

func AtualizarCartaoStatus(db *sql.DB, numCartao, status string) error {
    _, err := db.Exec(`UPDATE Cartao SET Status = $1 WHERE Num_cartao = $2`, status, numCartao)
    return err
}

func DeletarCartao(db *sql.DB, numCartao string) error {
    _, err := db.Exec(`DELETE FROM Cartao WHERE Num_cartao = $1`, numCartao)
    return err
}

// ===== TRANSACAO CRUD =====
func CriarTransacao(db *sql.DB, data, hora string, valor float64, tipo string, numOrigem, numDestino *int) (int, error) {
    var id int
    err := db.QueryRow(`INSERT INTO Transacao (Data, Hora, Valor, Tipo, Num_conta_origem, Num_conta_destino) VALUES ($1, $2, $3, $4, $5, $6) RETURNING Id_transacao`, data, hora, valor, tipo, numOrigem, numDestino).Scan(&id)
    return id, err
}

func BuscarTransacoes(db *sql.DB) (*sql.Rows, error) {
    return db.Query(`SELECT Id_transacao, Data, Hora, Valor, Tipo FROM Transacao`)
}

func DeletarTransacao(db *sql.DB, id int) error {
    _, err := db.Exec(`DELETE FROM Transacao WHERE Id_transacao = $1`, id)
    return err
}

// ===== RELACIONAMENTOS ASSOCIATIVOS =====
func DarAcessoDependenteConta(db *sql.DB, idDependente, numConta int) error {
    _, err := db.Exec(`INSERT INTO Dependente_Conta VALUES ($1, $2)`, idDependente, numConta)
    return err
}

func RemoverAcessoDependenteConta(db *sql.DB, idDependente, numConta int) error {
    _, err := db.Exec(`DELETE FROM Dependente_Conta WHERE Id_dependente = $1 AND Num_conta = $2`, idDependente, numConta)
    return err
}
