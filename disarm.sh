#!/bin/bash
service vlc stop
ssh -p 22095 al@office.webdevs.us "sudo service uvc stop"