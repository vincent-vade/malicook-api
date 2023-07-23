#!/bin/bash
cd /home/deploy/malicook-api
git pull
make down && make build && docker image prune -a -f