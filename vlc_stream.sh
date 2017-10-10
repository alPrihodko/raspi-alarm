#!/bin/bash

log_term() {
        echo "SIGTERM received"
	exit 0
}


trap log_term SIGINT SIGTERM SIGKILL

while [  1 -eq 1 ]; do
    cvlc  --color --http-port=8090 v4l:///dev/video0:norm=secam:frequency=543250:size=1280x768:channel=0 --sout '#transcode{vcodec=mp4v,acodec=mpga,vb=512,ab=32,venc=ffmpeg{keyint=80,hurry-up,vt=800000},deinterlace}:standard{access=http,mux=ogg}' --ttl 12 vlc://quit
    sleep 2
done