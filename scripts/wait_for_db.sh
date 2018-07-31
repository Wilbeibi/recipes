#!/bin/bahs
util $(nc -z ${DB_HOST} ${DB_PORT})
do
    echo "$(date +%T) - ${DB_HOST}:${DB_PORT} service is not ready, wait 3 seconds..."
    sleep 3
done

echo "$(date +%T) - service is ready"