CREATE DATABASE IF NOT EXISTS go_restfulapi;

CREATE TABLE IF NOT EXISTS category(
id int primary key auto_increment,
name varchar(200) not null
) engine = innodb;