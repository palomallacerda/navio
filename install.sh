#!/bin/bash

# CREATE NAVIO DATABASE
mysqlversion=`mysql --version`
goversion=`go version`
if [ -n "$goversion"  ] && [ -n "$mysqlversion"  ]; then 
    go run ./database/up.go
fi

# ...
sudo cp ./navio /usr/local/bin
