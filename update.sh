#!/bin/bash
ls
cd /home/deploy/malicook-api
pwd
git pull
docker compose down && docker-compose pull && docker-compose up -d