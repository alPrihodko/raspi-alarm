#!/bin/bash

XML='<?xml version="1.0" encoding="UTF-8"?>
<request>
    <auth>
        <login>380939760324</login>
        <password>fktrc-fk</password>
    </auth>
    <message>
          <from>Test</from>
          <text>Move detected</text>
          <recipient>380939760324</recipient>
    </message>
</request>'

curl -vvv -d "$XML" http://letsads.com/api
