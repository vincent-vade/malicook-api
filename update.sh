#!/bin/bash
cd /malicook-api
git pull
docker-compose pull && docker-compose up -d