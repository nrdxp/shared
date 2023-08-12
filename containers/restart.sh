#!/bin/bash

stop() {
	kill -9 "$(pgrep "${1}")"
	exit
}

trap stop

until ./m "${2}"; do
	echo "Restarting..."
	sleep 1
done
