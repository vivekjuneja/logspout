package main

import (
        _ "github.com/gliderlabs/logspout/adapters/syslog"
        _ "github.com/gliderlabs/logspout/transports/tcp"
        _ "github.com/vivekjuneja/logspout-logstash"
        _ "github.com/gliderlabs/logspout/transports/udp"
)
