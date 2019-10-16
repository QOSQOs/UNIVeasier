# Development setup

# Linux Ubuntu

## Backend tools

### Install Golang
Choose the archive file appropriate for your installation. For instance, if you are installing Go version `1.2.1` for `64-bit x86` on `Linux`: 

`tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz`

Then, download the archive and extract it into `/usr/local`.

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
First, add the MySQL APT repository to your system's software repository list.

```bash
$ wget –c https://dev.mysql.com/get/mysql-apt-config_0.8.10-1_all.deb
$ sudo dpkg --install mysql-apt-config_0.8.10-1_all.deb
```

Update package information from the MySQL APT repository.

```bash
$ sudo apt-get update
```

Later, install MySQL and start the server. During the installation, you are asked to supply a password for the root user for your MySQL installation. Therefore provide as user `root` and password `qosqo123`.

```bash
$ sudo apt-get install mysql-server
$ sudo service mysql status
$ sudo service mysql start
```

To install and run MySQL Workbench:

```bash
$ sudo apt install mysql-workbench
$ mysql-workbench
```

### Run project

Add to the PATH environment variable.

```bash
$ export REPO=$GOPATH/src/github.com/QOSQOs/UNIVeasier
```

First we need to execute the database scripts. Then, open MySQL Workbench and connect with the user `root` , password `qosqo123` and execute the following scripts:

``` bash
$ cd $REPO/internal/scripts/databases/dbtest.sql
$ cd $REPO/internal/scripts/dummy_data/dbtest.sql
$ cd $REPO/internal/scripts/procedures/test.sql
```

Clone the project from our github repository

```bash
$ go get github.com/QOSQOs/UNIVeasier
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

In order to run the unit test. To run an especific test, go to the folder where the file containing the tests is located and run:

```bash
$ cd $REPO/pkg/model/test
$ go test -v
```

To run all unit tests of the project:

``` bash
$ cd $REPO
$ go test ./...
```

## Front-end tools
### Install nodejs and npm
``` bash
$ curl -sL https://deb.nodesource.com/setup_12.x | sudo -E bash -
$ sudo apt-get install -y nodejs
$ node -v
$ npm -v
```

### Install Vue
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
$ npm run serve
```

To test, open your browser and go to the server url

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