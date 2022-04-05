#!/usr/bin/env bash

set -x

rm build/pfcalcli

go build -o build/pfcalcli -ldflags "-X main.build=$(date --iso-8601=minutes --utc)" pfcalcli/cmd/pfcalcli &&
	strip build/pfcalcli &&
	upx --ultra-brute build/pfcalcli
