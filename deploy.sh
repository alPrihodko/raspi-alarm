#!/bin/bash
HOST=192.168.1.46
deploy=pi\@"$HOST":/usr/local/bin

scp alarm.sh $deploy
scp alarm_sms.sh $deploy
scp arm.sh $deploy
scp disarm.sh $deploy

export GOOS=linux
export GOARCH=arm
export GOARM=7
go build || exit 1

ssh pi\@$HOST "sudo service raspi-alarm stop"
scp raspi-alarm $deploy
ssh pi\@$HOST "sudo service raspi-alarm start"
echo "Success"
