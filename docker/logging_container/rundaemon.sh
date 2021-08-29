#!/usr/bin/env bash

docker rm -f BBP-DAEMON-A
docker run --name BBP-DAEMON-A -d -v /bitlog:/bitlog --restart=always -t bbplogger /bin/logger -log_dir /bitlog/BBP -flag_file  /bitlog/BBPFLAG -exit_wait 240
sleep 240


docker rm -f BBP-DAEMON-B
docker run --name BBP-DAEMON-B -d -v /bitlog:/bitlog --restart=always -t bbplogger /bin/logger -log_dir /bitlog/BBP -flag_file  /bitlog/BBPFLAG -exit_wait 240






