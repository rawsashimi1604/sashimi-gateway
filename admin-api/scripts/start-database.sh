#!/bin/bash

echo "****************************************"
echo "Start Admin-API Database"
echo "****************************************"

cd docker;
docker compose down;
docker compose up --detach;