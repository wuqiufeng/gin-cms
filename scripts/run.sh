#!/bin/bash

export MYSQL_PASSWORD=123456
export MYSQL_HOST=127.0.0.1
export MYSQL_DB=lime
export MYSQL_PORT=3306
export MYSQL_USERNAME=root

./cmsgo server ./config/in-local.yaml -p 8080