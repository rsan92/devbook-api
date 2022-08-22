-- Active: 1646103809076@@127.0.0.1@3306@devbook
CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;


CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) NOT NULL,
    nick varchar(50) NOT NULL UNIQUE,
    email varchar(50) NOT NULL UNIQUE,
    senha varchar(100) NOT NULL,
    criadoEm TIMESTAMP DEFAULT current_timestamp()
) ENGINE=INNODB;

CREATE TABLE seguidores(
    usuario_id int not null,
    FOREIGN KEY (usuario_id) REFERENCES usuarios (id)
    ON DELETE CASCADE,

    seguidor_id int not null,
    FOREIGN KEY (usuario_id) REFERENCES usuarios (id)
    ON DELETE CASCADE,

    PRIMARY KEY (usuario_id, seguidor_id)
) ENGINE=INNODB;

CREATE TABLE publicacoes(
    id int AUTO_INCREMENT PRIMARY KEY,
    titulo VARCHAR(50) not null,
    conteudo VARCHAR(300) not null,
    autor_id int not null,
    FOREIGN key(autor_id) REFERENCES usuarios(id) ON DELETE CASCADE,
    curtidas int DEFAULT 0,
    criadaEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=INNODB;