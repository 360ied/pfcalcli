#!/usr/bin/bash

set -x

rm pfcalcli

go build && \
	strip pfcalcli && \
	upx --ultra-brute pfcalcli
