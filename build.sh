#!/usr/bin/bash

set -x

rm build/pfcalcli

go build -o build/pfcalcli pfcalcli/cmd/pfcalcli && \
	strip build/pfcalcli && \
	upx --ultra-brute build/pfcalcli
