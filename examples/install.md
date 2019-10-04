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
    name VARCHAR(255) NOT NULL
)  ENGINE=INNODB;

SHOW TABLES IN dbtest;

INSERT INTO dbtest.tabletest(name) VALUES ('Kleiber');
INSERT INTO dbtest.tabletest(name) VALUES ('Darwin');
INSERT INTO dbtest.tabletest(name) VALUES ('Dany');

DELETE FROM dbtest.tabletest WHERE name = 'Kleiber';

UPDATE dbtest.tabletest SET name = 'unknow' WHERE name = 'Kleiber';

SELECT * FROM dbtest.tabletest;

DROP TABLE dbtest.tabletest;

DROP DATABASE dbtest;
