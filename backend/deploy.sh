#!/bin/bash
GOOS=linux GOARCH=amd64 go build -ldflags="-X main.isBuild=prod" -o pms
#strip -s  postpusher

ssh root@178.62.14.228 'rm -rf /opt/pms/pms_app'
scp pms root@178.62.14.228:/opt/pms/pms_app
scp -r /Users/mattik/work/src/gitlab.com/pasinworld/pms/backend/* root@178.62.14.228:/opt/pms/
ssh -t root@178.62.14.228 'sudo killall pms_app'
