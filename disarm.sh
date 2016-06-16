#!/bin/bash
/etc/init.d/uv4l_uvc stop
ssh -p 22095 al@office.webdevs.com "sudo service uvc stop"