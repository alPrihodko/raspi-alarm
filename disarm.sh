#!/bin/bash
/etc/init.d/uv4l_uvc stop
ssh -p 22095 al@82.117.238.247 "sudo service uvc stop"