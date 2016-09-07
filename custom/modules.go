package main

import (
        _ "github.com/gliderlabs/logspout/adapters/syslog"
        _ "github.com/gliderlabs/logspout/transports/tcp"
        _ "gitlab.com/itsdp/logspout-logstash"
        _ "github.com/gliderlabs/logspout/transports/udp"
)
