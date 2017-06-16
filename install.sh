#!/bin/bash
set -x
set -e

cd /root/trackserver
cp trackserver.upstart /etc/init/trackserver.conf
start trackserver