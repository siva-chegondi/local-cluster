#!/bin/sh

if [ $1 = "vault" ]; then
    which go > /dev/null
    if [ "$?" -gt "0" ]; then
        echo "No Go environment"
    else
        echo "*** Building helper tools ***"
        go build -o bin/read_token helpers/read_token/read_token.go
        go build -o bin/store_token helpers/store_token/store_token.go
    fi
fi

echo "*** Destroy the old state ***"
vagrant destroy -f

echo "*** Running docker manager instance ***"
vagrant up docker-manager

echo "*** Running docker node instance ***"
vagrant up docker-node
