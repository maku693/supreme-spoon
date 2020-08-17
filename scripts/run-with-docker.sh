#!/bin/sh
set -eu

CONTAINER_ID=$(docker run -d --rm -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=test -p 127.0.0.1:3307:3306 mysql:8.0)
printf 'waiting for db to start up'
until mysql -h 127.0.0.1 -P 3307 -uroot -proot -e 'SELECT 1;' &>/dev/null; do
    printf .
    sleep 1
done
echo 'done.'
SQL_DSN="root:root@tcp(127.0.0.1:3307)/test" go run github.com/maku693/supreme-spoon
docker kill $CONTAINER_ID &>/dev/null
