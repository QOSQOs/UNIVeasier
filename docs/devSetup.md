# Development setup

# Linux Ubuntu

## Backend tools

### Install Golang
Choose the archive file appropriate for your installation. For instance, if you are installing Go version `1.2.1` for `64-bit x86` on `Linux`: 

`tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz`

Download the archive and extract it into `/usr/local`.

```bash
$ tar -C /usr/local -xzf go1.13.1.linux-amd64.tar.gz
```

Add to the PATH environment variable.

```bash
$ export GOPATH=$HOME/go
$ export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

Check if the installation was successful.

```bash
$ go version`
```

### Install MySQL database
Add the MySQL APT repository to your system's software repository list.

```bash
$ wget –c https://dev.mysql.com/get/mysql-apt-config_0.8.10-1_all.deb
$ sudo dpkg --install mysql-apt-config_0.8.10-1_all.deb
```

Update package information from the MySQL APT repository.

```bash
$ sudo apt-get update
```

Install MySQL and start the server. During the installation, you are asked to supply a password for the root user for your MySQL installation. Therefore provide as user `root` and choose a password.

```bash
$ sudo apt-get install mysql-server
$ sudo service mysql status
$ sudo service mysql start
```

Install and run MySQL Workbench:

```bash
$ sudo apt install mysql-workbench
$ mysql-workbench
```

### Run project

Add to the PATH environment variable. The configuration filepath `$CONFIG` and the project path `$REPO`.

```bash
$ export CONFIG=$REPO/internal/config/config.json
$ export REPO=$GOPATH/src/github.com/QOSQOs/UNIVeasier
```

Execute the database scripts. Then, open MySQL Workbench and connect with the user `root` , password `qosqo123` and run the following scripts:

``` bash
$ cd $REPO/internal/scripts/databases/dbtest.sql
$ cd $REPO/internal/scripts/dummy_data/dbtest.sql
$ cd $REPO/internal/scripts/procedures/test.sql
```

Clone the project from our github repository

```bash
$ go get github.com/QOSQOs/UNIVeasier
```

Change the default values ​​in the configuration file located in `$REPO/internal/config/config.json`

```json
{
    "server": {
        "host": "localhost",
        "port": "8000"
    },
    "db": {
        "driver": "mysql",
        "name": "dbtest",
        "username": "root",
        "password": "your-user-password",
        "protocol": "tcp",
        "host": "localhost",
        "port": "your-database-port"
    }
}
```

Run the following command to find the port on which your MySQL installation is serving.

```sql
SHOW GLOBAL VARIABLES LIKE 'PORT';
```

Build the project binary

```bash
$ go install $REPO/cmd
```

Run the project executable

```bash
$ cmd
```

Test the REST API

``` bash
$ curl -s -X GET http://localhost:8000/test
$ curl -s -X GET http://localhost:8000/test/{name}
```

### Run project unit test

Run an especific unit test of the project. Go to the folder where the file containing it test is located and run:

```bash
$ cd $REPO/pkg/model/test
$ go test -v
```

Run all unit tests of the project:

``` bash
$ cd $REPO
$ go test ./...
```

## Front-end tools
### Install Node.js and npm

Installation instructions:

``` bash
$ curl -sL https://deb.nodesource.com/setup_12.x | sudo -E bash -
$ sudo apt-get install -y nodejs
$ node -v
$ npm -v
```

### Install Vue CLI

To install the new package, use one of the following commands. You need administrator privileges to execute these unless npm was installed on your system through a Node.js version manager.

``` bash
$ sudo npm install -g @vue/cli
$ vue --version
$ sudo npm install -g @vue/cli-service-global
```

### Run project

Clone the project from our github repository

```bash
$ go get github.com/QOSQOs/UNIVeasier-ui
```

Init the server

``` bash
$ cd $REPO/../UNIVeasier-ui
$ npm install
$ npm run serve
```

Open your browser and go to the server url

```
 App running at:
  - Local:   http://localhost:8080
  - Network: http://localhost:8080
```

### References
1. Golang: https://golang.org/doc/install
2. MySQL: https://dev.mysql.com/doc/mysql-apt-repo-quick-guide/
3. NodeJS: https://github.com/nodesource/distributions/blob/master/README.md
4. Vue: https://cli.vuejs.org/guide/installation.html
