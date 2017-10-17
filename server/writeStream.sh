#!/bin/bash

log_term() {
        echo "SIGTERM received"
	exit 0
}


trap log_term SIGINT SIGTERM SIGKILL


#cameraSet.sh

sleep 30

while :
do
    timestamp=$( date +%F_%T )
    #ffmpeg -i http://user:fktrc-fk@109.86.198.152:8090/stream/video.mjpeg -fs 1000000000 -vcodec flv "/home/al/motion/home-$timestamp.flv"
    cvlc http://109.86.198.152:8090 --run-time 7200 --sout="#std{access=file,mux=ogg,dst=/home/al/motion/home-$timestamp.ogg}" vlc://quit
    
    if [[ $? -ne 0 ]]; then
        echo 'Error reading data, timeout 10 secs'
        sleep 30
	#cameraSet.sh
    fi
done
