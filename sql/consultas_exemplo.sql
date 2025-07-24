-- CONSULTAS SQL – Sistema Bancário

-- a) Listar todas as contas correntes de clientes Pessoa Física
SELECT c.Num_conta, pf.Nome_completo, pf.CPF, c.Saldo
FROM Conta c
JOIN Conta_corrente cc ON c.Num_conta = cc.Num_conta
JOIN Cliente cli ON c.Id_cliente = cli.Id_cliente
JOIN Pessoa_Fisica pf ON cli.Id_cliente = pf.Id_cliente;

-- b) Listar todas as contas vinculadas a uma agência específica
SELECT c.Num_conta, c.Tipo, c.Saldo, a.Endereco AS agencia_endereco
FROM Conta c
JOIN Agencia a ON c.Num_agencia = a.Num_agencia
WHERE a.Num_agencia = 1;

-- c) Obter todos os dependentes que têm acesso a uma conta específica
SELECT d.Nome_dependente, d.Telefone
FROM Dependente_Conta dc
JOIN Dependente d ON dc.Id_dependente = d.Id_dependente
WHERE dc.Num_conta = 1001;

-- d) Extrato (todas as transações de uma conta, ordenadas)
SELECT t.Data, t.Hora, t.Tipo, t.Valor,
       t.Num_conta_origem, t.Num_conta_destino
FROM Transacao t
WHERE t.Num_conta_origem = 1001 OR t.Num_conta_destino = 1001
ORDER BY t.Data, t.Hora;

-- e) Contas do tipo poupança ligadas (erroneamente) a pessoa jurídica
SELECT c.Num_conta, c.Tipo, pj.Razao_social
FROM Conta c
JOIN Conta_poupanca cp ON c.Num_conta = cp.Num_conta
JOIN Cliente cli ON c.Id_cliente = cli.Id_cliente
LEFT JOIN Pessoa_Juridica pj ON cli.Id_cliente = pj.Id_cliente
WHERE pj.CNPJ IS NOT NULL;

-- f) Total de transações feitas por um cliente (pelo Id_cliente)
SELECT cli.Id_cliente, COALESCE(pf.Nome_completo, pj.Razao_social) AS nome, COUNT(t.Id_transacao) AS total_transacoes
FROM Cliente cli
LEFT JOIN Pessoa_Fisica pf ON cli.Id_cliente = pf.Id_cliente
LEFT JOIN Pessoa_Juridica pj ON cli.Id_cliente = pj.Id_cliente
JOIN Conta c ON cli.Id_cliente = c.Id_cliente
LEFT JOIN Transacao t ON c.Num_conta = t.Num_conta_origem OR c.Num_conta = t.Num_conta_destino
GROUP BY cli.Id_cliente, pf.Nome_completo, pj.Razao_social;

-- g) Listar os cartões ativos de todos os clientes com suas contas
SELECT cartao.Num_cartao, cartao.Bandeira, cartao.Tipo, c.Num_conta, pf.Nome_completo, pj.Razao_social
FROM Cartao cartao
JOIN Conta c ON cartao.Num_conta = c.Num_conta
JOIN Cliente cli ON c.Id_cliente = cli.Id_cliente
LEFT JOIN Pessoa_Fisica pf ON cli.Id_cliente = pf.Id_cliente
LEFT JOIN Pessoa_Juridica pj ON cli.Id_cliente = pj.Id_cliente
WHERE cartao.Status = 'ativo';

-- h) Listar funcionários administrativos de determinada agência
SELECT f.Nome, a.Num_agencia, adm.Setor, adm.Nivel_acesso
FROM Administrativo adm
JOIN Funcionario f ON adm.Matricula = f.Matricula
JOIN Agencia a ON adm.Num_agencia = a.Num_agencia
WHERE a.Num_agencia = 1;

-- i) Ver saldo total de todas as contas de um cliente
SELECT cli.Id_cliente, COALESCE(pf.Nome_completo, pj.Razao_social) AS nome, SUM(c.Saldo) AS saldo_total
FROM Cliente cli
LEFT JOIN Pessoa_Fisica pf ON cli.Id_cliente = pf.Id_cliente
LEFT JOIN Pessoa_Juridica pj ON cli.Id_cliente = pj.Id_cliente
JOIN Conta c ON cli.Id_cliente = c.Id_cliente
GROUP BY cli.Id_cliente, pf.Nome_completo, pj.Razao_social;

-- j) Mostrar todas as transações de transferência (exibindo origem, destino e valor)
SELECT t.Id_transacao, t.Data, t.Hora, t.Valor, t.Num_conta_origem, t.Num_conta_destino
FROM Transacao t
WHERE t.Tipo = 'transferencia'
ORDER BY t.Data, t.Hora;

-- k) Listar os clientes (nome ou razão social) que têm pelo menos um dependente
SELECT DISTINCT COALESCE(pf.Nome_completo, pj.Razao_social) AS titular, cli.Id_cliente
FROM Cliente cli
LEFT JOIN Pessoa_Fisica pf ON cli.Id_cliente = pf.Id_cliente
LEFT JOIN Pessoa_Juridica pj ON cli.Id_cliente = pj.Id_cliente
JOIN Dependente d ON cli.Id_cliente = d.Id_cliente;

-- l) Listar o número de contas por tipo (corrente/poupança)
SELECT Tipo, COUNT(*) AS quantidade
FROM Conta
GROUP BY Tipo;

-- m) Mostrar todos os dependentes e quais contas eles têm acesso
SELECT d.Nome_dependente, dc.Num_conta
FROM Dependente d
JOIN Dependente_Conta dc ON d.Id_dependente = dc.Id_dependente
ORDER BY d.Nome_dependente;

-- n) Listar todos os gerentes e as agências que gerenciam
SELECT f.Nome AS nome_gerente, g.Matricula, a.Num_agencia, a.Endereco AS agencia
FROM Gerente g
JOIN Funcionario f ON g.Matricula = f.Matricula
JOIN Agencia a ON a.Gerente_agencia = g.Matricula;

-- o) Listar todos os cartões vencidos
SELECT Num_cartao, Bandeira, Tipo, Validade, Status
FROM Cartao
WHERE Validade < CURRENT_DATE;

-- p) Mostrar todas as contas e seu saldo ordenado do maior para o menor
SELECT Num_conta, Saldo
FROM Conta
ORDER BY Saldo DESC;

-- q) Mostrar todas as contas inativas e seus titulares
SELECT c.Num_conta, COALESCE(pf.Nome_completo, pj.Razao_social) AS titular, c.Status
FROM Conta c
JOIN Cliente cli ON c.Id_cliente = cli.Id_cliente
LEFT JOIN Pessoa_Fisica pf ON cli.Id_cliente = pf.Id_cliente
LEFT JOIN Pessoa_Juridica pj ON cli.Id_cliente = pj.Id_cliente
WHERE c.Status = 'inativa';

-- r) Listar o número de cartões ativos por conta
SELECT Num_conta, COUNT(*) AS total_cartoes_ativos
FROM Cartao
WHERE Status = 'ativo'
GROUP BY Num_conta;

-- s) Listar todas as transações acima de R$ 1000
SELECT *
FROM Transacao
WHERE Valor > 1000
ORDER BY Valor DESC;
