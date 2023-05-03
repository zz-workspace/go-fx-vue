#!/bin/sh
set -e
echo "starting psql"

env="$1"

su postgres -c 'pg_ctl start -D /var/lib/postgresql/data'

if su postgres -c "psql -tAc \"SELECT 1 FROM pg_roles WHERE rolname='myuser'\"" | grep -q 1; then
    echo "User already exists"
else
    su postgres -c "psql -c \"CREATE ROLE myuser WITH LOGIN PASSWORD 'abc@123';\""
    su postgres -c "psql -c \"ALTER ROLE myuser WITH CREATEDB;\""
    su postgres -c 'createdb -U myuser demo'
    echo "User created"
fi

if [ $env = "development" ]
then
    echo "AIR WORKING"
    air
else
    /golang-my-app
fi
