CREATE DATABASE go_restfulapi;

CREATE TABLE category(
id int primary key auto_increment,
name varchar(200) not null
) engine = innodb;

SELECT * FROM category;