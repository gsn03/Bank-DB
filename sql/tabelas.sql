-- =========================
-- FUNCIONÁRIOS E SUBCLASSES
-- =========================

CREATE TABLE Funcionario (
    Matricula      SERIAL PRIMARY KEY,      -- SERIAL para auto-incremento
    Nome           VARCHAR(100) NOT NULL,
    Telefone       VARCHAR(20),
    Cargo          VARCHAR(50) NOT NULL,      -- Exemplo: 'Gerente' ou 'Administrativo'
    Data_admissao  DATE NOT NULL,
    Login          VARCHAR(50) UNIQUE NOT NULL,
    Senha          VARCHAR(255) NOT NULL
);

CREATE TABLE Gerente (
    Matricula      INT PRIMARY KEY,
    Aprovar_limite BOOLEAN NOT NULL,
    FOREIGN KEY (Matricula) REFERENCES Funcionario(Matricula)
);

CREATE TABLE Administrativo (
    Matricula      INT PRIMARY KEY,
    Setor          VARCHAR(50),
    Nivel_acesso   INT,
    Gerente_resp   INT,  -- Supervisor direto (Gerente)
    Num_agencia    INT,  -- Agência onde trabalha (FK criada depois)
    FOREIGN KEY (Matricula) REFERENCES Funcionario(Matricula),
    FOREIGN KEY (Gerente_resp) REFERENCES Gerente(Matricula)
);

-- =============
-- CLIENTES E SUBCLASSES
-- =============

CREATE TABLE Cliente (
    Id_cliente     SERIAL PRIMARY KEY,
    Endereco       VARCHAR(255),
    Telefone       VARCHAR(20)
);

CREATE TABLE Pessoa_Fisica (
    Id_cliente     INT PRIMARY KEY,
    CPF            VARCHAR(14) UNIQUE NOT NULL,
    Nome_completo  VARCHAR(100) NOT NULL,
    Data_nasc      DATE,
    FOREIGN KEY (Id_cliente) REFERENCES Cliente(Id_cliente)
);

CREATE TABLE Pessoa_Juridica (
    Id_cliente     INT PRIMARY KEY,
    CNPJ           VARCHAR(18) UNIQUE NOT NULL,
    Razao_social   VARCHAR(100) NOT NULL,
    FOREIGN KEY (Id_cliente) REFERENCES Cliente(Id_cliente)
);

-- =============
-- DEPENDENTES
-- =============

CREATE TABLE Dependente (
    Id_dependente    SERIAL PRIMARY KEY,
    Nome_dependente  VARCHAR(100) NOT NULL,
    Endereco         VARCHAR(255),
    Telefone         VARCHAR(20),
    Id_cliente       INT NOT NULL,  -- Cliente do qual depende
    FOREIGN KEY (Id_cliente) REFERENCES Cliente(Id_cliente)
);

-- =============
-- AGÊNCIAS
-- =============

CREATE TABLE Agencia (
    Num_agencia       SERIAL PRIMARY KEY,       -- SERIAL para auto-incremento
    Endereco          VARCHAR(255) NOT NULL,
    Telefone          VARCHAR(20),
    Gerente_agencia   INT UNIQUE NOT NULL,      -- Um gerente só pode gerenciar uma agência
    FOREIGN KEY (Gerente_agencia) REFERENCES Gerente(Matricula)
);

-- FK para agência em Administrativo (após criação de Agencia)
ALTER TABLE Administrativo
ADD CONSTRAINT fk_agencia_admin FOREIGN KEY (Num_agencia) REFERENCES Agencia(Num_agencia);

-- =============
-- CONTAS E RELAÇÕES
-- =============

CREATE TABLE Conta (
    Num_conta        SERIAL PRIMARY KEY,       -- SERIAL para auto-incremento
    Tipo             VARCHAR(10) NOT NULL CHECK (Tipo IN ('corrente', 'poupanca')),
    Limite           DECIMAL(12,2),
    Saldo            DECIMAL(12,2) DEFAULT 0,
    Status           VARCHAR(10) NOT NULL,     -- ativa/inativa
    Data_abertura    DATE NOT NULL,
    Senha            VARCHAR(255) NOT NULL,
    Aberta_por       INT NOT NULL,             -- Matrícula do funcionário (gerente)
    Id_cliente       INT NOT NULL,             -- Titular da conta
    Num_agencia      INT NOT NULL,             -- Agência vinculada
    FOREIGN KEY (Aberta_por) REFERENCES Funcionario(Matricula),
    FOREIGN KEY (Id_cliente) REFERENCES Cliente(Id_cliente),
    FOREIGN KEY (Num_agencia) REFERENCES Agencia(Num_agencia)
    -- Regra de negócio: PJ não pode abrir conta poupança (aplicada no backend, não no SQL puro)
);

-- Subclasses de Conta
CREATE TABLE Conta_corrente (
    Num_conta INT PRIMARY KEY,
    FOREIGN KEY (Num_conta) REFERENCES Conta(Num_conta)
);

CREATE TABLE Conta_poupanca (
    Num_conta INT PRIMARY KEY,
    FOREIGN KEY (Num_conta) REFERENCES Conta(Num_conta)
);

-- Dependentes com acesso a contas
CREATE TABLE Dependente_Conta (
    Id_dependente  INT,
    Num_conta      INT,
    PRIMARY KEY (Id_dependente, Num_conta),
    FOREIGN KEY (Id_dependente) REFERENCES Dependente(Id_dependente),
    FOREIGN KEY (Num_conta) REFERENCES Conta(Num_conta)
);

-- =============
-- CARTÕES
-- =============

CREATE TABLE Cartao (
    Num_cartao        VARCHAR(25) PRIMARY KEY,
    Cod_verificacao   VARCHAR(10) NOT NULL,
    Limite            DECIMAL(12,2),
    Validade          DATE NOT NULL,
    Status            VARCHAR(20) NOT NULL,
    Bandeira          VARCHAR(30),
    Tipo              VARCHAR(10) NOT NULL,    -- credito/debito
    Num_conta         INT NOT NULL,
    FOREIGN KEY (Num_conta) REFERENCES Conta(Num_conta)
);

-- =============
-- TRANSACOES
-- =============

CREATE TABLE Transacao (
    Id_transacao    SERIAL PRIMARY KEY,
    Data            DATE NOT NULL,
    Hora            TIME NOT NULL,
    Valor           DECIMAL(12,2) NOT NULL,
    Tipo            VARCHAR(20) NOT NULL,         -- deposito, saque, transferencia
    Num_conta_origem   INT,                       -- pode ser NULL para depósito
    Num_conta_destino  INT,                       -- pode ser NULL para saque
    FOREIGN KEY (Num_conta_origem) REFERENCES Conta(Num_conta),
    FOREIGN KEY (Num_conta_destino) REFERENCES Conta(Num_conta)
);
