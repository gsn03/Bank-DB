-- ==========================
-- FUNCIONÁRIOS, GERENTES, ADMINISTRATIVOS
-- ==========================
INSERT INTO Funcionario VALUES (1, 'Ana Gerente', '81990000001', 'Gerente', '2020-01-10', 'ana.gerente', 'senha123');
INSERT INTO Funcionario VALUES (2, 'Bruno Adm', '81990000002', 'Administrativo', '2022-03-12', 'bruno.adm', 'senha456');
INSERT INTO Funcionario VALUES (3, 'Carlos Gerente', '81990000003', 'Gerente', '2018-09-05', 'carlos.gerente', 'senha789');

INSERT INTO Gerente VALUES (1, TRUE);
INSERT INTO Gerente VALUES (3, TRUE);

INSERT INTO Administrativo VALUES (2, 'RH', 1, 1, 1);

-- ==========================
-- CLIENTES (PF e PJ)
-- ==========================
INSERT INTO Cliente (Endereco, Telefone) VALUES ('Rua A, 123', '81990000010'); -- Id_cliente = 1
INSERT INTO Cliente (Endereco, Telefone) VALUES ('Av. B, 456', '81990000020'); -- Id_cliente = 2
INSERT INTO Cliente (Endereco, Telefone) VALUES ('Rua C, 789', '81990000030'); -- Id_cliente = 3

-- PESSOA FÍSICA
INSERT INTO Pessoa_Fisica VALUES (1, '111.222.333-44', 'João Silva', '1995-06-12');
INSERT INTO Pessoa_Fisica VALUES (3, '222.333.444-55', 'Maria Souza', '1988-11-20');

-- PESSOA JURÍDICA
INSERT INTO Pessoa_Juridica VALUES (2, '12.345.678/0001-99', 'Tech Solutions LTDA');

-- ==========================
-- DEPENDENTES
-- ==========================
INSERT INTO Dependente (Nome_dependente, Endereco, Telefone, Id_cliente)
VALUES ('Lucas Filho', 'Rua D, 100', '81990000011', 1); -- Id_dependente = 1
INSERT INTO Dependente (Nome_dependente, Endereco, Telefone, Id_cliente)
VALUES ('Pedro Menor', 'Av. E, 200', '81990000012', 3); -- Id_dependente = 2

-- ==========================
-- AGÊNCIAS
-- ==========================
INSERT INTO Agencia VALUES (1, 'Centro', '8133334444', 1);
INSERT INTO Agencia VALUES (2, 'Boa Viagem', '8144445555', 3);

-- ==========================
-- CONTAS
-- ==========================
INSERT INTO Conta VALUES (1001, 'corrente', 5000, 1200, 'ativa', '2023-01-01', 'senhaConta1', 1, 1, 1); -- PF, João
INSERT INTO Conta VALUES (1002, 'corrente', 20000, 8000, 'ativa', '2023-01-15', 'senhaConta2', 3, 2, 2); -- PJ, Tech Solutions
INSERT INTO Conta VALUES (1003, 'poupanca', 3000, 1500, 'ativa', '2023-01-20', 'senhaConta3', 1, 1, 1); -- PF, João

INSERT INTO Conta_corrente VALUES (1001);
INSERT INTO Conta_corrente VALUES (1002);
INSERT INTO Conta_poupanca VALUES (1003);

-- ==========================
-- DEPENDENTE_CONTA (acesso do dependente à conta)
-- ==========================
INSERT INTO Dependente_Conta VALUES (1, 1001); -- Lucas Filho acesso à conta 1001
INSERT INTO Dependente_Conta VALUES (2, 1003); -- Pedro Menor acesso à conta 1003

-- ==========================
-- CARTÕES
-- ==========================
INSERT INTO Cartao VALUES ('5555666677778888', '123', 3000, '2028-12-01', 'ativo', 'Visa', 'credito', 1001);
INSERT INTO Cartao VALUES ('9999000011112222', '456', 1500, '2027-06-15', 'ativo', 'Mastercard', 'debito', 1002);

-- ==========================
-- TRANSACOES
-- ==========================
INSERT INTO Transacao (Data, Hora, Valor, Tipo, Num_conta_origem, Num_conta_destino)
VALUES ('2024-06-01', '08:00', 1000, 'deposito', NULL, 1001);
INSERT INTO Transacao (Data, Hora, Valor, Tipo, Num_conta_origem, Num_conta_destino)
VALUES ('2024-06-02', '10:30', 500, 'saque', 1001, NULL);
INSERT INTO Transacao (Data, Hora, Valor, Tipo, Num_conta_origem, Num_conta_destino)
VALUES ('2024-06-03', '14:10', 300, 'transferencia', 1001, 1002);
INSERT INTO Transacao (Data, Hora, Valor, Tipo, Num_conta_origem, Num_conta_destino)
VALUES ('2024-06-04', '09:20', 1200, 'deposito', NULL, 1002);
INSERT INTO Transacao (Data, Hora, Valor, Tipo, Num_conta_origem, Num_conta_destino)
VALUES ('2024-06-05', '15:15', 700, 'transferencia', 1002, 1001);
