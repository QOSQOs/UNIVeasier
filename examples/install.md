## Install MySQL database

`wget –c https://dev.mysql.com/get/mysql-apt-config_0.8.10-1_all.deb`

`sudo dpkg --install mysql-apt-config_0.8.10-1_all.deb`

`sudo apt-get update`

`sudo apt-get upgrade`

`sudo apt-get install mysql-server`

`USER: root`

`PASSWORD: qosqo123`

`sudo service mysql status`

`sudo service mysql stop`

`sudo service mysql start`

`sudo mysql -u {USER} -p {PASSWORD}`
`sudo mysql -u root -p`

## Install Golang language

`tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz`

`tar -C /usr/local -xzf go1.13.1.linux-amd64.tar.gz.`

`export PATH=$PATH:/usr/local/go/bin`

`export GOPATH=$HOME/go`

## Test database connection

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

DELIMITER //

CREATE PROCEDURE dbtest.GetListTest()
BEGIN
    SELECT *  FROM dbtest.tabletest;
END //

CREATE PROCEDURE dbtest.GetTest(_name VARCHAR(255))
BEGIN
    SELECT *  FROM dbtest.tabletest WHERE name = _name;
END //

DELIMITER ;

SHOW TABLES IN dbtest;

INSERT INTO dbtest.tabletest(name, age, email) VALUES ('Kleiber', '27', 'kleiber100892@gmail.com');
INSERT INTO dbtest.tabletest(name, age, email) VALUES ('Darwin', '24', 'darwin.ttito.c@gmail.com');
INSERT INTO dbtest.tabletest(name, age, email) VALUES ('Dany', '26', 'florezatauchi@gmail.com');

SELECT * FROM dbtest.tabletest;

DELETE FROM dbtest.tabletest WHERE name = 'Kleiber';

UPDATE dbtest.tabletest SET name = 'unknow' WHERE name = 'Kleiber';

SELECT * FROM dbtest.tabletest;

DROP TABLE dbtest.tabletest;

DROP DATABASE dbtest;
