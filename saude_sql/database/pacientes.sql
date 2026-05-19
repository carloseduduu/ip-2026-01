CREATE TABLE IF NOT EXISTS pacientes (
    cod_SUS VARCHAR(20) PRIMARY KEY,
    nome VARCHAR(100),
    data_nascimento DATE,
    sexo VARCHAR(10),
    tipo_sanguineo VARCHAR(3) CHECK (tipo_sanguineo IN ('A+', 'A-', 'B+', 'B-', 'AB+', 'AB-', 'O+', 'O-'))
);