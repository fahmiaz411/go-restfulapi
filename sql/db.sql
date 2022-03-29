CREATE DATABASE go_restfulapi;

CREATE TABLE category(
id int primary key auto_increment,
name varchar(200) not null
) engine = innodb;

SELECT * FROM category;

INSERT INTO category(name) VALUES ("fahmi");
DELETE FROM category WHERE id = 1