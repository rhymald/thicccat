#!/bin/bash

sudo docker rm $(sudo docker ps -a -f status=exited -q) && sudo docker rmi $(sudo docker images -a -q)
sudo sh -c 'truncate -s 0 /var/lib/docker/containers/*/*-json.log'
set -eu ; LANG=en_US.UTF-8 snap list --all | awk '/disabled/{print $1, $3}' | while read snapname revision; do ; snap remove "$snapname" --revision="$revision" ; done