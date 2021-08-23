#!/bin/bash
echo "build image -> start"
docker build --no-cache -t kevin24ec/tvseries_db:dev --file build/Dockerfile .
echo "build image -> complete"
echo "------------------------------------------------------------------------------"
echo "push image -> start"
docker push kevin24ec/tvseries_db:dev
echo "push image -> complete"