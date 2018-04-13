#!/bin/bash
# NOTE: move each program to $GOPATH/src or etc...

go build -o domainfinder

cd synonyms
go build -o ../lib/synonyms

cd whois
go build -o ../lib/whois

cd sprinkle
go build -o ../lib/sprinkle

cd coolify
go build -o ../lib/coolify

cd domainify
go build -o ../lib/domainify
