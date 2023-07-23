#!/bin/bash
cd /home/deploy/malicook-api/
git pull
docker compose down && docker-compose pull && docker-compose up -d