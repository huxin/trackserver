#!/bin/bash
set -x
set -e

cd /root/trackserver
go get -u github.com/gorilla/mux
go build
