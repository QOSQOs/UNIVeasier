SHOW GLOBAL VARIABLES LIKE 'PORT';

CREATE DATABASE dbtest;

SHOW DATABASES;

USE dbtest;

CREATE TABLE IF NOT EXISTS dbtest.tabletest (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    email VARCHAR(50) NOT NULL
)  ENGINE=INNODB;

SHOW TABLES IN dbtest;