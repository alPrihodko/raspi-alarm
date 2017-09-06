#!/bin/bash

cameraSet.sh

while :
do
    #for i in {1..200}
    #do
    timestamp=$( date +%F_%T )
    ffmpeg -i http://user:fktrc-fk@109.86.198.152:8090/stream/video.mjpeg -fs 1000000000 -vcodec flv "/home/al/motion/home-$timestamp.flv"

    if [[ $? -ne 0 ]]; then
        echo 'Error reading data, timeout 10 secs'
        sleep 5
	#cameraSet.sh
    fi
    #done
done
