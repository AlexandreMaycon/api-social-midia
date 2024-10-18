CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE user(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(100) not null,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE seguidores(
    usuario_id int not null,
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    seguidor_id int not null,
    FOREIGN KEY (seguidor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    primary key(usuario_id, seguidor_id)
) ENGINE=INNODB;

CREATE TABLE publicacoes(
    id int auto_increment primary key,
    titulo varchar(50) not null,
    conteudo varchar(300) not null,

    autor_id int not null,
    FOREIGN KEY (autor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    curtidas int default 0,
    criadaEm timestamp default current_timestamp
) ENGINE=INNODB;

-- DATA

insert into usuarios (nome, nick, email, senha)
values
("Usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"), -- usuario1
("Usuário 2", "usuario_2", "usuario2@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"), -- usuario2
("Usuário 3", "usuario_3", "usuario3@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"); -- usuario3

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicação do Usuário 1", "Essa é a publicação do usuário 1! Oba!", 1),
("Publicação do Usuário 2", "Essa é a publicação do usuário 2! Oba!", 2),
("Publicação do Usuário 3", "Essa é a publicação do usuário 3! Oba!", 3);