#!/usr/bin/env bash

set -x

rm build/pfcalcli

go build -o build/pfcalcli -ldflags \
	"-X main.build=$(date -uIminutes) -w -s" \
	pfcalcli/cmd/pfcalcli &&
	upx --ultra-brute build/pfcalcli
