#!/bin/bash
HOST=192.168.1.46
deploy=pi\@"$HOST":/usr/local/bin

scp alarm.sh $deploy || exit 1
scp alarm_sms.sh $deploy || exit 1
scp arm.sh $deploy || exit 1
scp disarm.sh $deploy || exit 1

export GOOS=linux
export GOARCH=arm
export GOARM=7

if [ "$1" -eq "all" ]; then
    go build || exit 1
    ssh pi\@$HOST "sudo service raspi-alarm stop" || exit 1
    scp raspi-alarm $deploy || exit 1
    ssh pi\@$HOST "sudo service raspi-alarm start" || exit 1
    echo "Success"
fi
