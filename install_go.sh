#!/usr/bin/bash
set -x
set -e

cd /root
wget "https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz" -O go.tar.gz
tar zxvf go.tar.gz
echo "export GOROOT=/root/go/" >> /root/.bashrc
mkdir /root/proj
echo "export GOPATH=/root/proj" >> /root/.bashrc
echo "export PATH=$PATH:$GOROOT/bin/"
