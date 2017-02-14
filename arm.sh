#!/bin/bash
#/etc/init.d/uv4l_uvc stop
#/etc/init.d/uv4l_uvc start 046d:081b
#sleep 2
curl 'http://192.168.1.46:8090/panel?width=1280&height=720&format=1196444237&apply_changed=1' -H 'Pragma: no-cache' -H 'Accept-Encoding: gzip, deflate, sdch' -H 'Accept-Language: en-US,en;q=0.8' -H 'Upgrade-Insecure-Requests: 1' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.84 Safari/537.36' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8' -H 'Cache-Control: no-cache' -H 'Authorization: Basic YWRtaW46Zmt0cmMtZms=' -H 'Connection: keep-alive' -H 'Referer: http://192.168.1.46:8090/panel?width=1280&height=720&format=1196444237&apply_changed=1' --compressed

ssh -p22095 al@82.117.238.247 "sudo service uvc stop;sudo service uvc start"
