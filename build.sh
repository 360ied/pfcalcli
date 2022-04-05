#!/usr/bin/env bash

set -x

rm build/pfcalcli

go build -o build/pfcalcli -trimpath -ldflags \
	"-X main.build=$(cat version.txt) -w -s" \
	pfcalcli/cmd/pfcalcli &&
	upx --ultra-brute build/pfcalcli
